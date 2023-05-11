package chunk

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"sync"

	"github.com/caivega/glog"
	libstore "github.com/tokentransfer/interfaces/store"
	"github.com/tokentransfer/node/core"
	"github.com/tokentransfer/node/crypto"
	"github.com/tokentransfer/node/db"
	"github.com/tokentransfer/node/util"
)

var cs = &crypto.CryptoService{}

type chunkStorage struct {
	mutex   sync.Mutex
	entries libstore.KvService

	bytesMax  int64
	bytesUsed int64
	version   int64
}

// reference from https://github.com/indyjo/cafs/blob/master/ram/ramstorage.go
func NewStorage(maxBytes int64) (core.Storage, error) {
	dbs := &db.MemoryService{
		Name: "memory",
	}
	// dbs := &db.LevelService{
	// 	Name: "test",
	// 	Path: "./test",
	// }
	err := dbs.Init(nil)
	if err != nil {
		return nil, err
	}
	err = dbs.Start()
	if err != nil {
		return nil, err
	}
	return NewStorageWith(dbs, maxBytes)
}

func NewStorageWith(db libstore.KvService, maxBytes int64) (core.Storage, error) {
	return &chunkStorage{
		entries: db,

		bytesMax: maxBytes,
	}, nil
}

func LoadStorageWith(db libstore.KvService, maxBytes int64) (core.Storage, error) {
	_s, err := NewStorageWith(db, maxBytes)
	if err != nil {
		return nil, err
	}
	s := (_s).(*chunkStorage)
	err = s.load()
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (s *chunkStorage) Root() core.Key {
	v := s.Version()
	if v >= 0 {
		key, _ := s.GetVersion(v)
		if key != nil {
			return key
		}
	}
	return core.Key{}
}

func (s *chunkStorage) Version() int64 {
	return s.version - 1
}

func (s *chunkStorage) GetVersion(n int64) (core.Key, error) {
	indexString := fmt.Sprintf("%d", n)
	link := core.GetKey(indexString)
	linkEntry, ok := s.getEntry(link)
	if !ok {
		return nil, util.ErrorOfNotFound("link", indexString)
	}
	k, err := core.ParseKey(linkEntry.name)
	if err != nil {
		return nil, err
	}
	return k, nil
}

func (s *chunkStorage) addIndex(key core.Key) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	indexString := fmt.Sprintf("%d", s.version)
	link := core.GetKey(indexString)
	if s.Exists(link) {
		return util.ErrorOf("already exists", "index", indexString)
	}
	err := s.addLink(indexString, key)
	if err != nil {
		return err
	}
	s.version++
	return nil
}

func (s *chunkStorage) addLink(name string, key core.Key) error {
	link := core.GetKey(name)
	linkEntry := &chunkEntry{
		name:   key.String(),
		data:   []byte{},
		chunks: []chunkRef{},
		refs:   1,
	}
	if _, _, err := s.putEntry(link, linkEntry); err != nil {
		return err
	}
	return nil
}

func (s *chunkStorage) add(name string, key core.Key) error {
	err := s.addLink(name, key)
	if err != nil {
		return err
	}
	err = s.addIndex(key)
	if err != nil {
		return err
	}
	return nil
}

func (s *chunkStorage) Update() error {
	return s.load()
}

func (s *chunkStorage) load() error {
	rootLink := core.GetKey("*")
	entry, ok := s.getEntry(rootLink)
	if ok {
		var bytesUsed int64
		buf := bytes.NewBuffer(entry.data)
		err := binary.Read(buf, util.BYTE_ORDER, &bytesUsed)
		if err != nil {
			return err
		}
		s.bytesUsed = bytesUsed

		var version int64
		err = binary.Read(buf, util.BYTE_ORDER, &version)
		if err != nil {
			return err
		}
		s.version = version
	} else {
		s.bytesUsed = 0
		s.entries.ListData(func(k []byte, v []byte) error {
			key, err := core.ParseKey(string(k))
			if err != nil {
				return err
			}
			var entry chunkEntry
			err = entry.UnmarshalBinary(v)
			if err != nil {
				return err
			}
			_, ok := s.getEntry(key)
			if ok {
				storageSize := entry.storageSize()
				s.bytesUsed += storageSize
			}
			return nil
		})

		version := int64(0)
		for {
			indexString := fmt.Sprintf("%d", version)
			link := core.GetKey(indexString)
			if s.Exists(link) {
				version++
			} else {
				break
			}
		}
		s.version = version
		err := s.flush()
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *chunkStorage) flush() error {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, util.BYTE_ORDER, s.bytesUsed)
	if err != nil {
		return err
	}
	err = binary.Write(buf, util.BYTE_ORDER, s.version)
	if err != nil {
		return err
	}
	root := s.Root()

	rootLink := core.GetKey("*")
	_, _, err = s.putEntry(rootLink, &chunkEntry{
		name: root.String(),
		data: buf.Bytes(),
		refs: 1,
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *chunkStorage) used(size int64) {
	s.bytesUsed += size
	err := s.flush()
	if err != nil {
		glog.Error(err)
	}
}

func (s *chunkStorage) Close() error {
	if s.entries != nil {
		err := s.entries.Close()
		if err != nil {
			return err
		}
		s.entries = nil
	}
	return nil
}

func (s *chunkStorage) _putEntry(key core.Key, entry *chunkEntry) error {
	k := []byte(key.String())
	v, err := entry.MarshalBinary()
	if err != nil {
		return err
	}
	err = s.entries.PutData(k, v)
	if err != nil {
		return err
	}
	return nil
}

func (s *chunkStorage) putEntry(key core.Key, entry *chunkEntry) (int64, bool, error) {
	oldEntry, ok := s.getEntry(key)
	if ok {
		oldSize := oldEntry.storageSize()
		storageSize := entry.storageSize()
		newSize := storageSize - oldSize
		if newSize > 0 {
			err := s.reserveBytes("update entry", newSize)
			if err != nil {
				return 0, false, err
			}
		}
		err := s._putEntry(key, entry)
		if err != nil {
			return 0, false, err
		}
		if newSize != 0 {
			s.used(newSize)
		}
		return newSize, true, nil
	}

	storageSize := entry.storageSize()
	err := s.reserveBytes("put entry", storageSize)
	if err != nil {
		return 0, false, err
	}
	err = s._putEntry(key, entry)
	if err != nil {
		return 0, false, err
	}
	s.used(storageSize)
	return storageSize, false, nil
}

func (s *chunkStorage) getEntry(key core.Key) (*chunkEntry, bool) {
	data, err := s.entries.GetData([]byte(key.String()))
	if err != nil {
		return nil, false
	}
	if len(data) == 0 {
		return nil, false
	}
	var entry chunkEntry
	err = entry.UnmarshalBinary(data)
	if err != nil {
		return nil, false
	}
	return &entry, true
}

func (s *chunkStorage) removeEntry(key core.Key) error {
	err := s.entries.RemoveData([]byte(key.String()))
	if err != nil {
		return err
	}
	return nil
}

func (s *chunkStorage) GetUsed() int64 {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.bytesUsed
}

func (s *chunkStorage) GetCapacity() int64 {
	return s.bytesMax
}

func (s *chunkStorage) Exists(key core.Key) bool {
	_, ok := s.getEntry(key)
	return ok
}

func (s *chunkStorage) Get(key core.Key) (core.Data, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	entry, ok := s.getEntry(key)
	if ok {
		entry.refs++
		_, _, err := s.putEntry(key, entry)
		if err != nil {
			return nil, err
		}
		return &chunkData{s, key, false}, nil
	}
	return nil, util.ErrorOfNotFound("file", key.String())
}

func (s *chunkStorage) Chunk(key core.Key) ([]byte, error) {
	entry, ok := s.getEntry(key)
	if ok {
		return entry.data, nil
	}
	return nil, util.ErrorOfNotFound("chunk", key.String())
}

func (s *chunkStorage) Size(key core.Key) (int64, error) {
	entry, ok := s.getEntry(key)
	if ok {
		if len(entry.chunks) > 0 {
			return entry.chunks[len(entry.chunks)-1].nextPos, nil
		}
		return int64(len(entry.data)), nil
	}
	return 0, util.ErrorOfNotFound("entry", key.String())
}

func (s *chunkStorage) Reference(key core.Key) (int64, error) {
	entry, ok := s.getEntry(key)
	if ok {
		return int64(entry.refs), nil
	}
	return 0, util.ErrorOfNotFound("entry", key.String())
}

func (s *chunkStorage) Commit(storage core.Storage, key core.Key) error {
	fb, _ := storage.Get(key)
	if fb == nil {
		fa, err := s.Get(key)
		if err != nil {
			return err
		}
		defer fa.Dispose()

		far := fa.Open()
		kvf := fa.(*chunkData)

		fb := storage.Create(kvf.Entry().name)
		io.Copy(fb, far)
		err = far.Close()
		if err != nil {
			return err
		}
		err = fb.Close()
		if err != nil {
			return err
		}
	} else {
		fb.Dispose()
	}
	return nil
}

func (s *chunkStorage) Remove(key core.Key) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	entry, ok := s.getEntry(key)
	if ok {
		err := s.release(key)
		if err != nil {
			return err
		}

		for _, chunk := range entry.chunks {
			err := s.release(chunk.key)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *chunkStorage) DumpStatistics(log core.Printer) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	link := func(k core.Key, n int, local bool) string {
		zero := core.Key{}
		if k.Equals(zero) {
			return fmt.Sprintf("%x", k[:n])
		} else if local {
			return fmt.Sprintf(`<a href="#%v">%x</a>`, k, k[:n])
		} else {
			return fmt.Sprintf(`<a href="/file/address=%v">%x</a>`, k.Address(), k[:n])
		}
	}

	log.Printf("<html><head><title>Storage Statistics</title></head><body><pre>")
	log.Printf("Root: %s", s.Root().String())
	version := s.Version()
	log.Printf("Version: %d", version+1)
	for i := int64(0); i <= version; i++ {
		k, err := s.GetVersion(i)
		if err != nil {
			log.Printf("%d: [%v]", i, err)
		} else {
			log.Printf("%d, [%v]", i, k.String())
		}
	}
	log.Printf("Bytes used: %d, max: %d", s.bytesUsed, s.bytesMax)
	s.entries.ListData(func(k []byte, v []byte) error {
		key, err := core.ParseKey(string(k))
		if err != nil {
			return err
		}
		var entry chunkEntry
		err = entry.UnmarshalBinary(v)
		if err != nil {
			return err
		}
		log.Printf(`<a name="%v"> [%v] refs=%d size=%d [%v]</a>`, key, link(key, 4, false), entry.refs, entry.storageSize(), entry.name)

		zero := core.Key{}
		if key.Equals(zero) {
			for i, chunk := range entry.chunks {
				log.Printf("               chunk %4d: %v (length %6d)", i, link(chunk.key, 4, true), chunk.nextPos)
			}
		} else {
			prevPos := int64(0)
			for i, chunk := range entry.chunks {
				log.Printf("               chunk %4d: %v (length %6d, ends at %7d)", i, link(chunk.key, 4, true), chunk.nextPos-prevPos, chunk.nextPos)
				prevPos = chunk.nextPos
			}
		}
		return nil
	})
	log.Printf("</pre></body></html>")
}

func (s *chunkStorage) DumpLog(log core.Printer) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	link := func(k core.Key, local bool) string {
		if local {
			return fmt.Sprintf("[%s *]", k.Address())
		} else {
			return fmt.Sprintf("[%s]", k.Address())
		}
	}
	log.Printf("Root: %s", s.Root().String())
	version := s.Version()
	log.Printf("Version: %d", version+1)
	for i := int64(0); i <= version; i++ {
		k, err := s.GetVersion(i)
		if err != nil {
			log.Printf("%d: [%v]", i, err)
		} else {
			log.Printf("%d: [%v]", i, k.String())
		}
	}
	log.Printf("Bytes used: %d, max %d", s.bytesUsed, s.bytesMax)
	err := s.entries.ListData(func(k []byte, v []byte) error {
		key, err := core.ParseKey(string(k))
		if err != nil {
			glog.Error(err)
			return err
		}
		var entry chunkEntry
		err = entry.UnmarshalBinary(v)
		if err != nil {
			glog.Error(err)
			return err
		}
		log.Printf("%v [%s] chunks=%d refs=%d size=%v (%v)", link(key, false), key.String(), len(entry.chunks), entry.refs, entry.storageSize(), entry.name)

		zero := core.Key{}
		if key.Equals(zero) {
			for i, chunk := range entry.chunks {
				log.Printf("               chunk %4d: %v (length %6d)", i, link(chunk.key, true), chunk.nextPos)
			}
		} else {
			prevPos := int64(0)
			for i, chunk := range entry.chunks {
				log.Printf("               chunk %4d: %v (length %6d, ends at %7d)", i, link(chunk.key, true), chunk.nextPos-prevPos, chunk.nextPos)
				prevPos = chunk.nextPos
			}
		}
		return nil
	})
	if err != nil {
		log.Printf(">>> list data: %s", err.Error())
	}
}

func (s *chunkStorage) reserveBytes(info string, numBytes int64) error {
	if numBytes > s.bytesMax {
		return util.ErrorOf("not enough", "space", fmt.Sprintf("%d", numBytes))
	}
	bytesFree := s.bytesMax - s.bytesUsed
	if bytesFree < numBytes {
		return util.ErrorOf("not enough", "space", fmt.Sprintf("%d", numBytes))
	}
	return nil
}

func (s *chunkStorage) storeEntry(key core.Key, data []byte, chunks []chunkRef, name string) error {
	if len(data) > 0 && len(chunks) > 0 {
		return util.ErrorOf("illegal", "entry", name)
	}
	s.mutex.Lock()
	defer s.mutex.Unlock()

	var newEntry *chunkEntry
	if oldEntry, _ := s.getEntry(key); oldEntry != nil {
		if len(oldEntry.data) != len(data) || len(oldEntry.chunks) != len(chunks) {
			return util.ErrorOf("key collision", "entry", fmt.Sprintf("[%v]: %v[%v]", name, key, oldEntry.name))
		}
		// glog.Infof("[%v] Recycling key: %v [%v] (data: %d bytes, chunks: %d)", info, key, oldEntry.info, len(data), len(chunks))

		s.lock(key)
		for _, chunk := range chunks {
			s.release(chunk.key)
		}

		// newEntry = oldEntry
	} else {
		newEntry = &chunkEntry{
			name:   name,
			data:   data,
			chunks: chunks,
			refs:   1,
		}
		if _, _, err := s.putEntry(key, newEntry); err != nil {
			return err
		}
		// glog.Infof("[%v] Stored key: %v (data: %d bytes, chunks: %d)", info, key, len(data), len(chunks))
	}
	return nil
}

func (s *chunkStorage) lockL(key core.Key) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.lock(key)
}

func (s *chunkStorage) lock(key core.Key) error {
	entry, ok := s.getEntry(key)
	if !ok {
		return util.ErrorOfNotFound("entry", key.String())
	}
	entry.refs++
	_, _, err := s.putEntry(key, entry)
	if err != nil {
		return err
	}
	return nil
}

func (s *chunkStorage) releaseL(key core.Key) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.release(key)
}

func (s *chunkStorage) release(key core.Key) error {
	entry, ok := s.getEntry(key)
	if !ok {
		return util.ErrorOfNotFound("entry", key.String())
	}
	if entry.refs == 0 {
		return util.ErrorOfInvalid("state", key.String())
	}
	entry.refs--
	if entry.refs == 0 {
		err := s.removeEntry(key)
		if err != nil {
			return err
		}
		s.used(-entry.storageSize())
	} else {
		_, _, err := s.putEntry(key, entry)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *chunkStorage) Create(info string) core.Temporary {
	return &chunkTemporary{
		storage:   s,
		info:      info,
		fileHash:  cs.NewHasher(),
		chunkHash: cs.NewHasher(),
		valid:     true,
		open:      true,
		chunker:   NewChunker(),
		chunks:    make([]chunkRef, 0, 16),
	}
}

func (s *chunkStorage) Group(name string) (core.Group, error) {
	dir, err := NewGroup(s, name)
	if err != nil {
		return nil, err
	}
	return dir, nil
}

type LogPrinter struct {
}

func (p LogPrinter) Printf(format string, v ...interface{}) {
	fmt.Printf(format+"\n", v...)
}

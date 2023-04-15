package chunk

import (
	"bytes"
	"fmt"
	"hash"
	"sync"

	"github.com/caivega/glog"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagcbor"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/fluent/qp"
	"github.com/ipld/go-ipld-prime/node/basicnode"
	"github.com/tokentransfer/node/core"
	"github.com/tokentransfer/node/util"
)

type chunkItem struct {
	name  string
	key   core.Key
	group bool

	g *chunkGroup
}

func (e *chunkItem) MarshalBinary() ([]byte, error) {
	node, err := qp.BuildMap(basicnode.Prototype.Any, -1, func(ma datamodel.MapAssembler) {
		qp.MapEntry(ma, "name", qp.String(e.name))
		qp.MapEntry(ma, "key", qp.Bytes(e.key[:]))
		qp.MapEntry(ma, "group", qp.Bool(e.group))
	})
	if err != nil {
		return nil, err
	}
	data, err := ipld.Encode(node, dagcbor.Encode)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (e *chunkItem) UnmarshalBinary(data []byte) error {
	node, err := ipld.Decode(data, dagcbor.Decode)
	if err != nil {
		return err
	}
	e.name = util.GetStringFromNode(node, "name")

	if k, err := util.GetBytesFromNode(node, "key"); err != nil {
		return err
	} else {
		e.key = core.Key(k)
	}

	e.group = util.GetBooleanFromNode(node, "group")

	return nil
}

type chunkGroup struct {
	storage *chunkStorage
	name    string
	hash    hash.Hash
	parent  *chunkGroup

	locker    sync.Mutex
	itemArray []*chunkItem
	itemMap   map[string]int

	itemKey   *core.Key
	itemEntry *chunkEntry
}

func NewGroup(s *chunkStorage, name string) (core.Group, error) {
	g := &chunkGroup{
		storage: s,
		name:    name,
		hash:    cs.NewHasher(),
		parent:  nil,

		locker:    sync.Mutex{},
		itemArray: make([]*chunkItem, 0),
		itemMap:   make(map[string]int),
	}
	err := g.load(name)
	if err != nil {
		return nil, err
	}
	return g, nil
}

func (f *chunkGroup) Parent() core.Group {
	return f.parent
}

func (f *chunkGroup) Group(name string) (core.Group, error) {
	f.locker.Lock()
	defer f.locker.Unlock()

	if !f.Exists(name) {
		return nil, util.ErrorOfNonexists("group", name)
	}

	item, err := f.get(name)
	if err != nil {
		return nil, err
	}
	if !item.group {
		return nil, util.ErrorOfInvalid("group", name)
	}
	g := &chunkGroup{
		storage: f.storage,
		name:    name,
		hash:    cs.NewHasher(),
		parent:  f,

		locker:    sync.Mutex{},
		itemArray: make([]*chunkItem, 0),
		itemMap:   make(map[string]int),
	}
	err = g.Load(item.key)
	if err != nil {
		return nil, err
	}
	return g, nil
}

func (f *chunkGroup) AddGroup(name string) (core.Group, error) {
	f.locker.Lock()
	defer f.locker.Unlock()

	if f.Exists(name) {
		return nil, util.ErrorOf("Already exists", "group", name)
	}
	g := &chunkGroup{
		storage: f.storage,
		name:    name,
		hash:    cs.NewHasher(),
		parent:  f,

		locker:    sync.Mutex{},
		itemArray: make([]*chunkItem, 0),
		itemMap:   make(map[string]int),
	}
	err := g.update()
	if err != nil {
		return nil, err
	}
	item := &chunkItem{
		name:  name,
		key:   *g.itemKey,
		group: true,

		g: g,
	}
	f.itemMap[name] = len(f.itemArray)
	// f.itemMap[item.key.String()] = len(f.itemArray)
	f.itemArray = append(f.itemArray, item)
	err = f.update()
	if err != nil {
		return nil, err
	}

	return g, nil
}

func (f *chunkGroup) Dump(log core.Printer) {
	log.Printf("name: %s, size: %d, parent: %s", f.name, len(f.itemMap), func(d *chunkGroup) string {
		if d != nil {
			return d.name
		}
		return "nil"
	}(f.parent))

	for i := 0; i < len(f.itemArray); i++ {
		item := f.itemArray[i]
		if item != nil {
			log.Printf("%d: %s, %s, %v", i, item.name, item.key.String(), item.group)
		}
	}
}

func (f *chunkGroup) dispose(name string, key core.Key) {
	for i := 0; i < len(f.itemArray); i++ {
		item := f.itemArray[i]
		if item != nil {
			if item.group {
				item.g.Dispose()
			}
			err := f.storage.releaseL(item.key)
			if err != nil {
				glog.Error(err)
			}
		}
	}
	err := f.storage.releaseL(key)
	if err != nil {
		glog.Error(err)
	}

	f.itemArray = make([]*chunkItem, 0)
	f.itemMap = make(map[string]int)

	f.itemKey = nil
	f.itemEntry = nil
}

func (f *chunkGroup) Dispose() {
	if f.parent != nil {
		if f.itemKey != nil {
			f.dispose(f.name, *f.itemKey)
		}
	} else {
		link := core.GetKey(f.name)
		linkEntry, ok := f.storage.getEntry(link)
		if ok {
			key, err := core.ParseKey(linkEntry.name)
			if err != nil {
				glog.Error(err)
			} else {
				f.dispose(linkEntry.name, key)
			}

			err = f.storage.releaseL(link)
			if err != nil {
				glog.Error(err)
			}
		}
	}
}

func (f *chunkGroup) Load(key core.Key) error {
	entry, ok := f.storage.getEntry(key)
	if ok {
		f.name = entry.name
		f.itemKey = &key
		f.itemEntry = entry

		for i := 0; i < len(entry.chunks); i++ {
			chunk := entry.chunks[i]
			item := &chunkItem{}
			if i == 0 {
				err := item.UnmarshalBinary(entry.data[:chunk.nextPos])
				if err != nil {
					return err
				}
			} else {
				err := item.UnmarshalBinary(entry.data[entry.chunks[i-1].nextPos:chunk.nextPos])
				if err != nil {
					return err
				}
			}
			f.itemMap[item.name] = len(f.itemArray)
			// f.itemMap[item.key.String()] = len(f.itemArray)

			f.itemArray = append(f.itemArray, item)
		}
		return nil
	} else {
		return util.ErrorOfNonexists("key", key.String())
	}
}

func (f *chunkGroup) get(name string) (*chunkItem, error) {
	index, ok := f.itemMap[name]
	if !ok {
		return nil, util.ErrorOfNonexists("dir", name)
	}
	item := f.itemArray[index]
	if item == nil {
		return nil, util.ErrorOfInvalid("state", fmt.Sprintf("%d", index))
	}
	return item, nil
}

func (f *chunkGroup) exists(name string) (core.Key, bool, error) {
	link := core.GetKey(name)
	_, ok := f.storage.getEntry(link)
	if ok {
		return link, ok, nil
	} else {
		return nil, false, nil
	}
}

func (f *chunkGroup) load(name string) error {
	link, ok, err := f.exists(name)
	if err != nil {
		return err
	}
	if ok {
		linkEntry, ok := f.storage.getEntry(link)
		if ok {
			key, err := core.ParseKey(linkEntry.name)
			if err != nil {
				return err
			}
			err = f.Load(key)
			if err != nil {
				return err
			}
		} else {
			return util.ErrorOfNonexists("dir", name)
		}
	}

	err = f.update()
	if err != nil {
		return err
	}

	return nil
}

func (f *chunkGroup) update() error {
	lm := len(f.itemMap)
	la := len(f.itemArray)
	if la > 32 && (la+1) > 2*(lm+1) {
		itemMap := make(map[string]int)
		itemArray := make([]*chunkItem, lm)
		c := 0
		for i := 0; i < la; i++ {
			item := f.itemArray[i]
			if item != nil {
				itemArray[c] = item
				itemMap[item.name] = c
				c++
			}
		}
		f.itemMap = itemMap
		f.itemArray = itemArray
	}

	buf := new(bytes.Buffer)
	chunks := make([]chunkRef, 0)
	for i := 0; i < len(f.itemArray); i++ {
		item := f.itemArray[i]
		if item != nil {
			data, err := item.MarshalBinary()
			if err != nil {
				return err
			}
			_, err = buf.Write(data)
			if err != nil {
				return err
			}
			chunks = append(chunks, chunkRef{
				key:     item.key,
				nextPos: int64(buf.Len()),
			})
		}
	}
	entry := &chunkEntry{
		name:   f.name,
		data:   buf.Bytes(),
		chunks: chunks,
		refs:   1,
	}
	data, err := entry.MarshalBinary()
	if err != nil {
		return err
	}
	_, err = f.hash.Write(data)
	if err != nil {
		return err
	}
	key := core.Key(f.hash.Sum(nil))
	f.hash.Reset()

	f.itemKey = &key
	f.itemEntry = entry

	return nil
}

func (f *chunkGroup) Name() string {
	return f.name
}

func (f *chunkGroup) Key() core.Key {
	if f.itemKey != nil {
		return *f.itemKey
	}
	return core.Key{}
}

func (f *chunkGroup) AddData(name string, key core.Key) (bool, error) {
	f.locker.Lock()
	defer f.locker.Unlock()

	var item *chunkItem = nil

	index, ok := f.itemMap[name]
	if ok {
		item = f.itemArray[index]
		if item != nil {
			if item.key.String() == key.String() {
				return true, nil
			} else {
				item.key = key
			}
		}
	}
	err := f.storage.lockL(key)
	if err != nil {
		return false, err
	}
	if item == nil {
		item = &chunkItem{
			name:  name,
			key:   key,
			group: false,
		}
		f.itemMap[name] = len(f.itemArray)
		// f.itemMap[item.key.String()] = len(f.itemArray)
		f.itemArray = append(f.itemArray, item)
	}
	err = f.update()
	if err != nil {
		return false, err
	}

	return false, nil
}

func (f *chunkGroup) Remove(name string) error {
	f.locker.Lock()
	defer f.locker.Unlock()

	index, ok := f.itemMap[name]
	if ok {
		item := f.itemArray[index]
		if item != nil {
			if item.group {
				item.g.Dispose()
			} else {
				err := f.storage.releaseL(item.key)
				if err != nil {
					return err
				}
			}
		}

		delete(f.itemMap, name)
		f.itemArray[index] = nil

		err := f.update()
		if err != nil {
			return err
		}
	}

	return nil
}

func (f *chunkGroup) Exists(name string) bool {
	_, ok := f.itemMap[name]
	return ok
}

func (f *chunkGroup) GetKey(name string) (core.Key, error) {
	f.locker.Lock()
	defer f.locker.Unlock()

	index, ok := f.itemMap[name]
	if ok {
		item := f.itemArray[index]
		if item != nil {
			return item.key, nil
		} else {
			return nil, util.ErrorOfInvalid("state", fmt.Sprintf("%d", index))
		}
	}
	return nil, util.ErrorOfNonexists("key", name)
}

func (f *chunkGroup) IsGroup(name string) (bool, error) {
	f.locker.Lock()
	defer f.locker.Unlock()

	index, ok := f.itemMap[name]
	if ok {
		item := f.itemArray[index]
		if item != nil {
			return item.group, nil
		} else {
			return false, util.ErrorOfInvalid("state", fmt.Sprintf("%d", index))
		}
	}
	return false, util.ErrorOfNonexists("dir", name)
}

func (f *chunkGroup) List() []string {
	f.locker.Lock()
	defer f.locker.Unlock()

	list := make([]string, 0)
	for i := 0; i < len(f.itemArray); i++ {
		item := f.itemArray[i]
		if item != nil {
			list = append(list, item.name)
		}
	}
	return list
}

func (f *chunkGroup) Commit() error {
	f.locker.Lock()
	defer f.locker.Unlock()

	if f.itemKey != nil && f.itemEntry != nil {
		if _, _, err := f.storage.putEntry(*f.itemKey, f.itemEntry); err != nil {
			return err
		}

		if f.parent != nil {
			item, err := f.parent.get(f.name)
			if err != nil {
				return err
			}
			item.key = *f.itemKey
			err = f.parent.update()
			if err != nil {
				return err
			}
			err = f.parent.Commit()
			if err != nil {
				return err
			}
		} else {
			err := f.storage.add(f.name, *f.itemKey)
			if err != nil {
				return err
			}
			err = f.storage.flush()
			if err != nil {
				return err
			}
			err = f.storage.entries.Flush()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

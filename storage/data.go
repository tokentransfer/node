package storage

import (
	"io"

	"github.com/tokentransfer/node/core"
	"github.com/tokentransfer/node/util"
)

type chunkData struct {
	storage  *chunkStorage
	key      core.Key
	disposed bool
}

func (f *chunkData) Key() core.Key {
	return f.key
}

func (f *chunkData) Entry() *chunkEntry {
	entry, _ := f.storage.getEntry(f.key)
	return entry
}

func (f *chunkData) Open() io.ReadCloser {
	entry := f.Entry()
	if len(entry.chunks) > 0 {
		f.storage.lockL(f.key)
		return &chunkReader{
			storage:    f.storage,
			key:        f.key,
			chunksTail: entry.chunks,
			closed:     false,
		}
	} else {
		return &chunkDataReader{entry.data, 0}
	}
}

func (f *chunkData) Size() int64 {
	entry := f.Entry()
	if len(entry.chunks) > 0 {
		return entry.chunks[len(entry.chunks)-1].nextPos
	}
	return int64(len(entry.data))
}

func (f *chunkData) Dispose() {
	if !f.disposed {
		f.disposed = true
		f.storage.releaseL(f.key)
	}
}

func (f *chunkData) Remove() {
	if !f.disposed {
		f.disposed = true

		f.storage.releaseL(f.key)
		entry := f.Entry()
		for _, chunk := range entry.chunks {
			f.storage.releaseL(chunk.key)
		}
	}
}

func (f *chunkData) checkValid() error {
	if f.disposed {
		return util.ErrorOfInvalid("chunk", "already disposed")
	}
	return nil
}

func (f *chunkData) Duplicate() (core.Data, error) {
	if err := f.checkValid(); err != nil {
		return nil, err
	}
	file, err := f.storage.Get(f.key)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (f *chunkData) IsChunked() (bool, error) {
	if err := f.checkValid(); err != nil {
		return false, err
	}

	entry := f.Entry()
	return len(entry.chunks) > 0, nil
}

func (f *chunkData) Chunks() core.DataIterator {
	var chunks []chunkRef
	entry := f.Entry()
	if len(entry.chunks) > 0 {
		chunks = entry.chunks
	} else {
		chunks = make([]chunkRef, 1)
		chunks[0] = chunkRef{f.key, f.Size()}
	}
	f.storage.lockL(f.key)
	return &chunksIterator{
		storage:      f.storage,
		key:          f.key,
		chunks:       chunks,
		chunkIdx:     0,
		lastChunkIdx: -1,
		disposed:     false,
	}
}

func (f *chunkData) NumChunks() int64 {
	entry := f.Entry()
	if len(entry.chunks) > 0 {
		return int64(len(entry.chunks))
	}
	return 1
}

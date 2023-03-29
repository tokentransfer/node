package chunk

import (
	"fmt"
	"io"

	"github.com/tokentransfer/node/core"
	"github.com/tokentransfer/node/util"
)

type chunkRef struct {
	key     core.Key
	nextPos int64
}

type chunkDataReader struct {
	data  []byte
	index int
}

type chunkReader struct {
	storage    *chunkStorage
	key        core.Key
	chunksTail []chunkRef
	closed     bool
	dataReader io.ReadCloser
}

type chunksIterator struct {
	storage      *chunkStorage
	key          core.Key
	chunks       []chunkRef
	chunkIdx     int
	lastChunkIdx int
	disposed     bool
}

func (ci *chunksIterator) checkValid() error {
	if ci.disposed {
		return util.ErrorOfInvalid("chunk", "already disposed")
	}
	return nil
}

func (ci *chunksIterator) Dispose() error {
	if !ci.disposed {
		ci.disposed = true
		if err := ci.storage.releaseL(ci.key); err != nil {
			return err
		}
	}
	return nil
}

func (ci *chunksIterator) Duplicate() (core.DataIterator, error) {
	if err := ci.checkValid(); err != nil {
		return nil, err
	}
	ci.storage.lockL(ci.key)
	return &chunksIterator{
		storage:  ci.storage,
		key:      ci.key,
		chunks:   ci.chunks,
		chunkIdx: ci.chunkIdx,
		disposed: false,
	}, nil
}

func (ci *chunksIterator) Next() (bool, error) {
	if err := ci.checkValid(); err != nil {
		return false, err
	}
	if ci.chunkIdx == len(ci.chunks) {
		ci.Dispose()
		return false, nil
	} else {
		ci.lastChunkIdx = ci.chunkIdx
		ci.chunkIdx++
		return true, nil
	}
}

func (ci *chunksIterator) Key() (core.Key, error) {
	if err := ci.checkValid(); err != nil {
		return nil, err
	}
	return ci.chunks[ci.lastChunkIdx].key, nil
}

func (ci *chunksIterator) Size() (int64, error) {
	if err := ci.checkValid(); err != nil {
		return 0, err
	}
	startPos := int64(0)
	if ci.lastChunkIdx > 0 {
		startPos = ci.chunks[ci.lastChunkIdx-1].nextPos
	}
	return ci.chunks[ci.lastChunkIdx].nextPos - startPos, nil
}

func (ci *chunksIterator) Data() (core.Data, error) {
	err := ci.checkValid()
	if err != nil {
		return nil, err
	}
	if f, err := ci.storage.Get(ci.chunks[ci.lastChunkIdx].key); err != nil {
		return nil, err
	} else {
		return f, nil
	}
}

func (r *chunkDataReader) Read(b []byte) (n int, err error) {
	if len(b) == 0 {
		return 0, nil
	}
	if r.index >= len(r.data) {
		return 0, io.EOF
	}
	n = copy(b, r.data[r.index:])
	r.index += n
	return
}

func (r *chunkDataReader) Close() error {
	return nil
}

func (r *chunkReader) Read(b []byte) (n int, err error) {
	if r.closed {
		err = util.ErrorOfInvalid("state", fmt.Sprintf("%v", r.closed))
		return
	}
	if n == 0 && err == nil {
		if r.dataReader == nil {
			if len(r.chunksTail) > 0 {
				if data, e := r.storage.Chunk(r.chunksTail[0].key); e != nil {
					return 0, err
				} else {
					r.dataReader = &chunkDataReader{data, 0}
					r.chunksTail = r.chunksTail[1:]
				}
			} else {
				return 0, io.EOF
			}
		}

		n, err = r.dataReader.Read(b)
		if err == io.EOF {
			err = r.dataReader.Close()
			r.dataReader = nil
		}
	}
	return
}

func (r *chunkReader) Close() (err error) {
	if r.closed {
		return nil
	}
	r.closed = true

	r.storage.releaseL(r.key)
	if r.dataReader != nil {
		err = r.dataReader.Close()
		r.dataReader = nil
	}
	return
}

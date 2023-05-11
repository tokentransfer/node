package storage

import (
	"bytes"
	"fmt"
	"hash"

	"github.com/tokentransfer/node/core"
	"github.com/tokentransfer/node/util"
)

type chunkTemporary struct {
	storage   *chunkStorage
	info      string
	buffer    bytes.Buffer
	fileHash  hash.Hash
	chunkHash hash.Hash
	valid     bool
	open      bool
	chunker   Chunker
	chunks    []chunkRef
}

func (t *chunkTemporary) flushBufferIntoChunk() error {
	if t.buffer.Len() == 0 {
		return nil
	}

	chunkInfo := fmt.Sprintf("%v #%d", t.info, len(t.chunks))
	chunkData := make([]byte, t.buffer.Len())
	copy(chunkData, t.buffer.Bytes())

	key := core.Key(t.chunkHash.Sum(nil))
	t.chunkHash.Reset()

	if err := t.storage.storeEntry(key, chunkData, nil, chunkInfo); err != nil {
		return err
	}
	chunk := chunkRef{
		key:     key,
		nextPos: int64(t.buffer.Len()),
	}
	if len(t.chunks) > 0 {
		chunk.nextPos += t.chunks[len(t.chunks)-1].nextPos
	}
	t.chunks = append(t.chunks, chunk)

	t.buffer.Reset()
	return nil
}

func (t *chunkTemporary) Write(b []byte) (int, error) {
	if !t.valid || !t.open {
		return 0, util.ErrorOfInvalid("state", t.info)
	}
	t.valid = false

	nBytes := len(b)

	for len(b) > 0 {
		nBoundary := t.chunker.Scan(b)
		if _, err := t.buffer.Write(b[:nBoundary]); err != nil {
			return 0, err
		}
		t.chunkHash.Write(b[:nBoundary])
		t.fileHash.Write(b[:nBoundary])
		if nBoundary < len(b) {
			if err := t.flushBufferIntoChunk(); err != nil {
				return 0, err
			}
			b = b[nBoundary:]
		} else {
			b = nil
		}
	}

	t.valid = true
	return nBytes, nil
}

func (t *chunkTemporary) Close() error {
	if !t.valid || !t.open {
		return util.ErrorOfInvalid("state", t.info)
	}
	t.open = false
	t.valid = false
	key := core.Key(t.fileHash.Sum(nil))

	if len(t.chunks) == 0 {
		data := make([]byte, t.buffer.Len())
		copy(data, t.buffer.Bytes())
		if err := t.storage.storeEntry(key, data, nil, t.info); err != nil {
			return err
		}
	} else {
		if err := t.flushBufferIntoChunk(); err != nil {
			return err
		}
		finalChunks := make([]chunkRef, len(t.chunks))
		copy(finalChunks, t.chunks)
		if err := t.storage.storeEntry(key, nil, finalChunks, t.info); err != nil {
			return err
		}
	}
	t.valid = true
	return nil
}

func (t *chunkTemporary) Data() (core.Data, error) {
	if !t.valid {
		return nil, util.ErrorOfInvalid("temporary state", t.info)
	}
	if t.open {
		return nil, util.ErrorOf("still open", "temporary", t.info)
	}

	key := core.Key(t.fileHash.Sum(nil))

	file, err := t.storage.Get(key)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (t *chunkTemporary) Dispose() {
	if t.chunks == nil {
		return
	}

	t.releaseFromStorage()

	t.valid = false
	t.open = false
	t.buffer = bytes.Buffer{}
	t.chunker = nil
	t.chunks = nil
}

func (t *chunkTemporary) releaseFromStorage() {
	t.storage.mutex.Lock()
	defer t.storage.mutex.Unlock()

	if !t.open && t.valid {
		key := core.Key(t.fileHash.Sum(nil))
		t.storage.release(key)
	} else {
		for _, chunk := range t.chunks {
			t.storage.release(chunk.key)
		}
	}
}

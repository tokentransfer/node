package chunk

import (
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagcbor"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/fluent/qp"
	"github.com/ipld/go-ipld-prime/node/basicnode"
	"github.com/tokentransfer/node/core"
	"github.com/tokentransfer/node/util"
)

const (
	entrySize = 112
	chunkSize = 40
)

type chunkEntry struct {
	name   string
	data   []byte
	chunks []chunkRef
	refs   int
}

func (e *chunkEntry) storageSize() int64 {
	return int64(entrySize + len(e.data) + chunkSize*len(e.chunks))
}

func (e *chunkEntry) MarshalBinary() ([]byte, error) {
	node, err := qp.BuildMap(basicnode.Prototype.Any, -1, func(ma datamodel.MapAssembler) {
		qp.MapEntry(ma, "name", qp.String(e.name))
		qp.MapEntry(ma, "data", qp.Bytes(e.data))
		qp.MapEntry(ma, "chunks", qp.List(int64(len(e.chunks)), func(la datamodel.ListAssembler) {
			for _, c := range e.chunks {
				key := c.key
				pos := c.nextPos
				qp.ListEntry(la, qp.Map(2, func(mi datamodel.MapAssembler) {
					qp.MapEntry(mi, "key", qp.Bytes(key[:]))
					qp.MapEntry(mi, "nextPos", qp.Int(pos))
				}))
			}
		}))
		qp.MapEntry(ma, "refs", qp.Int(int64(e.refs)))
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

func (e *chunkEntry) UnmarshalBinary(data []byte) error {
	node, err := ipld.Decode(data, dagcbor.Decode)
	if err != nil {
		return err
	}
	e.name = util.GetStringFromNode(node, "name")
	if data, err := util.GetBytesFromNode(node, "data"); err != nil {
		return err
	} else {
		e.data = data
	}
	list := util.GetListFromNode(node, "chunks")
	if list != nil {
		chunks := make([]chunkRef, 0)
		for !list.Done() {
			c := chunkRef{}
			_, item, err := list.Next()
			if err != nil {
				return err
			}
			k, err := util.GetBytesFromNode(item, "key")
			if err != nil {
				return err
			}
			c.key = core.Key(k)
			c.nextPos = util.GetIntFromNode(item, "nextPos")

			chunks = append(chunks, c)
		}
		e.chunks = chunks
	}
	e.refs = int(util.GetIntFromNode(node, "refs"))

	return nil
}

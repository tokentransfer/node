package consensus

import (
	"time"

	libaccount "github.com/tokentransfer/interfaces/account"
	libblock "github.com/tokentransfer/interfaces/block"
	libcore "github.com/tokentransfer/interfaces/core"
	"github.com/tokentransfer/node/util"
)

type Status int

const (
	PeerNone Status = iota
	PeerConnected
	PeerKnown
	PeerConsensused
)

func (s Status) String() string {
	switch s {
	case PeerNone:
		return "none"
	case PeerConnected:
		return "connected"
	case PeerKnown:
		return "known"
	case PeerConsensused:
		return "consensused"
	}
	return "unknown status"
}

type peerEntry struct {
	BlockNumber int64
	PeerCount   int64
	Status      Status

	lastSendBlock int64
	lastSendTime  time.Time
}

type Peer struct {
	Id  string
	Key libaccount.PublicKey

	address string
	index   uint64
	peermap map[string]*peerEntry
}

func (p *Peer) GetPublicKey() libaccount.PublicKey {
	return p.Key
}

func (p *Peer) getPeerEntry(rootAccount libcore.Address) *peerEntry {
	root := util.GetString(rootAccount)
	entry, ok := p.peermap[root]
	if !ok {
		entry = &peerEntry{
			lastSendBlock: int64(-1),
			lastSendTime:  time.Now(),
		}
		p.peermap[root] = entry
	}
	return entry
}

func (p *Peer) GetIndex(n *Node) uint64 {
	if p.index == 0 {
		address, err := p.Key.GenerateAddress()
		if err != nil {
			return 0
		}
		index, err := n.GetIndex(address.String())
		if err != nil {
			return 0
		}

		p.address = address.String()
		p.index = index
	}
	return p.index
}

func (p *Peer) GetAddress(n *Node) string {
	if len(p.address) == 0 {
		address, err := p.Key.GenerateAddress()
		if err != nil {
			return ""
		}
		index, err := n.GetIndex(address.String())
		if err != nil {
			return ""
		}

		p.address = address.String()
		p.index = index
	}
	return p.address
}

type Entry struct {
	Consensused    bool
	ValidatedBlock libblock.Block

	txlist []libblock.TransactionWithData
	txmap  map[string]libblock.TransactionWithData
}

func (entry *Entry) GetBlockNumber() int64 {
	if entry.ValidatedBlock != nil {
		return int64(entry.ValidatedBlock.GetIndex())
	}
	return -1
}

func (entry *Entry) GetBlockHash() string {
	if entry.ValidatedBlock != nil {
		return entry.ValidatedBlock.GetHash().String()
	}
	return ""
}

func (entry *Entry) GetBlock() libblock.Block {
	if entry.ValidatedBlock != nil {
		return entry.ValidatedBlock
	}
	return nil
}

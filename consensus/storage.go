package consensus

import (
	"sync"

	libaccount "github.com/tokentransfer/interfaces/account"
	libcore "github.com/tokentransfer/interfaces/core"
	libcrypto "github.com/tokentransfer/interfaces/crypto"
	libstore "github.com/tokentransfer/interfaces/store"
	"github.com/tokentransfer/node/chunk"
	"github.com/tokentransfer/node/config"
	"github.com/tokentransfer/node/merkle"
	"github.com/tokentransfer/node/util"
)

type StorageService struct {
	config *config.Config
	cs     libcrypto.CryptoService
	as     libaccount.AccountService

	merkleMap    map[string]libstore.MerkleService
	merkleLocker *sync.Mutex

	chunkMap    map[string]*chunk.ChunkService
	chunkLocker *sync.Mutex
}

func NewStorageService(config *config.Config, cs libcrypto.CryptoService, as libaccount.AccountService) (*StorageService, error) {
	return &StorageService{
		config: config,
		cs:     cs,
		as:     as,

		merkleMap:    make(map[string]libstore.MerkleService),
		merkleLocker: new(sync.Mutex),
		chunkMap:     make(map[string]*chunk.ChunkService),
		chunkLocker:  new(sync.Mutex),
	}, nil
}

func (service *StorageService) GetMerkleService(rootAccount libcore.Address) libstore.MerkleService {
	service.merkleLocker.Lock()
	defer service.merkleLocker.Unlock()

	root := util.GetString(rootAccount)
	ms, ok := service.merkleMap[root]
	if !ok {
		s, err := merkle.NewMerkleService(service.config, service.cs, rootAccount)
		if err != nil {
			panic(err)
		}
		service.merkleMap[root] = s

		ms = s
	}
	return ms
}

func (service *StorageService) GetChunkService(rootAccount libcore.Address) *chunk.ChunkService {
	service.chunkLocker.Lock()
	defer service.chunkLocker.Unlock()

	root := util.GetString(rootAccount)
	ss, ok := service.chunkMap[root]
	if !ok {
		s, err := chunk.NewChunkService(service.config, rootAccount)
		if err != nil {
			panic(err)
		}
		service.chunkMap[root] = s

		ss = s
	}
	return ss
}

func (service *StorageService) Close() error {
	for _, s := range service.merkleMap {
		if err := s.Close(); err != nil {
			return err
		}
	}
	for _, s := range service.chunkMap {
		if err := s.Close(); err != nil {
			return err
		}
	}
	return nil
}

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
)

type StorageService struct {
	config *config.Config
	cs     libcrypto.CryptoService
	as     libaccount.AccountService

	merkleMap    map[libcore.Address]libstore.MerkleService
	merkleLocker *sync.Mutex

	chunkMap    map[libcore.Address]*chunk.ChunkService
	chunkLocker *sync.Mutex
}

func NewStorageService(config *config.Config, cs libcrypto.CryptoService, as libaccount.AccountService) (*StorageService, error) {
	return &StorageService{
		config: config,
		cs:     cs,
		as:     as,

		merkleMap:    make(map[libcore.Address]libstore.MerkleService),
		merkleLocker: new(sync.Mutex),
		chunkMap:     make(map[libcore.Address]*chunk.ChunkService),
		chunkLocker:  new(sync.Mutex),
	}, nil
}

func (service *StorageService) GetMerkleService(rootAccount libcore.Address) libstore.MerkleService {
	service.merkleLocker.Lock()
	defer service.merkleLocker.Unlock()

	ms, ok := service.merkleMap[rootAccount]
	if !ok {
		s, err := merkle.NewMerkleService(service.config, service.cs, rootAccount)
		if err != nil {
			panic(err)
		}
		service.merkleMap[rootAccount] = s

		ms = s
	}
	return ms
}

func (service *StorageService) GetChunkService(rootAccount libcore.Address) *chunk.ChunkService {
	service.chunkLocker.Lock()
	defer service.chunkLocker.Unlock()

	ss, ok := service.chunkMap[rootAccount]
	if !ok {
		s, err := chunk.NewChunkService(service.config, rootAccount)
		if err != nil {
			panic(err)
		}
		service.chunkMap[rootAccount] = s

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

package consensus

import (
	"bytes"
	"fmt"
	"io"

	libcore "github.com/tokentransfer/interfaces/core"
	"github.com/tokentransfer/node/chunk"
	"github.com/tokentransfer/node/core"
	"github.com/tokentransfer/node/store"
	"github.com/tokentransfer/node/vm"
)

type storageService struct {
	storage core.Storage
	stackdb *store.StackService
}

func NewStorageService(c libcore.Config) (*storageService, error) {
	service := &storageService{}
	err := service.Init(c)
	if err != nil {
		return nil, err
	}
	err = service.Start()
	if err != nil {
		return nil, err
	}
	return service, nil
}

func (s *storageService) CreateSandbox() error {
	memdb := &store.MemoryService{
		Name: "memory",
	}
	err := memdb.Init(nil)
	if err != nil {
		return err
	}
	err = memdb.Start()
	if err != nil {
		return err
	}
	err = s.stackdb.Push(memdb)
	if err != nil {
		return err
	}
	return nil
}

func (s *storageService) CommitSandbox() error {
	err := s.stackdb.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (s *storageService) CancelSandbox() error {
	err := s.stackdb.Cancel()
	if err != nil {
		return err
	}
	return nil
}

func getGroup(parent core.Group, name string) (core.Group, error) {
	g, _ := parent.Group(name)
	if g == nil {
		group, err := parent.AddGroup(name)
		if err != nil {
			return nil, err
		}
		err = group.Commit()
		if err != nil {
			return nil, err
		}
		g = group
	}
	return g, nil
}

func (s *storageService) CreateContract(account libcore.Address, code []byte) (libcore.Hash, libcore.Hash, error) {
	rootGroup, err := s.storage.Group("/")
	if err != nil {
		return nil, nil, err
	}
	codeGroup, err := getGroup(rootGroup, "code")
	if err != nil {
		return nil, nil, err
	}

	address := account.String()
	t := s.storage.Create(address)
	_, err = t.Write(code)
	if err != nil {
		return nil, nil, err
	}
	err = t.Close()
	if err != nil {
		return nil, nil, err
	}

	d := t.Data()
	_, err = codeGroup.AddData(address, d.Key())
	if err != nil {
		return nil, nil, err
	}

	err = codeGroup.Commit()
	if err != nil {
		return nil, nil, err
	}
	d.Dispose()
	t.Dispose()

	return libcore.Hash(s.storage.Root()), libcore.Hash(d.Key()), nil
}

func (s *storageService) CreateData(account libcore.Address, data []byte) (libcore.Hash, libcore.Hash, error) {
	rootGroup, err := s.storage.Group("/")
	if err != nil {
		return nil, nil, err
	}
	dataGroup, err := getGroup(rootGroup, "data")
	if err != nil {
		return nil, nil, err
	}

	address := account.String()
	t := s.storage.Create(address)
	_, err = t.Write(data)
	if err != nil {
		return nil, nil, err
	}
	err = t.Close()
	if err != nil {
		return nil, nil, err
	}

	d := t.Data()
	_, err = dataGroup.AddData(address, d.Key())
	if err != nil {
		return nil, nil, err
	}

	err = dataGroup.Commit()
	if err != nil {
		return nil, nil, err
	}
	d.Dispose()
	t.Dispose()

	return libcore.Hash(s.storage.Root()), libcore.Hash(d.Key()), nil
}

func (s *storageService) RunContract(cost int64, codeAccount libcore.Address, dataAccount libcore.Address, method string, params [][]byte) (libcore.Hash, libcore.Hash, error) {
	rootGroup, err := s.storage.Group("/")
	if err != nil {
		return nil, nil, err
	}
	codeGroup, err := getGroup(rootGroup, "code")
	if err != nil {
		return nil, nil, err
	}

	address := codeAccount.String()
	codeKey, err := codeGroup.GetKey(address)
	if err != nil {
		return nil, nil, err
	}
	codeData, err := s.storage.Get(codeKey)
	if err != nil {
		return nil, nil, err
	}
	codeReader := codeData.Open()
	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, codeReader); err != nil {
		return nil, nil, err
	}
	codeData.Dispose()
	codeReader.Close()

	_, result, err := vm.RunWasm(cost, buf.Bytes(), method, params) // remainCost
	if err != nil {
		return nil, nil, err
	}
	rootHash, dataHash, err := s.CreateData(dataAccount, result)
	if err != nil {
		return nil, nil, err
	}
	fmt.Println("> run contract", codeAccount.String(), dataAccount.String(), method, string(result))
	return rootHash, dataHash, nil
}

func (s *storageService) Init(c libcore.Config) error {
	datadb := &store.LevelService{Name: "data"}
	err := datadb.Init(c)
	if err != nil {
		return err
	}
	err = datadb.Start()
	if err != nil {
		return err
	}
	stackdb, err := store.CreateStackService(datadb)
	if err != nil {
		return err
	}
	storage, err := chunk.LoadStorageWith(stackdb, 256*1024*1024)
	if err != nil {
		return err
	}
	s.storage = storage
	s.stackdb = stackdb
	return nil
}

func (s *storageService) Start() error {
	return nil
}

func (s *storageService) Close() error {
	err := s.storage.Close()
	if err != nil {
		return err
	}
	return nil
}

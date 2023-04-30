package consensus

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"sync"

	"github.com/caivega/glog"
	libcore "github.com/tokentransfer/interfaces/core"
	libcrypto "github.com/tokentransfer/interfaces/crypto"
	libstore "github.com/tokentransfer/interfaces/store"
	"github.com/tokentransfer/node/chunk"
	"github.com/tokentransfer/node/core"
	"github.com/tokentransfer/node/core/pb"
	"github.com/tokentransfer/node/store"
	"github.com/tokentransfer/node/util"
	"github.com/tokentransfer/node/vm"
)

type StorageService struct {
	storage core.Storage
	stackdb *store.StackService

	locker sync.Mutex
}

func NewStorageService(c libcore.Config) (*StorageService, error) {
	service := &StorageService{
		locker: sync.Mutex{},
	}
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

func (s *StorageService) Dump(printer core.Printer) {
	s.locker.Lock()
	defer s.locker.Unlock()

	s.storage.DumpLog(printer)
}

func (s *StorageService) CreateSandbox() error {
	s.locker.Lock()
	defer s.locker.Unlock()

	// s.dump("before create")
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
	// s.dump("after create")
	return nil
}

func (s *StorageService) CommitSandbox() error {
	s.locker.Lock()
	defer s.locker.Unlock()

	// s.dump("before commit")
	_, err := s.stackdb.Commit()
	if err != nil {
		return err
	}
	if err := s.stackdb.Flush(); err != nil {
		return err
	}
	if err := s.storage.Update(); err != nil {
		return err
	}
	// s.dumpWith("after commit", popKv)
	return nil
}

func (s *StorageService) CancelSandbox() error {
	s.locker.Lock()
	defer s.locker.Unlock()

	// s.dump("before cancel")
	_, err := s.stackdb.Cancel()
	if err != nil {
		return err
	}
	if err := s.storage.Update(); err != nil {
		return err
	}
	// s.dumpWith("after cancel", popKv)
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

func (s *StorageService) ReadCode(codeAccount libcore.Address) ([]byte, []byte, error) {
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
	wasmCode, err := util.ReadBytes(codeReader)
	if err != nil {
		return nil, nil, err
	}
	abiCode, err := util.ReadBytes(codeReader)
	if err != nil {
		return nil, nil, err
	}
	codeReader.Close()
	codeData.Dispose()
	glog.Infoln("> read code", codeAccount.String(), codeKey.String(), len(wasmCode), len(abiCode))

	return wasmCode, abiCode, nil
}

func (s *StorageService) WriteCode(account libcore.Address, wasmCode []byte, abiCode []byte) (core.Key, error) {
	rootGroup, err := s.storage.Group("/")
	if err != nil {
		return nil, err
	}
	codeGroup, err := getGroup(rootGroup, "code")
	if err != nil {
		return nil, err
	}
	// s.dump("create code group")

	address := account.String()
	t := s.storage.Create(address)
	err = util.WriteBytes(t, wasmCode)
	if err != nil {
		return nil, err
	}
	if abiCode == nil {
		abiCode = make([]byte, 0)
	}
	err = util.WriteBytes(t, abiCode)
	if err != nil {
		return nil, err
	}
	err = t.Close()
	if err != nil {
		return nil, err
	}

	d, err := t.Data()
	if err != nil {
		return nil, err
	}
	_, err = codeGroup.AddData(address, d.Key())
	if err != nil {
		return nil, err
	}
	err = codeGroup.Commit()
	if err != nil {
		return nil, err
	}
	// s.dump("create code account")
	d.Dispose()
	t.Dispose()
	glog.Infoln("> write code", address, d.Key().String(), len(wasmCode), len(abiCode))

	return d.Key(), nil
}

func (s *StorageService) WriteData(dataAccount libcore.Address, codeAccount libcore.Address, data []byte) (libcore.Hash, libcore.Hash, error) {
	rootGroup, err := s.storage.Group("/")
	if err != nil {
		return nil, nil, err
	}
	dataGroup, err := getGroup(rootGroup, "data")
	if err != nil {
		return nil, nil, err
	}
	accountGroup, err := getGroup(dataGroup, dataAccount.String())
	if err != nil {
		return nil, nil, err
	}

	address := codeAccount.String()
	t := s.storage.Create(address)
	_, err = t.Write(data)
	if err != nil {
		return nil, nil, err
	}
	err = t.Close()
	if err != nil {
		return nil, nil, err
	}

	d, err := t.Data()
	if err != nil {
		return nil, nil, err
	}
	_, err = accountGroup.AddData(address, d.Key())
	if err != nil {
		return nil, nil, err
	}

	err = accountGroup.Commit()
	if err != nil {
		return nil, nil, err
	}
	d.Dispose()
	t.Dispose()
	glog.Infoln("> write data", dataAccount.String(), address, d.Key().String(), len(data))

	return libcore.Hash(s.storage.Root()), libcore.Hash(d.Key()), nil
}

func (s *StorageService) ReadData(dataAccount libcore.Address, codeAccount libcore.Address) ([]byte, error) {
	root, err := s.storage.Group("/")
	if err != nil {
		return nil, err
	}
	group, err := getGroup(root, "data")
	if err != nil {
		return nil, err
	}
	account, err := getGroup(group, dataAccount.String())
	if err != nil {
		return nil, err
	}
	key, err := account.GetKey(codeAccount.String())
	if err != nil {
		return nil, err
	}
	data, err := s.storage.Get(key)
	if err != nil {
		return nil, err
	}
	reader := data.Open()
	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, reader); err != nil {
		return nil, err
	}
	reader.Close()
	data.Dispose()
	glog.Infoln("> read data", dataAccount.String(), codeAccount.String(), key.String(), data.Size())

	return buf.Bytes(), nil
}

func (s *StorageService) CreateUserData(account libcore.Address, destination libcore.Address, code libcore.Address, info *pb.UserInfo) (libcore.Hash, libcore.Address, error) {
	var newInfo *pb.UserInfo
	var dataAccount libcore.Address
	var contentData []byte
	if len(info.Data.Content) > 0 {
		contentData = info.Data.Content

		t := s.storage.Create(account.String())
		_, err := t.Write(contentData)
		if err != nil {
			return nil, nil, err
		}
		err = t.Close()
		if err != nil {
			return nil, nil, err
		}

		d, err := t.Data()
		if err != nil {
			return nil, nil, err
		}
		newInfo = &pb.UserInfo{
			Account: info.Account,
			Key:     info.Key,
			Nonce:   info.Nonce,
			Data: &pb.DataInfo{
				Hash: []byte(d.Key()),
			},
		}
		d.Dispose()

		dataAccount = account
	} else {
		k := core.Key(info.Data.Hash)
		if !s.storage.Exists(k) {
			return nil, nil, util.ErrorOfNotFound("user data", k.String())
		}
		newInfo = &pb.UserInfo{
			Account: info.Account,
			Key:     info.Key,
			Nonce:   info.Nonce,
			Data: &pb.DataInfo{
				Hash: []byte(k),
			},
		}

		dataAccount = destination
	}

	newData, err := core.Marshal(newInfo)
	if err != nil {
		return nil, nil, err
	}

	rootGroup, err := s.storage.Group("/")
	if err != nil {
		return nil, nil, err
	}
	dataGroup, err := getGroup(rootGroup, "user")
	if err != nil {
		return nil, nil, err
	}
	accountGroup, err := getGroup(dataGroup, dataAccount.String())
	if err != nil {
		return nil, nil, err
	}

	name := code.String()
	t := s.storage.Create(name)
	_, err = t.Write(newData)
	if err != nil {
		return nil, nil, err
	}
	err = t.Close()
	if err != nil {
		return nil, nil, err
	}

	d, err := t.Data()
	if err != nil {
		return nil, nil, err
	}
	_, err = accountGroup.AddData(name, d.Key())
	if err != nil {
		return nil, nil, err
	}

	err = accountGroup.Commit()
	if err != nil {
		return nil, nil, err
	}
	d.Dispose()
	t.Dispose()
	glog.Infoln("> write user", dataAccount.String(), name, d.Key().String(), len(newData), len(contentData))

	return libcore.Hash(d.Key()), dataAccount, nil
}

func (s *StorageService) GetUserData(userAccount libcore.Address, codeAccount libcore.Address) (*pb.UserInfo, error) {
	root, err := s.storage.Group("/")
	if err != nil {
		return nil, err
	}
	group, err := getGroup(root, "user")
	if err != nil {
		return nil, err
	}
	account, err := getGroup(group, userAccount.String())
	if err != nil {
		return nil, err
	}
	key, err := account.GetKey(codeAccount.String())
	if err != nil {
		return nil, err
	}
	data, err := s.storage.Get(key)
	if err != nil {
		return nil, err
	}
	reader := data.Open()
	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, reader); err != nil {
		return nil, err
	}
	reader.Close()
	data.Dispose()

	meta, msg, err := core.Unmarshal(buf.Bytes())
	if err != nil {
		return nil, err
	}
	if meta != core.CORE_USER_INFO {
		return nil, util.ErrorOfUnmatched("data", "user info", core.CORE_USER_INFO, meta)
	}
	info := msg.(*pb.UserInfo)
	contentKey := core.Key(info.Data.Hash)
	if !s.storage.Exists(key) {
		return nil, util.ErrorOfNotFound("content", contentKey.String())
	}
	contentSize, err := s.storage.Size(contentKey)
	if err != nil {
		return nil, err
	}

	glog.Infoln("> read user", userAccount.String(), codeAccount.String(), key.String(), data.Size(), contentKey.String(), contentSize)
	return info, nil
}

func (s *StorageService) CreatePage(name string, account libcore.Address, data []byte) (libcore.Hash, libcore.Hash, error) {
	rootGroup, err := s.storage.Group("/")
	if err != nil {
		return nil, nil, err
	}
	pageGroup, err := getGroup(rootGroup, "page")
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
	d, err := t.Data()
	if err != nil {
		return nil, nil, err
	}
	_, err = pageGroup.AddData(address, d.Key())
	if err != nil {
		return nil, nil, err
	}
	if len(name) > 0 {
		_, err = pageGroup.AddData(name, d.Key())
		if err != nil {
			return nil, nil, err
		}
	}
	err = pageGroup.Commit()
	if err != nil {
		return nil, nil, err
	}
	d.Dispose()
	t.Dispose()
	glog.Infoln("> create page", name, address, d.Key().String(), len(data))

	return libcore.Hash(s.storage.Root()), libcore.Hash(d.Key()), nil
}

func (s *StorageService) ReadPageByName(name string) ([]byte, error) {
	rootGroup, err := s.storage.Group("/")
	if err != nil {
		return nil, err
	}
	pageGroup, err := getGroup(rootGroup, "page")
	if err != nil {
		return nil, err
	}
	pageKey, err := pageGroup.GetKey(name)
	if err != nil {
		return nil, err
	}
	pageData, err := s.storage.Get(pageKey)
	if err != nil {
		return nil, err
	}
	pageReader := pageData.Open()
	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, pageReader); err != nil {
		return nil, err
	}
	pageReader.Close()
	pageData.Dispose()
	glog.Infoln("> read page", name, pageKey.String(), pageData.Size())

	return buf.Bytes(), nil
}

func (s *StorageService) ReadPageByAddress(account libcore.Address) ([]byte, error) {
	rootGroup, err := s.storage.Group("/")
	if err != nil {
		return nil, err
	}
	pageGroup, err := getGroup(rootGroup, "page")
	if err != nil {
		return nil, err
	}
	address := account.String()
	pageKey, err := pageGroup.GetKey(address)
	if err != nil {
		return nil, err
	}
	pageData, err := s.storage.Get(pageKey)
	if err != nil {
		return nil, err
	}
	pageReader := pageData.Open()
	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, pageReader); err != nil {
		return nil, err
	}
	pageReader.Close()
	pageData.Dispose()
	glog.Infoln("> read page", address, pageKey.String(), pageData.Size())

	return buf.Bytes(), nil
}

func (s *StorageService) CreateContract(account libcore.Address, wasmCode []byte, abiCode []byte) (libcore.Hash, libcore.Hash, error) {
	err := vm.VerifyWasm(wasmCode, abiCode)
	if err != nil {
		return nil, nil, err
	}
	k, err := s.WriteCode(account, wasmCode, abiCode)
	if err != nil {
		return nil, nil, err
	}
	return libcore.Hash(s.storage.Root()), libcore.Hash(k), nil
}

func (s *StorageService) RunContract(cs libcrypto.CryptoService, cost uint64, fromAccount libcore.Address, toAccount libcore.Address, method string, params [][]byte, inputs []libcore.Address, outputs []libcore.Address) (int64, libcore.Address, libcore.Hash, libcore.Hash, []byte, error) {
	wasmCode, abiCode, err := s.ReadCode(toAccount)
	if err != nil {
		return 0, nil, nil, nil, nil, err
	}
	var wasmData []byte
	if len(inputs) == 0 {
		wasmData, _ = s.ReadData(fromAccount, toAccount)
	} else {
		inputDatas := &pb.DataList{}
		for _, input := range inputs {
			wasmData, _ := s.ReadData(fromAccount, input)
			inputDatas.List = append(inputDatas.List, &pb.Data{
				Bytes: wasmData,
			})
		}
		wasmData, err = core.MarshalData(inputDatas)
		if err != nil {
			return 0, nil, nil, nil, nil, err
		}
	}
	usedCost, newWasmData, resultData, err := vm.RunWasm(int64(cost), wasmCode, abiCode, wasmData, method, params) // remainCost
	if err != nil {
		return 0, nil, nil, nil, nil, err
	}
	if newWasmData != nil {
		if len(outputs) == 0 {
			rootHash, dataHash, err := s.WriteData(fromAccount, toAccount, newWasmData)
			if err != nil {
				return 0, nil, nil, nil, nil, err
			}
			glog.Infoln("> run contract", usedCost, fromAccount.String(), toAccount.String(), method, len(inputs), len(outputs), len(wasmData), len(newWasmData), string(resultData))
			return usedCost, fromAccount, rootHash, dataHash, resultData, nil
		} else {
			t, msg, err := core.UnmarshalData(newWasmData)
			if err != nil {
				return 0, nil, nil, nil, nil, err
			}
			if t != core.CORE_DATA_LIST {
				return 0, nil, nil, nil, nil, util.ErrorOfInvalid("return data", "should be data list")
			}
			outputDatas := msg.(*pb.DataList)
			if len(outputDatas.List) != len(outputs) {
				return 0, nil, nil, nil, nil, util.ErrorOfInvalid("return data", fmt.Sprintf("%d != %d", len(outputDatas.List), len(outputs)))
			}
			resultInfos := &pb.DataList{}
			for index, outputData := range outputDatas.List {
				output := outputs[index]
				_, dataHash, err := s.WriteData(toAccount, output, outputData.Bytes)
				if err != nil {
					return 0, nil, nil, nil, nil, err
				}
				hashBytes, err := core.MarshalData([]byte(dataHash))
				if err != nil {
					return 0, nil, nil, nil, nil, err
				}
				resultInfos.List = append(resultInfos.List, &pb.Data{
					Bytes: hashBytes,
				})
			}
			resultDatas, err := core.MarshalData(resultInfos)
			if err != nil {
				return 0, nil, nil, nil, nil, err
			}
			dataHash, err := cs.Hash(resultDatas)
			if err != nil {
				return 0, nil, nil, nil, nil, err
			}
			glog.Infoln("> run contract", usedCost, fromAccount.String(), toAccount.String(), method, len(inputs), len(outputs), len(wasmData), len(newWasmData), len(resultDatas), string(resultData))
			return usedCost, toAccount, nil, dataHash, resultDatas, nil
		}
	} else {
		glog.Infoln("> run contract", usedCost, fromAccount.String(), toAccount.String(), method, resultData)
		return usedCost, fromAccount, nil, nil, resultData, nil
	}
}

func (s *StorageService) CallContract(dataAccount libcore.Address, codeAccount libcore.Address, method string, params [][]byte) (int64, interface{}, error) {
	wasmCode, abiCode, err := s.ReadCode(codeAccount)
	if err != nil {
		return 0, nil, err
	}
	wasmData, _ := s.ReadData(dataAccount, codeAccount)
	usedCost, _, resultData, err := vm.RunWasm(int64(1000000), wasmCode, abiCode, wasmData, method, params) // remainCost
	if err != nil {
		return 0, nil, err
	}
	r, err := core.AsData(resultData)
	if err != nil {
		return 0, nil, err
	}
	glog.Infoln("> call contract", usedCost, dataAccount.String(), codeAccount.String(), len(resultData), method, r)
	return usedCost, r, nil
}

func (s *StorageService) GetContractData(dataAccount libcore.Address, codeAccount libcore.Address, format string) (int64, interface{}, error) {
	wasmData, _ := s.ReadData(dataAccount, codeAccount)
	var r interface{}
	var e error
	switch format {
	case "data":
		r, e = core.AsData(wasmData)
		if e != nil {
			return 0, nil, e
		}
	case "string":
		r = string(wasmData)
	default:
		r = hex.EncodeToString(wasmData)
	}
	glog.Infoln("> get contract data", dataAccount.String(), codeAccount.String(), len(wasmData), r)
	return int64(len(wasmData)), r, nil
}

func (s *StorageService) GetData(hash libcore.Hash, format string) (int64, interface{}, error) {
	key := core.Key(hash)
	data, err := s.storage.Get(key)
	if err != nil {
		return 0, nil, err
	}
	reader := data.Open()
	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, reader); err != nil {
		return 0, nil, err
	}
	reader.Close()
	data.Dispose()

	contentData := buf.Bytes()
	var r interface{}
	var e error
	switch format {
	case "data":
		r, e = core.AsData(contentData)
		if e != nil {
			return 0, nil, e
		}
	case "string":
		r = string(contentData)
	default:
		r = hex.EncodeToString(contentData)
	}
	glog.Infoln("> get data", key.String(), len(contentData), r)

	return int64(len(contentData)), r, nil
}

func (s *StorageService) Init(c libcore.Config) error {
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

func (s *StorageService) Start() error {
	return nil
}

func (s *StorageService) Close() error {
	err := s.storage.Close()
	if err != nil {
		return err
	}
	return nil
}

// test
func (s *StorageService) dump(name string) {
	name = fmt.Sprintf("%s(%d)", name, s.stackdb.Len())

	dumpStorage(s.storage, name+": current storage")
	reloadStorage, err := chunk.LoadStorageWith(s.stackdb, 2048*2048*8)
	if err != nil {
		glog.Error(err)
	} else {
		dumpStorage(reloadStorage, name+": reload storage")
	}
	for i := (s.stackdb.Len() - 1); i >= 0; i-- {
		db := s.stackdb.Get(i)
		s, err := chunk.LoadStorageWith(db, 2048*2048*8)
		if err != nil {
			glog.Error(err)
		} else {
			dumpStorage(s, fmt.Sprintf(name+": s%d", i))
		}
	}
	dumpKv(s.stackdb, name+":stackdb")
	for i := (s.stackdb.Len() - 1); i >= 0; i-- {
		db := s.stackdb.Get(i)
		dumpKv(db, fmt.Sprintf(name+":db%d", i))
	}
}

func (s *StorageService) dumpWith(name string, popKv libstore.KvService) {
	s.dump(name)

	if popKv != nil {
		name = fmt.Sprintf("%s(%d)", name, s.stackdb.Len())
		s, err := chunk.LoadStorageWith(popKv, 2048*2048*8)
		if err != nil {
			glog.Error(err)
		} else {
			dumpStorage(s, name+":pop storage")
		}
		dumpKv(popKv, name+":pop db")
	}
}

func dumpKv(s libstore.KvService, name string) {
	fmt.Println("~~~~~~~~~~~~~~~~~~~" + name + "~~~~~~~~~~~~~~~~~~~")
	s.ListData(func(key []byte, value []byte) error {
		fmt.Println(string(key), key, value)
		return nil
	})
}
func dumpStorage(s core.Storage, name string) {
	fmt.Println("=================" + name + "=================")
	s.DumpLog(chunk.LogPrinter{})
}

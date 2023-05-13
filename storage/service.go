package storage

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"path"
	"sync"

	"github.com/caivega/glog"
	"github.com/tokentransfer/node/account"
	"github.com/tokentransfer/node/block"
	"github.com/tokentransfer/node/core"
	"github.com/tokentransfer/node/core/pb"
	"github.com/tokentransfer/node/crypto"
	"github.com/tokentransfer/node/db"
	"github.com/tokentransfer/node/util"
	"github.com/tokentransfer/node/vm"

	libcore "github.com/tokentransfer/interfaces/core"
	libstore "github.com/tokentransfer/interfaces/store"
)

var as = account.NewAccountService()
var cs = &crypto.CryptoService{}

type StorageService struct {
	storage core.Storage
	stackdb *db.StackService

	categories []string
	locker     sync.Mutex
}

func NewStorageService(c libcore.Config, account libcore.Address) (*StorageService, error) {
	service := &StorageService{
		categories: func(a libcore.Address) []string {
			if a != nil {
				return []string{
					"account",
					a.String(),
				}
			}
			return []string{}
		}(account),
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

func (s *StorageService) RootHash() libcore.Hash {
	return libcore.Hash(s.storage.Root())
}

func (s *StorageService) GetStorage() core.Storage {
	return s.storage
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
	memdb := &db.MemoryService{
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

func (s *StorageService) getRoot(rootAccount libcore.Address) (core.Group, error) {
	root, err := s.storage.Group("/")
	if err != nil {
		return nil, err
	}
	if rootAccount == nil {
		return root, nil
	}
	g, err := getGroup(root, rootAccount.String())
	if err != nil {
		return nil, err
	}
	return g, nil
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

func (s *StorageService) read(key core.Key) ([]byte, error) {
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

	return buf.Bytes(), nil
}

func (s *StorageService) write(name string, data []byte) (core.Data, error) {
	t := s.storage.Create(name)
	_, err := t.Write(data)
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
	t.Dispose()
	return d, nil
}

func (s *StorageService) writeData(rootGroup core.Group, category string, dir string, name string, data []byte) (libcore.Hash, libcore.Hash, libcore.Hash, libcore.Hash, error) {
	var categoryGroup core.Group
	if len(category) > 0 {
		group, err := getGroup(rootGroup, category)
		if err != nil {
			return nil, nil, nil, nil, err
		}
		categoryGroup = group
	} else {
		categoryGroup = rootGroup
	}

	var dirGroup core.Group
	if len(dir) > 0 {
		group, err := getGroup(categoryGroup, dir)
		if err != nil {
			return nil, nil, nil, nil, err
		}
		dirGroup = group
	} else {
		dirGroup = categoryGroup
	}

	t := s.storage.Create(name)
	_, err := t.Write(data)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	err = t.Close()
	if err != nil {
		return nil, nil, nil, nil, err
	}

	d, err := t.Data()
	if err != nil {
		return nil, nil, nil, nil, err
	}
	_, err = dirGroup.AddData(name, d.Key())
	if err != nil {
		return nil, nil, nil, nil, err
	}

	err = dirGroup.Commit()
	if err != nil {
		return nil, nil, nil, nil, err
	}
	d.Dispose()
	t.Dispose()
	glog.Infoln("> write data", category, dir, name, len(data), categoryGroup.Key().Short(), dirGroup.Key().Short(), d.Key().String())

	return libcore.Hash(s.storage.Root()), libcore.Hash(categoryGroup.Key()), libcore.Hash(dirGroup.Key()), libcore.Hash(d.Key()), nil
}

func (s *StorageService) readData(rootGroup core.Group, category string, dir string, name string) (core.Key, core.Key, core.Key, []byte, error) {
	var categoryGroup core.Group
	if len(category) > 0 {
		group, err := rootGroup.Group(category)
		if err != nil {
			return nil, nil, nil, nil, err
		}
		categoryGroup = group
	} else {
		categoryGroup = rootGroup
	}

	var dirGroup core.Group
	if len(dir) > 0 {
		group, err := categoryGroup.Group(dir)
		if err != nil {
			return nil, nil, nil, nil, err
		}
		dirGroup = group
	} else {
		dirGroup = categoryGroup
	}

	key, err := dirGroup.GetKey(name)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	data, err := s.storage.Get(key)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	reader := data.Open()
	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, reader); err != nil {
		return nil, nil, nil, nil, err
	}
	reader.Close()
	data.Dispose()

	content := buf.Bytes()
	glog.Infoln("> read data", category, dir, name, len(content), categoryGroup.Key().Short(), dirGroup.Key().Short(), key.String())
	return categoryGroup.Key(), dirGroup.Key(), key, content, nil
}

func (s *StorageService) writeMap(rootGroup core.Group, category string, dir string, name string, m map[string]interface{}) (libcore.Hash, libcore.Hash, libcore.Hash, libcore.Hash, error) {
	data, err := core.MarshalData(m)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	return s.writeData(rootGroup, category, dir, name, data)
}

func (s *StorageService) readMap(rootGroup core.Group, category string, dir string, name string) (core.Key, core.Key, core.Key, map[string]interface{}, error) {
	categoryKey, dirKey, key, metaData, err := s.readData(rootGroup, category, dir, name)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	t, msg, err := core.UnmarshalData(metaData)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	if t != core.CORE_DATA_MAP {
		return nil, nil, nil, nil, util.ErrorOfInvalid("data map", "data")
	}
	m := msg.(map[string]interface{})
	return categoryKey, dirKey, key, m, nil
}

func (s *StorageService) ReadCode(codeAccount libcore.Address) ([]byte, []byte, error) {
	rootGroup, err := s.getRoot(nil)
	if err != nil {
		return nil, nil, err
	}
	_, _, _, wasmCode, err := s.readData(rootGroup, "code", codeAccount.String(), "code")
	if err != nil {
		return nil, nil, err
	}
	_, _, _, abiCode, _ := s.readData(rootGroup, "code", codeAccount.String(), "abi") // optional
	return wasmCode, abiCode, nil
}

func (s *StorageService) WriteCode(account libcore.Address, wasmCode []byte, abiCode []byte) (libcore.Hash, error) {
	rootGroup, err := s.getRoot(nil)
	if err != nil {
		return nil, err
	}
	_, _, _, codeHash, err := s.writeData(rootGroup, "code", account.String(), "code", wasmCode)
	if err != nil {
		return nil, err
	}
	if len(abiCode) > 0 {
		_, _, dirHash, _, err := s.writeData(rootGroup, "code", account.String(), "abi", abiCode)
		if err != nil {
			return nil, err
		}
		return dirHash, nil
	}
	return codeHash, nil
}

func (s *StorageService) ReadMeta(symbol libcore.Symbol) (uint64, libcore.Address, *pb.MetaInfo, error) {
	rootGroup, err := s.getRoot(nil)
	if err != nil {
		return 0, nil, nil, err
	}
	_, _, _, metaMap, err := s.readMap(rootGroup, "token", "meta", symbol.String())
	if err != nil {
		return 0, nil, nil, err
	}
	_, _, _, indexMap, err := s.readMap(rootGroup, "token", "index", symbol.String())
	if err != nil {
		return 0, nil, nil, err
	}
	account := util.ToString(&metaMap, "account")
	_, a, err := as.NewAccountFromAddress(account)
	if err != nil {
		return 0, nil, nil, err
	}
	info := toMetaInfo(&metaMap, "info")
	index := util.ToUint64(&indexMap, "index")
	return index, a, info, nil
}

func toMetaInfo(m *map[string]interface{}, key string) *pb.MetaInfo {
	if m != nil {
		data := (*m)[key]
		if data != nil {
			info, ok := data.(*pb.MetaInfo)
			if ok {
				return info
			}
		}
	}
	return nil
}

func (s *StorageService) WriteMeta(account libcore.Address, info *pb.MetaInfo) (*block.DataInfo, error) {
	rootGroup, err := s.getRoot(nil)
	if err != nil {
		return nil, err
	}
	_, _, _, metaHash, err := s.writeMap(rootGroup, "token", "meta", info.Symbol, map[string]interface{}{
		"account": account.String(),
		"info":    info,
	})
	if err != nil {
		return nil, err
	}
	_, _, _, indexHash, err := s.writeMap(rootGroup, "token", "index", info.Symbol, map[string]interface{}{
		"index": uint64(0),
	})
	if err != nil {
		return nil, err
	}

	group := &block.GroupInfo{}
	group.PutData("meta", &block.DataInfo{
		Hash: metaHash,
	})
	group.PutData("index", &block.DataInfo{
		Hash: indexHash,
	})
	dataInfo, err := block.GetDataInfo(group)
	if err != nil {
		return nil, err
	}
	return dataInfo, nil
}

func (s *StorageService) ReadToken(symbol libcore.Symbol, index uint64) (libcore.Address, *pb.TokenInfo, error) {
	rootGroup, err := s.getRoot(nil)
	if err != nil {
		return nil, nil, err
	}
	_, _, _, tokenMap, err := s.readMap(rootGroup, "token", "token", util.GetTokenKey(symbol.String(), index))
	if err != nil {
		return nil, nil, err
	}
	account := util.ToString(&tokenMap, "account")
	_, a, err := as.NewAccountFromAddress(account)
	if err != nil {
		return nil, nil, err
	}
	info := toTokenInfo(&tokenMap, "info")
	return a, info, nil
}

func toTokenInfo(m *map[string]interface{}, key string) *pb.TokenInfo {
	if m != nil {
		data := (*m)[key]
		if data != nil {
			info, ok := data.(*pb.TokenInfo)
			if ok {
				return info
			}
		}
	}
	return nil
}

func (s *StorageService) WriteToken(account libcore.Address, info *pb.TokenInfo) (*block.DataInfo, error) {
	rootGroup, err := s.getRoot(nil)
	if err != nil {
		return nil, err
	}
	_, _, _, tokenHash, err := s.writeMap(rootGroup, "token", "token", util.GetTokenKey(info.Symbol, info.Index), map[string]interface{}{
		"account": account.String(),
		"info":    info,
	})
	if err != nil {
		return nil, err
	}
	_, _, _, indexHash, err := s.writeMap(rootGroup, "token", "index", info.Symbol, map[string]interface{}{
		"index": info.Index,
	})
	if err != nil {
		return nil, err
	}
	group := &block.GroupInfo{}
	group.PutData("token", &block.DataInfo{
		Hash: tokenHash,
	})
	group.PutData("index", &block.DataInfo{
		Hash: indexHash,
	})
	dataInfo, err := block.GetDataInfo(group)
	if err != nil {
		return nil, err
	}
	return dataInfo, nil
}

func (s *StorageService) ReadMemory(rootAccount libcore.Address, dataAccount libcore.Address, codeAccount libcore.Address) ([]byte, error) {
	rootGroup, err := s.getRoot(rootAccount)
	if err != nil {
		return nil, err
	}
	_, _, _, data, err := s.readData(rootGroup, "memory", dataAccount.String(), codeAccount.String())
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *StorageService) WriteMemory(rootAccount libcore.Address, dataAccount libcore.Address, codeAccount libcore.Address, data []byte) (libcore.Hash, libcore.Hash, error) {
	rootGroup, err := s.getRoot(rootAccount)
	if err != nil {
		return nil, nil, err
	}
	rootHash, _, _, dataHash, err := s.writeData(rootGroup, "memory", dataAccount.String(), codeAccount.String(), data)
	if err != nil {
		return nil, nil, err
	}
	return rootHash, dataHash, nil
}

func (s *StorageService) ReadUser(rootAccount libcore.Address, userAccount libcore.Address, codeAccount libcore.Address) (*pb.UserInfo, error) {
	rootGroup, err := s.getRoot(rootAccount)
	if err != nil {
		return nil, err
	}
	_, _, _, data, err := s.readData(rootGroup, "user", userAccount.String(), codeAccount.String())
	if err != nil {
		return nil, err
	}
	meta, msg, err := core.Unmarshal(data)
	if err != nil {
		return nil, err
	}
	if meta != core.CORE_USER_INFO {
		return nil, util.ErrorOfUnmatched("data", "user info", core.CORE_USER_INFO, meta)
	}
	info := msg.(*pb.UserInfo)
	contentKey := core.Key(info.Data.Hash)
	if !s.storage.Exists(contentKey) {
		return nil, util.ErrorOfNotFound("content", contentKey.String())
	}
	contentSize, err := s.storage.Size(contentKey)
	if err != nil {
		return nil, err
	}
	glog.Infoln("> read user", contentKey.String(), contentSize)
	return info, nil
}

func (s *StorageService) WriteUser(account libcore.Address, destination libcore.Address, code libcore.Address, info *pb.UserInfo) (libcore.Hash, libcore.Address, error) {
	var newInfo *pb.UserInfo
	var dataAccount libcore.Address
	var contentKey core.Key
	var contentSize int64
	if len(info.Data.Content) > 0 {
		contentSize = int64(len(info.Data.Content))

		t := s.storage.Create(account.String())
		_, err := t.Write(info.Data.Content)
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
		contentKey = d.Key()
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
		contentKey = k
		size, err := s.storage.Size(k)
		if err != nil {
			return nil, nil, err
		}
		contentSize = size
	}
	newData, err := core.Marshal(newInfo)
	if err != nil {
		return nil, nil, err
	}

	rootGroup, err := s.getRoot(dataAccount)
	if err != nil {
		return nil, nil, err
	}
	_, _, _, dataHash, err := s.writeData(rootGroup, "user", dataAccount.String(), code.String(), newData)
	if err != nil {
		return nil, nil, err
	}
	glog.Infoln("> write user", contentKey.String(), contentSize)
	return dataHash, dataAccount, nil
}

func (s *StorageService) WritePage(name string, account libcore.Address, data []byte) (libcore.Hash, libcore.Hash, error) {
	rootGroup, err := s.getRoot(nil)
	if err != nil {
		return nil, nil, err
	}
	_, _, _, dataHash, err := s.writeData(rootGroup, "code", account.String(), "page", data)
	if err != nil {
		return nil, nil, err
	}
	key := core.Key(dataHash)
	if len(name) > 0 {
		pageGroup, err := getGroup(rootGroup, "page")
		if err != nil {
			return nil, nil, err
		}
		_, err = pageGroup.AddData(name, key)
		if err != nil {
			return nil, nil, err
		}
		err = pageGroup.Commit()
		if err != nil {
			return nil, nil, err
		}
	}
	glog.Infoln("> create page", name, account.String(), key.String(), len(data))
	return libcore.Hash(s.storage.Root()), dataHash, nil
}

func (s *StorageService) ReadPageByName(name string) ([]byte, error) {
	rootGroup, err := s.getRoot(nil)
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
	rootGroup, err := s.getRoot(nil)
	if err != nil {
		return nil, err
	}
	_, _, _, data, err := s.readData(rootGroup, "code", account.String(), "page")
	if err != nil {
		return nil, err
	}
	return data, nil
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

func (s *StorageService) RunContract(cost uint64, signAccount libcore.Address, fromAccount libcore.Address, toAccount libcore.Address, method string, params [][]byte, inputs []libcore.Address, outputs []libcore.Address) (int64, libcore.Address, libcore.Hash, libcore.Hash, []byte, error) {
	wasmCode, abiCode, err := s.ReadCode(toAccount)
	if err != nil {
		return 0, nil, nil, nil, nil, err
	}
	var wasmData []byte
	if len(inputs) == 0 {
		wasmData, _ = s.ReadMemory(fromAccount, fromAccount, toAccount)
	} else {
		inputDatas := &pb.DataList{}
		for _, input := range inputs {
			wasmData, _ := s.ReadMemory(fromAccount, fromAccount, input)
			inputDatas.List = append(inputDatas.List, &pb.Data{
				Bytes: wasmData,
			})
		}
		wasmData, err = core.MarshalData(inputDatas)
		if err != nil {
			return 0, nil, nil, nil, nil, err
		}
	}
	usedCost, newWasmData, resultData, err := vm.RunWasm(int64(cost), wasmCode, abiCode, wasmData, method, params, (signAccount.String() == fromAccount.String())) // remainCost
	if err != nil {
		return 0, nil, nil, nil, nil, err
	}
	if newWasmData != nil {
		if len(outputs) == 0 {
			rootHash, dataHash, err := s.WriteMemory(fromAccount, fromAccount, toAccount, newWasmData)
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
				_, dataHash, err := s.WriteMemory(toAccount, toAccount, output, outputData.Bytes)
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

func (s *StorageService) CallContract(rootAccount libcore.Address, dataAccount libcore.Address, codeAccount libcore.Address, method string, params [][]byte) (int64, interface{}, error) {
	wasmCode, abiCode, err := s.ReadCode(codeAccount)
	if err != nil {
		return 0, nil, err
	}
	wasmData, _ := s.ReadMemory(rootAccount, dataAccount, codeAccount)
	usedCost, _, resultData, err := vm.RunWasm(int64(1000000), wasmCode, abiCode, wasmData, method, params, true) // remainCost
	if err != nil {
		return 0, nil, err
	}
	t, r, err := core.UnmarshalData(resultData)
	if err != nil {
		return 0, nil, err
	}
	glog.Infoln("> call contract", usedCost, dataAccount.String(), codeAccount.String(), len(resultData), method, t, r)
	return usedCost, r, nil
}

func (s *StorageService) GetContractData(rootAccount libcore.Address, dataAccount libcore.Address, codeAccount libcore.Address, format string) (int64, interface{}, error) {
	wasmData, _ := s.ReadMemory(rootAccount, dataAccount, codeAccount)
	var r interface{}
	var e error
	switch format {
	case "data":
		_, r, e = core.UnmarshalData(wasmData)
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
		_, r, e = core.UnmarshalData(contentData)
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

func (s *StorageService) ReadState(rootAccount libcore.Address, theAccount libcore.Address, name string) ([]byte, error) {
	rootGroup, err := s.getRoot(rootAccount)
	if err != nil {
		return nil, err
	}
	_, _, _, data, err := s.readData(rootGroup, "state", theAccount.String(), name)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *StorageService) WriteState(rootAccount libcore.Address, theAccount libcore.Address, name string, data []byte) (libcore.Hash, libcore.Hash, error) {
	rootGroup, err := s.getRoot(rootAccount)
	if err != nil {
		return nil, nil, err
	}
	rootHash, _, _, dataHash, err := s.writeData(rootGroup, "state", theAccount.String(), name, data)
	if err != nil {
		return nil, nil, err
	}
	return rootHash, dataHash, nil
}

func (s *StorageService) GetGas(rootAccount libcore.Address, theAccount libcore.Address) (*util.Value, error) {
	data, err := s.ReadState(rootAccount, theAccount, "gas")
	if err != nil {
		return nil, err
	}
	value, err := util.NewValue(string(data))
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (s *StorageService) UpdateGas(rootAccount libcore.Address, theAccount libcore.Address, value util.Value) error {
	_, _, err := s.WriteState(rootAccount, theAccount, "gas", []byte(value.String()))
	if err != nil {
		return err
	}
	return nil
}

func (s *StorageService) GetAccountGas(rootAccount libcore.Address, account libcore.Address) (*util.Value, *util.Value, error) {
	value, err := s.GetGas(rootAccount, account)
	if err != nil {
		return nil, nil, err
	}
	localValue, err := s.GetGas(rootAccount, account)
	if err != nil {
		return nil, nil, err
	}
	return value, localValue, nil
}

func (s *StorageService) Init(c libcore.Config) error {
	dataDir := c.GetDataDir()
	if len(s.categories) > 0 {
		list := append([]string{dataDir}, s.categories...)
		dataDir = path.Join(list...)
	}
	dbPath := path.Join(dataDir, "data")

	datadb := &db.LevelService{Path: dbPath}
	err := datadb.Init(c)
	if err != nil {
		return err
	}
	err = datadb.Start()
	if err != nil {
		return err
	}
	stackdb, err := db.CreateStackService(datadb)
	if err != nil {
		return err
	}
	storage, err := LoadStorageWith(stackdb, 256*1024*1024)
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
	reloadStorage, err := LoadStorageWith(s.stackdb, 2048*2048*8)
	if err != nil {
		glog.Error(err)
	} else {
		dumpStorage(reloadStorage, name+": reload storage")
	}
	for i := (s.stackdb.Len() - 1); i >= 0; i-- {
		db := s.stackdb.Get(i)
		s, err := LoadStorageWith(db, 2048*2048*8)
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
		s, err := LoadStorageWith(popKv, 2048*2048*8)
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
	s.DumpLog(LogPrinter{})
}

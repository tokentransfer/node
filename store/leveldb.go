package store

import (
	"path"

	"github.com/syndtr/goleveldb/leveldb"

	libcore "github.com/tokentransfer/interfaces/core"
	"github.com/tokentransfer/node/util"
)

type LevelService struct {
	Path string
	Name string

	config libcore.Config
	db     *leveldb.DB
}

func (service *LevelService) Close() error {
	if service.db != nil {
		err := service.db.Close()
		if err != nil {
			if err == leveldb.ErrClosed {
				return nil
			}
			return err
		}
	}

	return nil
}

func (service *LevelService) open() error {
	if len(service.Path) > 0 {
		db, err := serviceForLevelDB(service.Path)
		if err != nil {
			return err
		}
		service.db = db
	} else {
		if service.config != nil {
			dataDir := service.config.GetDataDir()
			dbPath := path.Join(dataDir, service.Name)
			db, err := serviceForLevelDB(dbPath)
			if err != nil {
				return err
			}
			service.db = db
		} else {
			return util.ErrorOfEmpty("config or path", "for leveldb")
		}
	}
	return nil
}

func (service *LevelService) Init(c libcore.Config) error {
	service.config = c
	return service.open()
}

func (service *LevelService) Start() error {
	return nil
}

func (service *LevelService) PutData(key []byte, value []byte) error {
	db := service.db

	err := db.Put(key, value, nil)
	if err != nil {
		return err
	}
	return nil
}

func (service *LevelService) PutDatas(keys [][]byte, values [][]byte) error {
	db := service.db

	lk := len(keys)
	lv := len(values)
	if lk != lv {
		return util.ErrorOfUnmatched("the length", "datas", lk, lv)
	}
	bs := leveldb.MakeBatch(lk)
	for i := 0; i < lk; i++ {
		bs.Put(keys[i], values[i])
	}
	err := db.Write(bs, nil)
	if err != nil {
		return err
	}
	return nil
}

func (service *LevelService) Flush() error {
	return nil
}

func (service *LevelService) GetData(key []byte) ([]byte, error) {
	db := service.db
	bytes, err := db.Get(key, nil)
	if err != nil {
		if err == leveldb.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return bytes, nil
}

func (service *LevelService) GetDatas(keys [][]byte) ([][]byte, error) {
	db := service.db

	l := len(keys)
	bytes := make([][]byte, l)
	for i := 0; i < l; i++ {
		value, err := db.Get(keys[i], nil)
		if err != nil {
			if err == leveldb.ErrNotFound {
				return nil, nil
			}
			return nil, err
		}
		bytes[i] = value
	}
	return bytes, nil
}

func (service *LevelService) HasData(key []byte) bool {
	db := service.db

	ok, err := db.Has(key, nil)
	if err != nil {
		return false
	}
	if !ok {
		return false
	}

	return true
}

func (service *LevelService) RemoveData(key []byte) error {
	db := service.db

	err := db.Delete(key, nil)
	if err != nil {
		return err
	}

	return nil
}

func (service *LevelService) ListData(each func(key []byte, value []byte) error) error {
	db := service.db

	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()
		err := each(key, value)
		if err != nil {
			return err
		}
	}
	iter.Release()
	return iter.Error()
}

func serviceForLevelDB(dbPath string) (*leveldb.DB, error) {
	db, err := leveldb.OpenFile(dbPath, nil)
	if err != nil {
		return nil, err
	}
	return db, nil
}

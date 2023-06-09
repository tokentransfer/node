package db

import (
	"encoding/hex"
	"sync"

	"github.com/caivega/glog"
	libcore "github.com/tokentransfer/interfaces/core"
	libstore "github.com/tokentransfer/interfaces/store"
	"github.com/tokentransfer/node/util"
)

type StackService struct {
	list []libstore.KvService

	m sync.Map
}

// KV -> MEMORY -> MEMORY
func CreateStackService(validatedService libstore.KvService) (*StackService, error) {
	if validatedService == nil {
		return nil, util.ErrorOfInvalid("service", "validated")
	}
	return &StackService{
		list: []libstore.KvService{validatedService},
		m:    sync.Map{},
	}, nil
}

func (stack *StackService) Init(c libcore.Config) error {
	return nil
}

func (stack *StackService) Start() error {
	return nil
}

func (stack *StackService) Close() error {
	for i := len(stack.list) - 1; i >= 0; i-- {
		service := stack.list[i]
		if err := service.Close(); err != nil {
			return err
		}
	}
	return nil
}

func (stack *StackService) Push(s libstore.KvService) error {
	stack.list = append(stack.list, s)
	return nil
}

func (stack *StackService) Pop() libstore.KvService {
	l := len(stack.list)
	if l > 0 {
		s := stack.list[l-1]
		stack.list = stack.list[:l-1]
		return s
	}
	return nil
}

func (stack *StackService) Top() libstore.KvService {
	l := len(stack.list)
	if l > 0 {
		s := stack.list[l-1]
		return s
	}
	return nil
}

func (stack *StackService) Get(index int) libstore.KvService {
	return stack.list[index]
}

func (stack *StackService) Len() int {
	return len(stack.list)
}

func (stack *StackService) Empty() bool {
	return len(stack.list) == 0
}

func (stack *StackService) GetData(key []byte) ([]byte, error) {
	_, ok := stack.m.Load(hex.EncodeToString(key))
	if ok {
		return nil, nil
	}

	for i := len(stack.list) - 1; i >= 0; i-- {
		service := stack.list[i]
		if service.HasData(key) {
			if data, err := service.GetData(key); err != nil {
				return nil, err
			} else {
				return data, nil
			}
		}
	}
	return nil, nil
}

func (stack *StackService) GetDatas(keys [][]byte) ([][]byte, error) {
	l := len(keys)
	bytes := make([][]byte, l)
	for i := 0; i < l; i++ {
		_, ok := stack.m.Load(hex.EncodeToString(keys[i]))
		if ok {
			bytes[i] = nil
		} else {
			value, err := stack.GetData(keys[i])
			if err != nil {
				bytes[i] = nil
			} else {
				bytes[i] = value
			}
		}
	}
	return bytes, nil
}

func (stack *StackService) PutData(key []byte, value []byte) error {
	service := stack.Top()
	if service != nil {
		err := service.PutData(key, value)
		if err != nil {
			return err
		}
		stack.m.Delete(hex.EncodeToString(key))
	}
	return nil
}

func (stack *StackService) PutDatas(keys [][]byte, values [][]byte) error {
	service := stack.Top()
	if service != nil {
		err := service.PutDatas(keys, values)
		if err != nil {
			return err
		}
		l := len(keys)
		for i := 0; i < l; i++ {
			stack.m.Delete(hex.EncodeToString(keys[i]))
		}
	}
	return nil
}

func (stack *StackService) HasData(key []byte) bool {
	_, ok := stack.m.Load(hex.EncodeToString(key))
	if ok {
		return false
	}

	for i := len(stack.list) - 1; i >= 0; i-- {
		service := stack.list[i]
		if has := service.HasData(key); has {
			return true
		}
	}
	return false
}

func (stack *StackService) RemoveData(key []byte) error {
	service := stack.Top()
	if service != nil {
		_ = service.RemoveData(key)
	}
	stack.m.Store(hex.EncodeToString(key), struct{}{})
	return nil
}

func (stack *StackService) Flush() error {
	for i := len(stack.list) - 1; i >= 0; i-- {
		service := stack.list[i]
		if err := service.Flush(); err != nil {
			return err
		}
	}
	return nil
}

func (stack *StackService) ListData(each func(key []byte, value []byte) error) error {
	m := sync.Map{}
	for i := 0; i < len(stack.list); i++ {
		s := stack.list[i]
		s.ListData(func(key []byte, value []byte) error {
			_, ok := stack.m.Load(hex.EncodeToString(key))
			if !ok {
				m.Store(hex.EncodeToString(key), hex.EncodeToString(value))
			}
			return nil
		})
	}

	m.Range(func(k, v interface{}) bool {
		var key, value []byte

		if k != nil {
			data, err := hex.DecodeString(k.(string))
			if err != nil {
				glog.Error(err)
				key = nil
			} else {
				key = data
			}
		} else {
			key = nil
		}

		if v != nil {
			data, err := hex.DecodeString(v.(string))
			if err != nil {
				glog.Error(err)
				value = nil
			} else {
				value = data
			}
		} else {
			value = nil
		}

		err := each(key, value)
		if err != nil {
			glog.Error(err)
		}
		return err == nil
	})
	return nil
}

func (stack *StackService) Commit() (libstore.KvService, error) {
	l := len(stack.list)
	if l > 1 {
		top := stack.list[l-1]
		bottom := stack.list[l-2]
		err := top.ListData(func(key []byte, value []byte) error {
			err := bottom.PutData(key, value)
			if err != nil {
				glog.Error(err)
				return err
			}
			return nil
		})
		if err != nil {
			glog.Error(err)
			return nil, err
		}

		stack.m.Range(func(k, v interface{}) bool {
			var key []byte

			if k != nil {
				data, err := hex.DecodeString(k.(string))
				if err != nil {
					glog.Error(err)
					key = nil
				} else {
					key = data
				}
			} else {
				key = nil
			}

			if key != nil {
				err := bottom.RemoveData(key)
				if err != nil {
					glog.Error(err)
				}
				return err == nil
			}
			return true
		})

		popKv := stack.Pop()
		stack.m = sync.Map{}
		return popKv, nil
	}
	return nil, nil
}

func (stack *StackService) Cancel() (libstore.KvService, error) {
	l := len(stack.list)
	if l > 1 {
		popKv := stack.Pop()
		stack.m = sync.Map{}
		return popKv, nil
	}
	return nil, nil
}

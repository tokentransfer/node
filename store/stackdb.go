package store

import (
	"encoding/hex"
	"sync"

	libcore "github.com/tokentransfer/interfaces/core"
	libstore "github.com/tokentransfer/interfaces/store"
	"github.com/tokentransfer/node/core"
)

type StackService struct {
	list []libstore.KvService

	m sync.Map
}

// VALIDATED -> CURRENT -> CURRENT
func CreateStackService(validatedService libstore.KvService) (*StackService, error) {
	if validatedService == nil {
		return nil, core.ErrorOfInvalid("service", "validated")
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
		if data, err := service.GetData(key); err != nil {
			return nil, err
		} else {
			if len(data) > 0 {
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
		err := service.RemoveData(key)
		if err != nil {
			return err
		}
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
				m.Store(hex.EncodeToString(key), value)
			}
			return nil
		})
	}

	m.Range(func(k, v interface{}) bool {
		var key, value []byte

		if k != nil {
			data, err := hex.DecodeString(k.(string))
			if err != nil {
				key = nil
			} else {
				key = data
			}
		} else {
			key = nil
		}

		if v != nil {
			value = v.([]byte)
		} else {
			value = nil
		}

		err := each(key, value)
		return err == nil
	})
	return nil
}

func (stack *StackService) Commit() error {
	l := len(stack.list)
	if l > 1 {
		top := stack.list[l-1]
		bottom := stack.list[l-2]
		err := top.ListData(func(key []byte, value []byte) error {
			err := bottom.PutData(key, value)
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}

		stack.m.Range(func(k, v interface{}) bool {
			var key []byte

			if k != nil {
				data, err := hex.DecodeString(k.(string))
				if err != nil {
					key = nil
				} else {
					key = data
				}
			} else {
				key = nil
			}

			if key != nil {
				err := bottom.RemoveData(key)
				return err == nil
			}
			return true
		})

		stack.Pop()
		stack.m = sync.Map{}
	}
	return nil
}

func (stack *StackService) Cancel() error {
	l := len(stack.list)
	if l > 1 {
		stack.Pop()
		stack.m = sync.Map{}
	}
	return nil
}

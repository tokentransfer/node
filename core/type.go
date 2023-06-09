package core

import (
	"encoding/hex"
	"strconv"
	"strings"

	"github.com/tetratelabs/wazero/api"
	"github.com/tokentransfer/node/util"
)

type DataType byte

var nameMap = map[string]string{}
var dataTypes = map[string]DataType{}
var dataNames = map[DataType]string{}
var dataPrimaries = map[DataType]bool{}

func (t DataType) Create(name string, isPrimary bool) {
	n := strings.ToLower(name)
	nameMap[n] = name

	dataTypes[n] = t
	dataNames[t] = n
	dataPrimaries[t] = isPrimary
}

func (t DataType) Primary() bool {
	b := dataPrimaries[t]
	return b
}

func (t DataType) String() string {
	n := dataNames[t]
	return nameMap[n]
}

func AsBytes(data []byte) ([]byte, error) {
	meta, value, err := UnmarshalData(data)
	if err != nil {
		return nil, err
	}
	var v []byte
	var e error
	switch meta {
	case CORE_DATA_STRING:
		if value != nil {
			v = []byte(value.(string))
		} else {
			v = nil
		}
		e = nil
	case CORE_DATA_BYTES:
		if value != nil {
			v = value.([]byte)
		} else {
			v = nil
		}
		e = nil
	default:
		v = nil
		e = util.ErrorOfInvalid("type", "data")
	}
	if e != nil {
		return nil, e
	}
	return v, nil
}

func AsUint64(data []byte) (uint64, error) {
	meta, value, err := UnmarshalData(data)
	if err != nil {
		return 0, err
	}
	var v uint64
	var e error
	switch meta {
	case CORE_DATA_INT8:
		v = uint64(value.(int8))
		e = nil
	case CORE_DATA_UINT8:
		v = uint64(value.(uint8))
		e = nil
	case CORE_DATA_INT16:
		v = uint64(value.(int16))
		e = nil
	case CORE_DATA_UINT16:
		v = uint64(value.(uint16))
		e = nil
	case CORE_DATA_INT32:
		v = uint64(value.(int32))
		e = nil
	case CORE_DATA_UINT32:
		v = uint64(value.(uint32))
		e = nil
	case CORE_DATA_INT64:
		v = uint64(value.(int64))
		e = nil
	case CORE_DATA_UINT64:
		v = value.(uint64)
		e = nil
	case CORE_DATA_FLOAT32:
		v = uint64(api.EncodeF32(value.(float32)))
		e = nil
	case CORE_DATA_FLOAT64:
		v = uint64(api.EncodeF64(value.(float64)))
		e = nil

	case CORE_DATA_STRING:
		i32, err := strconv.ParseInt(value.(string), 10, 32)
		if err != nil {
			return 0, err
		}
		v = uint64(i32)
		e = nil
	default:
		v = 0
		e = util.ErrorOfInvalid("type", "data")
	}
	if e != nil {
		return 0, e
	}
	return v, nil
}

func AsFloat64(data []byte) (float64, error) {
	meta, value, err := UnmarshalData(data)
	if err != nil {
		return 0, err
	}
	var v float64
	var e error
	switch meta {
	case CORE_DATA_INT8:
		v = float64(value.(int8))
		e = nil
	case CORE_DATA_UINT8:
		v = float64(value.(uint8))
		e = nil
	case CORE_DATA_INT16:
		v = float64(value.(int16))
		e = nil
	case CORE_DATA_UINT16:
		v = float64(value.(uint16))
		e = nil
	case CORE_DATA_INT32:
		v = float64(value.(int32))
		e = nil
	case CORE_DATA_UINT32:
		v = float64(value.(uint32))
		e = nil
	case CORE_DATA_INT64:
		v = float64(value.(int64))
		e = nil
	case CORE_DATA_UINT64:
		v = float64(value.(uint64))
		e = nil
	case CORE_DATA_FLOAT32:
		v = float64(value.(float32))
		e = nil
	case CORE_DATA_FLOAT64:
		v = value.(float64)
		e = nil
	case CORE_DATA_STRING:
		v, e = strconv.ParseFloat(value.(string), 32)
	default:
		v = 0
		e = util.ErrorOfInvalid("type", "data")
	}
	if e != nil {
		return 0, e
	}
	return v, nil
}

func (t DataType) FromString(s string) ([]byte, error) {
	var v interface{}
	var e error
	switch t {
	case CORE_DATA_BOOLEAN:
		b, err := strconv.ParseBool(s)
		if err != nil {
			return nil, err
		}
		v = b
		e = nil
	case CORE_DATA_INT8:
		i64, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return nil, err
		}
		v = uint8(i64)
		e = nil
	case CORE_DATA_UINT8:
		i64, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return nil, err
		}
		v = uint8(i64)
		e = nil
	case CORE_DATA_INT16:
		i64, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return nil, err
		}
		v = int16(i64)
		e = nil
	case CORE_DATA_UINT16:
		i64, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return nil, err
		}
		v = uint16(i64)
		e = nil
	case CORE_DATA_INT32:
		i64, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return nil, err
		}
		v = int32(i64)
		e = nil
	case CORE_DATA_UINT32:
		i64, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return nil, err
		}
		v = uint32(i64)
		e = nil
	case CORE_DATA_INT64:
		i64, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return nil, err
		}
		v = i64
		e = nil
	case CORE_DATA_UINT64:
		i64, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return nil, err
		}
		v = uint64(i64)
		e = nil
	case CORE_DATA_FLOAT32:
		f64, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return nil, err
		}
		v = float32(f64)
		e = nil
	case CORE_DATA_FLOAT64:
		f64, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return nil, err
		}
		v = f64
		e = nil
	case CORE_DATA_STRING:
		v = s
		e = nil
	case CORE_DATA_BYTES:
		hexData, err := hex.DecodeString(s)
		if err != nil {
			return nil, err
		}
		v = hexData
		e = nil
	default:
		v = 0
		e = util.ErrorOfInvalid("type", "data")
	}
	if e != nil {
		return nil, e
	}
	data, err := MarshalData(v)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (t DataType) FromUint64(value uint64) ([]byte, error) {
	var v interface{}
	var e error
	switch t {
	case CORE_DATA_INT8:
		v = uint8(value)
		e = nil
	case CORE_DATA_UINT8:
		v = uint8(value)
		e = nil
	case CORE_DATA_INT16:
		v = int16(value)
		e = nil
	case CORE_DATA_UINT16:
		v = uint16(value)
		e = nil
	case CORE_DATA_INT32:
		v = int32(value)
		e = nil
	case CORE_DATA_UINT32:
		v = uint32(value)
		e = nil
	case CORE_DATA_INT64:
		v = int64(value)
		e = nil
	case CORE_DATA_UINT64:
		v = value
		e = nil
	case CORE_DATA_FLOAT32:
		v = api.DecodeF32(value)
		e = nil
	case CORE_DATA_FLOAT64:
		v = api.DecodeF64(value)
		e = nil
	default:
		v = 0
		e = util.ErrorOfInvalid("type", "data")
	}
	if e != nil {
		return nil, e
	}
	data, err := MarshalData(v)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (t DataType) FromBytes(value []byte) ([]byte, error) {
	var v interface{}
	var e error
	switch t {
	case CORE_DATA_STRING:
		v = string(value)
		e = nil
	case CORE_DATA_BYTES:
		v = value
		e = nil
	default:
		v = 0
		e = util.ErrorOfInvalid("type", "data")
	}
	if e != nil {
		return nil, e
	}
	data, err := MarshalData(v)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetDataTypeByName(name string) (DataType, error) {
	n := strings.ToLower(name)

	if dataType, ok := dataTypes[n]; ok {
		return dataType, nil
	}
	return 0, util.ErrorOfUnknown("DataType", name)
}

func GetDataTypes() []DataType {
	list := make([]DataType, 0)
	for _, t := range dataTypes {
		list = append(list, t)
	}
	return list
}

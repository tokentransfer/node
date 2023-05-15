package core

import (
	"bytes"
	"fmt"
	"math"
	"sort"

	"github.com/tokentransfer/node/core/pb"
	"github.com/tokentransfer/node/util"
	"google.golang.org/protobuf/proto"
)

const (
	CORE_MESSAGE = DataType(10)
	CORE_DATA    = DataType(11)

	CORE_DATA_NULL    = DataType(50)
	CORE_DATA_BOOLEAN = DataType(51)

	CORE_DATA_INT8  = DataType(52)
	CORE_DATA_INT16 = DataType(53)
	CORE_DATA_INT32 = DataType(54)
	CORE_DATA_INT64 = DataType(55)

	CORE_DATA_UINT8  = DataType(56)
	CORE_DATA_UINT16 = DataType(57)
	CORE_DATA_UINT32 = DataType(58)
	CORE_DATA_UINT64 = DataType(59)

	CORE_DATA_FLOAT32 = DataType(60)
	CORE_DATA_FLOAT64 = DataType(61)

	CORE_DATA_STRING = DataType(62)
	CORE_DATA_BYTES  = DataType(63)

	CORE_DATA_LIST = DataType(64)
	CORE_DATA_MAP  = DataType(65)

	CORE_BLOCK                 = DataType(100)
	CORE_TRANSACTION           = DataType(101)
	CORE_RECEIPT               = DataType(102)
	CORE_TRANSACTION_WITH_DATA = DataType(103)
	CORE_MESSAGE_KEY           = DataType(104)

	// CORE_STATE         = byte(120)
	CORE_ACCOUNT_STATE = DataType(121)

	// CORE_INFO
	CORE_PAYLOAD_INFO  = DataType(141)
	CORE_CONTRACT_INFO = DataType(142)
	CORE_META_INFO     = DataType(143)
	CORE_TOKEN_INFO    = DataType(144)
	CORE_DATA_INFO     = DataType(145)
	CORE_PEER_INFO     = DataType(146)
	CORE_PAGE_INFO     = DataType(147)
	CORE_CODE_INFO     = DataType(148)
	CORE_USER_INFO     = DataType(149)
	CORE_GROUP_INFO    = DataType(150)
)

func init() {
	DataType(CORE_DATA_BOOLEAN).Create("boolean", true)

	DataType(CORE_DATA_INT8).Create("int8", true)
	DataType(CORE_DATA_INT16).Create("int16", true)
	DataType(CORE_DATA_INT32).Create("int32", true)
	DataType(CORE_DATA_INT64).Create("int64", true)

	DataType(CORE_DATA_UINT8).Create("uint8", true)
	DataType(CORE_DATA_UINT16).Create("uint16", true)
	DataType(CORE_DATA_UINT32).Create("uint32", true)
	DataType(CORE_DATA_UINT64).Create("uint16", true)

	DataType(CORE_DATA_FLOAT32).Create("float32", true)
	DataType(CORE_DATA_FLOAT64).Create("float64", true)

	DataType(CORE_DATA_STRING).Create("string", false)
	DataType(CORE_DATA_BYTES).Create("bytes", false)

	DataType(CORE_DATA_LIST).Create("list", false)
	DataType(CORE_DATA_MAP).Create("map", false)
}

func GetInfo(data []byte) string {
	if len(data) > 0 {
		meta := data[0]
		switch DataType(meta) {
		case CORE_MESSAGE:
			return "message"
		case CORE_DATA:
			return "data"

		case CORE_DATA_NULL:
			return "data_null"
		case CORE_DATA_BOOLEAN:
			return "data_boolean"

		case CORE_DATA_INT8:
			return "data_int8"
		case CORE_DATA_INT16:
			return "data_int16"
		case CORE_DATA_INT32:
			return "data_int32"
		case CORE_DATA_INT64:
			return "data_int64"
		case CORE_DATA_UINT8:
			return "data_uint8"
		case CORE_DATA_UINT16:
			return "data_uint16"
		case CORE_DATA_UINT32:
			return "data_uint32"
		case CORE_DATA_UINT64:
			return "data_uint64"

		case CORE_DATA_FLOAT32:
			return "data_float32"
		case CORE_DATA_FLOAT64:
			return "data_float64"

		case CORE_DATA_STRING:
			return "data_string"
		case CORE_DATA_BYTES:
			return "data_bytes"

		case CORE_DATA_LIST:
			return "data_list"
		case CORE_DATA_MAP:
			return "data_map"

		case CORE_BLOCK:
			return "block"
		case CORE_TRANSACTION:
			return "transaction"
		case CORE_RECEIPT:
			return "receipt"
		case CORE_TRANSACTION_WITH_DATA:
			return "transaction_with_data"
		case CORE_MESSAGE_KEY:
			return "message_key"

		case CORE_ACCOUNT_STATE:
			return "account_state"

		case CORE_PAYLOAD_INFO:
			return "payload_info"
		case CORE_CONTRACT_INFO:
			return "contract_info"
		case CORE_META_INFO:
			return "meta_info"
		case CORE_TOKEN_INFO:
			return "token_info"
		case CORE_DATA_INFO:
			return "data_info"
		case CORE_PEER_INFO:
			return "peer_info"
		case CORE_PAGE_INFO:
			return "page_info"
		case CORE_CODE_INFO:
			return "code_info"
		case CORE_USER_INFO:
			return "user_info"
		case CORE_GROUP_INFO:
			return "group_info"
		default:
			return "unknown"
		}
	}
	return ""
}

func Clone(t proto.Message) (proto.Message, error) {
	data, err := Marshal(t)
	if err != nil {
		return nil, err
	}
	_, o, err := Unmarshal(data)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func MarshalData(data interface{}) ([]byte, error) {
	if data == nil {
		return []byte{byte(CORE_DATA_NULL)}, nil
	}

	var meta DataType
	var bs []byte
	switch data := data.(type) {
	case bool:
		meta = CORE_DATA_BOOLEAN
		bs = make([]byte, 1)
		if data {
			bs[0] = 1
		} else {
			bs[0] = 0
		}
	case *bool:
		meta = CORE_DATA_BOOLEAN
		bs = make([]byte, 1)
		if *data {
			bs[0] = 1
		} else {
			bs[0] = 0
		}
	case int8:
		meta = CORE_DATA_INT8
		bs = make([]byte, 1)
		bs[0] = byte(data)
	case *int8:
		meta = CORE_DATA_INT8
		bs = make([]byte, 1)
		bs[0] = byte(*data)
	case uint8:
		meta = CORE_DATA_UINT8
		bs = make([]byte, 1)
		bs[0] = byte(data)
	case *uint8:
		meta = CORE_DATA_UINT8
		bs = make([]byte, 1)
		bs[0] = byte(*data)
	case int16:
		meta = CORE_DATA_INT16
		bs = make([]byte, 2)
		util.BYTE_ORDER.PutUint16(bs, uint16(data))
	case *int16:
		meta = CORE_DATA_INT16
		bs = make([]byte, 2)
		util.BYTE_ORDER.PutUint16(bs, uint16(*data))
	case uint16:
		meta = CORE_DATA_UINT16
		bs = make([]byte, 2)
		util.BYTE_ORDER.PutUint16(bs, data)
	case *uint16:
		meta = CORE_DATA_UINT16
		bs = make([]byte, 2)
		util.BYTE_ORDER.PutUint16(bs, *data)
	case int32:
		meta = CORE_DATA_INT32
		bs = make([]byte, 4)
		util.BYTE_ORDER.PutUint32(bs, uint32(data))
	case *int32:
		meta = CORE_DATA_INT32
		bs = make([]byte, 4)
		util.BYTE_ORDER.PutUint32(bs, uint32(*data))
	case uint32:
		meta = CORE_DATA_UINT32
		bs = make([]byte, 4)
		util.BYTE_ORDER.PutUint32(bs, data)
	case *uint32:
		meta = CORE_DATA_UINT32
		bs = make([]byte, 4)
		util.BYTE_ORDER.PutUint32(bs, *data)
	case int64:
		meta = CORE_DATA_INT64
		bs = make([]byte, 8)
		util.BYTE_ORDER.PutUint64(bs, uint64(data))
	case *int64:
		meta = CORE_DATA_INT64
		bs = make([]byte, 8)
		util.BYTE_ORDER.PutUint64(bs, uint64(*data))
	case uint64:
		meta = CORE_DATA_UINT64
		bs = make([]byte, 8)
		util.BYTE_ORDER.PutUint64(bs, data)
	case *uint64:
		meta = CORE_DATA_UINT64
		bs = make([]byte, 8)
		util.BYTE_ORDER.PutUint64(bs, *data)
	case float32:
		meta = CORE_DATA_FLOAT32
		bs = make([]byte, 4)
		util.BYTE_ORDER.PutUint32(bs, math.Float32bits(data))
	case *float32:
		meta = CORE_DATA_FLOAT32
		bs = make([]byte, 4)
		util.BYTE_ORDER.PutUint32(bs, math.Float32bits(*data))
	case float64:
		meta = CORE_DATA_FLOAT64
		bs = make([]byte, 8)
		util.BYTE_ORDER.PutUint64(bs, math.Float64bits(data))
	case *float64:
		meta = CORE_DATA_FLOAT64
		bs = make([]byte, 8)
		util.BYTE_ORDER.PutUint64(bs, math.Float64bits(*data))
	case string:
		meta = CORE_DATA_STRING
		bytes := []byte(data)
		bs = make([]byte, len(bytes))
		copy(bs, bytes)
	case *string:
		meta = CORE_DATA_STRING
		bytes := []byte(*data)
		bs = make([]byte, len(bytes))
		copy(bs, bytes)
	case []byte:
		meta = CORE_DATA_BYTES
		bs = make([]byte, len(data))
		copy(bs, data)
	case *[]byte:
		meta = CORE_DATA_BYTES
		bs = make([]byte, len(*data))
		copy(bs, *data)

	case []interface{}:
		list := &pb.DataList{}
		for _, item := range data {
			data, err := MarshalData(item)
			if err != nil {
				return nil, err
			}
			list.List = append(list.List, &pb.Data{Bytes: data})
		}
		return Marshal(list)
	case *[]interface{}:
		list := &pb.DataList{}
		for _, item := range *data {
			data, err := MarshalData(item)
			if err != nil {
				return nil, err
			}
			list.List = append(list.List, &pb.Data{Bytes: data})
		}
		return Marshal(list)
	case map[string]interface{}:
		list := make([]string, 0)
		for k, _ := range data {
			list = append(list, k)
		}
		l, keys := GetList(list)
		m := &pb.DataMap{Map: make([]*pb.DataEntry, l)}
		for i := 0; i < l; i++ {
			k := keys[i]
			v := data[k]
			d, err := MarshalData(v)
			if err != nil {
				return nil, err
			}
			m.Map[i] = &pb.DataEntry{
				Name: k,
				Value: &pb.Data{
					Bytes: d,
				},
			}
		}
		return Marshal(m)
	case *map[string]interface{}:
		list := make([]string, 0)
		for k := range *data {
			list = append(list, k)
		}
		l, keys := GetList(list)
		m := &pb.DataMap{Map: make([]*pb.DataEntry, l)}
		for i := 0; i < l; i++ {
			k := keys[i]
			v := (*data)[k]
			d, err := MarshalData(v)
			if err != nil {
				return nil, err
			}
			m.Map[i] = &pb.DataEntry{
				Name: k,
				Value: &pb.Data{
					Bytes: d,
				},
			}
		}
		return Marshal(m)
	case proto.Message:
		return Marshal(data)
	default:
		err := util.ErrorOfInvalid("data type", fmt.Sprintf("%v", data))
		return nil, err
	}
	return append([]byte{byte(meta)}, bs...), nil
}

func UnmarshalData(data []byte) (DataType, interface{}, error) {
	if len(data) == 0 {
		return 0, nil, util.ErrorOfInvalid("data", "empty")
	}
	if len(data) == 1 {
		return DataType(data[0]), nil, nil
	}
	var o interface{}

	meta := DataType(data[0])
	bs := data[1:]
	switch meta {
	case CORE_DATA:
		meta, msg, err := UnmarshalData(data)
		if err != nil {
			return 0, nil, err
		}
		if meta != CORE_DATA {
			return 0, nil, util.ErrorOfInvalid("data", "data")
		}
		d := msg.(*pb.Data)
		_, o, err := UnmarshalData(d.Bytes)
		if err != nil {
			return 0, nil, err
		}
		return meta, o, nil
	case CORE_DATA_NULL:
		return meta, nil, nil
	case CORE_DATA_BOOLEAN:
		if len(bs) != 1 {
			return 0, nil, util.ErrorOfInvalid("data", "boolean")
		}
		o = (bs[0] != 0)
	case CORE_DATA_INT8:
		if len(bs) != 1 {
			return 0, nil, util.ErrorOfInvalid("data", "int8")
		}
		o = int8(bs[0])
	case CORE_DATA_UINT8:
		if len(bs) != 1 {
			return 0, nil, util.ErrorOfInvalid("data", "uint8")
		}
		o = uint8(bs[0])
	case CORE_DATA_INT16:
		if len(bs) != 2 {
			return 0, nil, util.ErrorOfInvalid("data", "int16")
		}
		o = int16(util.BYTE_ORDER.Uint16(bs))
	case CORE_DATA_UINT16:
		if len(bs) != 2 {
			return 0, nil, util.ErrorOfInvalid("data", "uint16")
		}
		o = util.BYTE_ORDER.Uint16(bs)
	case CORE_DATA_INT32:
		if len(bs) != 4 {
			return 0, nil, util.ErrorOfInvalid("data", "int32")
		}
		o = int32(util.BYTE_ORDER.Uint32(bs))
	case CORE_DATA_UINT32:
		if len(bs) != 4 {
			return 0, nil, util.ErrorOfInvalid("data", "uint32")
		}
		o = util.BYTE_ORDER.Uint32(bs)
	case CORE_DATA_INT64:
		if len(bs) != 8 {
			return 0, nil, util.ErrorOfInvalid("data", "int64")
		}
		o = int64(util.BYTE_ORDER.Uint64(bs))
	case CORE_DATA_UINT64:
		if len(bs) != 8 {
			return 0, nil, util.ErrorOfInvalid("data", "uint64")
		}
		o = util.BYTE_ORDER.Uint64(bs)
	case CORE_DATA_FLOAT32:
		if len(bs) != 4 {
			return 0, nil, util.ErrorOfInvalid("data", "float32")
		}
		o = math.Float32frombits(util.BYTE_ORDER.Uint32(bs))
	case CORE_DATA_FLOAT64:
		if len(bs) != 8 {
			return 0, nil, util.ErrorOfInvalid("data", "float64")
		}
		o = math.Float64frombits(util.BYTE_ORDER.Uint64(bs))
	case CORE_DATA_STRING:
		d := make([]byte, len(bs))
		copy(d, bs)
		o = string(d)
	case CORE_DATA_BYTES:
		d := make([]byte, len(bs))
		copy(d, bs)
		o = d
	case CORE_DATA_LIST:
		t, msg, err := Unmarshal(data)
		if err != nil {
			return 0, nil, err
		}
		if t != CORE_DATA_LIST {
			return 0, nil, util.ErrorOfInvalid("data list", "data")
		}
		list := msg.(*pb.DataList)
		rlist := make([]interface{}, 0)
		for _, item := range list.List {
			_, sdata, err := UnmarshalData(item.Bytes)
			if err != nil {
				return 0, nil, err
			}
			rlist = append(rlist, sdata)
		}
		return meta, rlist, nil
	case CORE_DATA_MAP:
		t, msg, err := Unmarshal(data)
		if err != nil {
			return 0, nil, err
		}
		if t != CORE_DATA_MAP {
			return 0, nil, util.ErrorOfInvalid("data map", "data")
		}
		m := msg.(*pb.DataMap)
		rm := make(map[string]interface{})
		for _, v := range m.Map {
			_, sdata, err := UnmarshalData(v.Value.Bytes)
			if err != nil {
				return 0, nil, err
			}
			rm[v.Name] = sdata
		}
		return meta, rm, nil
	default:
		return Unmarshal(data)
	}
	return meta, o, nil
}

func Marshal(message proto.Message) ([]byte, error) {
	var meta DataType
	switch message.(type) {
	case *pb.Message:
		meta = CORE_MESSAGE
	case *pb.Data:
		meta = CORE_DATA

	case *pb.DataList:
		meta = CORE_DATA_LIST
	case *pb.DataMap:
		meta = CORE_DATA_MAP

	case *pb.Block:
		meta = CORE_BLOCK
	case *pb.Transaction:
		meta = CORE_TRANSACTION
	case *pb.Receipt:
		meta = CORE_RECEIPT
	case *pb.TransactionWithData:
		meta = CORE_TRANSACTION_WITH_DATA
	case *pb.MessageKey:
		meta = CORE_MESSAGE_KEY

	case *pb.AccountState:
		meta = CORE_ACCOUNT_STATE

	case *pb.PayloadInfo:
		meta = CORE_PAYLOAD_INFO
	case *pb.ContractInfo:
		meta = CORE_CONTRACT_INFO
	case *pb.MetaInfo:
		meta = CORE_META_INFO
	case *pb.TokenInfo:
		meta = CORE_TOKEN_INFO
	case *pb.DataInfo:
		meta = CORE_DATA_INFO
	case *pb.PeerInfo:
		meta = CORE_PEER_INFO
	case *pb.PageInfo:
		meta = CORE_PAGE_INFO
	case *pb.CodeInfo:
		meta = CORE_CODE_INFO
	case *pb.UserInfo:
		meta = CORE_USER_INFO
	case *pb.GroupInfo:
		meta = CORE_GROUP_INFO

	default:
		err := util.ErrorOfInvalid("type", "data")
		return nil, err
	}
	data, err := proto.Marshal(message)
	if err != nil {
		return nil, err
	}
	return append([]byte{byte(meta)}, data...), nil
}

func GetMeta(data []byte) DataType {
	if len(data) > 0 {
		return DataType(data[0])
	}
	return 0
}

func GetList(m []string) (int, []string) {
	list := make([][]byte, 0)
	for _, k := range m {
		list = append(list, []byte(k))
	}
	sort.Slice(list, func(i, j int) bool {
		li := len(list[i])
		lj := len(list[j])
		if li != lj {
			return li < lj
		}
		return bytes.Compare(list[i], list[j]) < 0
	})
	l := len(list)
	keys := make([]string, l)
	for i := 0; i < len(list); i++ {
		k := string(list[i])
		keys[i] = k
	}
	return l, keys
}

func GetSort(list [][]byte) (int, [][]byte) {
	sort.Slice(list, func(i, j int) bool {
		li := len(list[i])
		lj := len(list[j])
		if li != lj {
			return li < lj
		}
		return bytes.Compare(list[i], list[j]) < 0
	})
	return len(list), list
}

func Unmarshal(data []byte) (DataType, proto.Message, error) {
	if len(data) == 0 {
		return 0, nil, util.ErrorOfInvalid("data", "null")
	}

	meta := DataType(data[0])
	bs := data[1:]

	var msg proto.Message
	switch meta {
	case CORE_MESSAGE:
		msg = &pb.Message{}
	case CORE_DATA:
		msg = &pb.Data{}

	case CORE_DATA_LIST:
		msg = &pb.DataList{}
	case CORE_DATA_MAP:
		msg = &pb.DataMap{}

	case CORE_BLOCK:
		msg = &pb.Block{}
	case CORE_TRANSACTION:
		msg = &pb.Transaction{}
	case CORE_RECEIPT:
		msg = &pb.Receipt{}
	case CORE_TRANSACTION_WITH_DATA:
		msg = &pb.TransactionWithData{}
	case CORE_MESSAGE_KEY:
		msg = &pb.MessageKey{}

	case CORE_ACCOUNT_STATE:
		msg = &pb.AccountState{}

	case CORE_PAYLOAD_INFO:
		msg = &pb.PayloadInfo{}
	case CORE_CONTRACT_INFO:
		msg = &pb.ContractInfo{}
	case CORE_META_INFO:
		msg = &pb.MetaInfo{}
	case CORE_TOKEN_INFO:
		msg = &pb.TokenInfo{}
	case CORE_DATA_INFO:
		msg = &pb.DataInfo{}
	case CORE_PEER_INFO:
		msg = &pb.PeerInfo{}
	case CORE_PAGE_INFO:
		msg = &pb.PageInfo{}
	case CORE_CODE_INFO:
		msg = &pb.CodeInfo{}
	case CORE_USER_INFO:
		msg = &pb.UserInfo{}
	case CORE_GROUP_INFO:
		msg = &pb.GroupInfo{}

	default:
		err := util.ErrorOfInvalid("format", "data")
		return 0, nil, err
	}

	err := proto.Unmarshal(bs, msg)
	if err != nil {
		return 0, nil, err
	}
	return meta, msg, nil
}

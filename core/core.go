package core

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"math"
	"math/big"

	"github.com/tokentransfer/node/core/pb"
	"github.com/tokentransfer/node/util"
	"google.golang.org/protobuf/proto"
)

const (
	CORE_MESSAGE = byte(10)
	CORE_DATA    = byte(11)

	CORE_DATA_NULL    = byte(50)
	CORE_DATA_BOOLEAN = byte(51)

	CORE_DATA_INT8  = byte(52)
	CORE_DATA_INT16 = byte(53)
	CORE_DATA_INT32 = byte(54)
	CORE_DATA_INT64 = byte(55)

	CORE_DATA_UINT8  = byte(56)
	CORE_DATA_UINT16 = byte(57)
	CORE_DATA_UINT32 = byte(58)
	CORE_DATA_UINT64 = byte(59)

	CORE_DATA_FLOAT32 = byte(60)
	CORE_DATA_FLOAT64 = byte(61)

	CORE_DATA_STRING = byte(62)
	CORE_DATA_BYTES  = byte(63)

	CORE_DATA_LIST = byte(64)
	CORE_DATA_MAP  = byte(65)

	CORE_BLOCK                 = byte(100)
	CORE_TRANSACTION           = byte(101)
	CORE_RECEIPT               = byte(102)
	CORE_TRANSACTION_WITH_DATA = byte(103)
	CORE_MESSAGE_KEY           = byte(104)

	// CORE_STATE         = byte(120)
	CORE_ACCOUNT_STATE = byte(121)

	// CORE_INFO
	CORE_PAYLOAD_INFO  = byte(141)
	CORE_CONTRACT_INFO = byte(142)
	CORE_META_INFO     = byte(143)
	CORE_TOKEN_INFO    = byte(144)
	CORE_DATA_INFO     = byte(145)
	CORE_PEER_INFO     = byte(146)
	CORE_PAGE_INFO     = byte(147)
)

var SYSTEM_CODE = "TEST"

func GetInfo(data []byte) string {
	if len(data) > 0 {
		meta := data[0]
		switch meta {
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
		return []byte{CORE_DATA_NULL}, nil
	}

	var meta byte
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
	default:
		err := util.ErrorOfInvalid("type", "data")
		return nil, err
	}
	return append([]byte{meta}, bs...), nil
}

func UnmarshalData(data []byte) (byte, interface{}, error) {
	if len(data) == 0 {
		return 0, nil, util.ErrorOfInvalid("data", "empty")
	}
	if len(data) == 1 {
		return CORE_DATA_NULL, nil, nil
	}
	var o interface{}

	meta := data[0]
	bs := data[1:]
	switch meta {
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
		bytes := []byte(data)
		bs = make([]byte, len(bytes))
		copy(bs, bytes)
		o = string(bs)
	case CORE_DATA_BYTES:
		bs = make([]byte, len(data))
		copy(bs, data)
		o = bs
	default:
		err := util.ErrorOfInvalid("type", "data")
		return 0, nil, err
	}
	return meta, o, nil
}

func Marshal(message proto.Message) ([]byte, error) {
	var meta byte
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

	default:
		err := util.ErrorOfInvalid("type", "data")
		return nil, err
	}
	data, err := proto.Marshal(message)
	if err != nil {
		return nil, err
	}
	return append([]byte{meta}, data...), nil
}

func GetMeta(data []byte) byte {
	if len(data) > 0 {
		return data[0]
	}
	return 0
}

func Unmarshal(data []byte) (byte, proto.Message, error) {
	if len(data) > 0 {
		meta := data[0]
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

		default:
			err := errors.New("error data format")
			return 0, nil, err
		}

		err := proto.Unmarshal(bs, msg)
		if err != nil {
			return 0, nil, err
		}
		return meta, msg, nil
	}
	return 0, nil, errors.New("null data")
}

func WriteBytes(w io.Writer, b []byte) error {
	l := len(b)
	err := binary.Write(w, util.BYTE_ORDER, uint32(l))
	if err != nil {
		return err
	}
	if l > 0 {
		n, err := w.Write(b)
		if err != nil {
			return err
		}
		if n != l {
			return errors.New("error write")
		}
	}
	return nil
}

func ReadBytes(r io.Reader, maxSize uint32) ([]byte, error) {
	l := uint32(0)
	err := binary.Read(r, util.BYTE_ORDER, &l)
	if err != nil {
		return nil, err
	}
	if l > maxSize {
		return nil, fmt.Errorf("[ERROR] read bytes limit: %d > %d", l, maxSize)
	}
	b := make([]byte, l)
	if l > 0 {
		n, err := io.ReadFull(r, b)
		if err != nil {
			return nil, err
		}
		if n != int(l) {
			return nil, fmt.Errorf("[ERROR] error read: %d != %d", n, int(l))
		}
	}
	return b, nil
}

func GetIndex(address string) (uint64, error) {
	_, a, err := as.NewAccountFromAddress(address)
	if err != nil {
		return 0, err
	}
	data, err := a.MarshalBinary()
	if err != nil {
		return 0, err
	}
	n := new(big.Int)
	m := n.SetBytes(data)
	return m.Uint64(), nil
}

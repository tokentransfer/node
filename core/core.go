package core

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"math/big"

	"github.com/tokentransfer/node/core/pb"
	"google.golang.org/protobuf/proto"
)

const (
	CORE_MESSAGE = byte(10)

	CORE_BLOCK                 = byte(100)
	CORE_TRANSACTION           = byte(101)
	CORE_RECEIPT               = byte(102)
	CORE_TRANSACTION_WITH_DATA = byte(103)
	CORE_MESSAGE_KEY           = byte(104)

	// CORE_STATE         = byte(110)
	CORE_ACCOUNT_STATE  = byte(111)
	CORE_CURRENCY_STATE = byte(112)
)

var SYSTEM_CODE = "TEST"

func GetInfo(data []byte) string {
	if len(data) > 0 {
		meta := data[0]
		switch meta {
		case CORE_MESSAGE:
			return "message"
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
		case CORE_CURRENCY_STATE:
			return "currency_state"
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

func Marshal(message proto.Message) ([]byte, error) {
	var meta byte
	switch message.(type) {
	case *pb.Message:
		meta = CORE_MESSAGE
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
	case *pb.CurrencyState:
		meta = CORE_CURRENCY_STATE

	default:
		err := errors.New("error data type")
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
		case CORE_CURRENCY_STATE:
			msg = &pb.CurrencyState{}

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
	err := binary.Write(w, binary.LittleEndian, uint32(l))
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
	err := binary.Read(r, binary.LittleEndian, &l)
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

package util

import (
	"crypto/sha512"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"os"

	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/tjfoc/gmsm/sm3"
)

var (
	BYTE_ORDER = binary.LittleEndian
)

func ToArray(m *map[string]interface{}, key string) []interface{} {
	if m != nil {
		data := (*m)[key]
		if data != nil {
			list, ok := data.([]interface{})
			if ok {
				return list
			}
		}
	}
	return nil
}

func ToMap(m *map[string]interface{}, key string) map[string]interface{} {
	if m != nil {
		data := (*m)[key]
		if data != nil {
			sm, ok := data.(map[string]interface{})
			if ok {
				return sm
			}
		}
	}
	return nil
}

func Has(m *map[string]interface{}, key string) bool {
	if m != nil {
		_, ok := (*m)[key]
		return ok
	}
	return false
}

func ToUint64(m *map[string]interface{}, key string) uint64 {
	s := ToString(m, key)
	if len(s) > 0 {
		n := new(big.Int)
		_, ok := n.SetString(s, 10)
		if ok {
			return n.Uint64()
		}
	}
	return 0
}

func AsUint64(m *map[string]interface{}, key string) uint64 {
	item, ok := (*m)[key]
	if ok {
		i, ok := item.(float64)
		if ok {
			return uint64(i)
		}
	}
	return 0
}

func ToInt64(m *map[string]interface{}, key string) int64 {
	s := ToString(m, key)
	if len(s) > 0 {
		n := new(big.Int)
		_, ok := n.SetString(s, 10)
		if ok {
			return n.Int64()
		}
	}
	return 0
}

func ToString(m *map[string]interface{}, key string) string {
	item, ok := (*m)[key]
	if ok {
		s, ok := item.(string)
		if ok {
			return s
		}
	}
	return ""
}

func ToBoolean(m *map[string]interface{}, key string) bool {
	item, ok := (*m)[key]
	if ok {
		s, ok := item.(bool)
		if ok {
			return s
		}
	}
	return false
}

func PrintJSON(name string, o interface{}) {
	jsonBytes, err := json.MarshalIndent(o, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Println(name, string(jsonBytes))
}

func GetStringFromNode(n datamodel.Node, p string) string {
	n, err := n.LookupByString(p)
	if err != nil {
		return ""
	}
	s, err := n.AsString()
	if err != nil {
		return ""
	}
	return s
}

func GetIntFromNode(n datamodel.Node, p string) int64 {
	n, err := n.LookupByString(p)
	if err != nil {
		return 0
	}
	i, err := n.AsInt()
	if err != nil {
		return 0
	}
	return i
}

func GetBooleanFromNode(n datamodel.Node, p string) bool {
	n, err := n.LookupByString(p)
	if err != nil {
		return false
	}
	b, err := n.AsBool()
	if err != nil {
		return false
	}
	return b
}

func GetBytesFromNode(n datamodel.Node, p string) ([]byte, error) {
	n, err := n.LookupByString(p)
	if err != nil {
		return nil, err
	}
	data, err := n.AsBytes()
	if err != nil {
		return nil, nil
	}
	return data, nil
}

func GetListFromNode(n datamodel.Node, p string) datamodel.ListIterator {
	n, err := n.LookupByString(p)
	if err != nil {
		return nil
	}
	return n.ListIterator()
}

func Sha512Half(b []byte) []byte {
	h := sha512.New()
	h.Write(b)
	return h.Sum(nil)[:32]
}

func Sm3Half(b []byte) []byte {
	h := sm3.New()
	h.Write(b)
	return h.Sum(nil)[:32]
}

func WriteBytes(w io.Writer, b []byte) error {
	l := len(b)
	err := binary.Write(w, BYTE_ORDER, uint32(l))
	if err != nil {
		return err
	}
	if l > 0 {
		n, err := w.Write(b)
		if err != nil {
			return err
		}
		if n != l {
			return ErrorOfInvalid("data", fmt.Sprintf("%d != %d", n, l))
		}
	}
	return nil
}

func ReadBytes(r io.Reader) ([]byte, error) {
	return ReadBytesWith(r, 0)
}

func ReadBytesWith(r io.Reader, maxSize uint32) ([]byte, error) {
	l := uint32(0)
	err := binary.Read(r, BYTE_ORDER, &l)
	if err != nil {
		return nil, err
	}
	if maxSize > 0 && l > maxSize {
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

func ReadFile(p string) ([]byte, error) {
	f, err := os.Open(p)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return data, nil
}

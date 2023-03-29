package core

import (
	"encoding/hex"
	"fmt"

	"github.com/tokentransfer/node/util"

	libcore "github.com/tokentransfer/interfaces/core"
)

type Key libcore.Hash

func (k Key) String() string {
	return libcore.EncodeToString(k[:])
}

func (k Key) Address() string {
	s := k.String()
	if len(s) > 16 {
		return s[:16]
	}
	return s
}

func (k Key) Short() string {
	s := k.String()
	if len(s) > 8 {
		return s[:8]
	}
	return s
}

func (k Key) Equals(key Key) bool {
	return k.String() == key.String()
}

func isHexChar(c byte) bool {
	switch {
	case '0' <= c && c <= '9':
		return true
	case 'a' <= c && c <= 'f':
		return true
	case 'A' <= c && c <= 'F':
		return true
	}

	return false
}

func IsKey(s string) bool {
	if len(s) != 64 {
		return false
	}
	for i := 0; i < len(s); i++ {
		ok := isHexChar(s[i])
		if !ok {
			return false
		}
	}
	return true
}

func ParseKey(s string) (Key, error) {
	if len(s)%2 != 0 {
		return nil, util.ErrorOfInvalid("key length", fmt.Sprintf("%d", len(s)))
	}

	if b, err := hex.DecodeString(s); err != nil {
		return nil, err
	} else {
		return Key(b), nil
	}
}

func GetKey(s string) Key {
	b := []byte(s)
	return Key(b[:])
}

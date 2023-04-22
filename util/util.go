package util

import (
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"time"

	"github.com/shopspring/decimal"

	libcore "github.com/tokentransfer/interfaces/core"
)

const (
	MAX_VALUE       = int64(9000000000000000000)
	MAX_MEMO_LENGTH = (256 * 1024) //918744 //(900 * 1024)

	timeEpoch  int64  = 946684800 //January 1st, 2000 (00:00 UTC)
	timeFormat string = "2006-Jan-02 15:04:05"
)

// Represents a time as the number of seconds since the epoch: January 1st, 2000 (00:00 UTC)
type Time struct {
	T uint32
}

type HumanTime struct {
	Time
}

func NewTime(t uint32) Time {
	return Time{t}
}

func convertToTime(t time.Time) uint32 {
	return uint32(t.Sub(time.Unix(timeEpoch, 0)).Nanoseconds() / 1000000000)
}

func (t Time) Time() time.Time {
	return time.Unix(int64(t.T)+timeEpoch, 0)
}

func Now() *Time {
	return &Time{convertToTime(time.Now())}
}

// Accepts time formatted as 2006-Jan-02 15:04:05
func (t *Time) SetString(s string) error {
	v, err := time.Parse(timeFormat, s)
	if err != nil {
		return err
	}
	t.SetUint32(convertToTime(v))
	return nil
}

func (t *Time) SetUint32(n uint32) {
	t.T = n
}

func (t Time) Uint32() uint32 {
	return t.T
}

func (t Time) Human() *HumanTime {
	return &HumanTime{t}
}

// Returns time formatted as 2006-Jan-02 15:04:05
func (t Time) String() string {
	return t.Time().UTC().Format(timeFormat)
}

// Returns time formatted as 15:04:05
func (t Time) Short() string {
	return t.Time().UTC().Format("15:04:05")
}

type Value struct {
	value string
}

func NewValue(s string) (*Value, error) {
	_, err := decimal.NewFromString(s)
	if err != nil {
		return nil, err
	}
	return &Value{s}, nil
}

func (v *Value) Value() string {
	return v.value
}

func (v Value) Clone() Value {
	return Value{v.value}
}

func (v Value) ZeroClone() Value {
	return Value{"0"}
}

func (v Value) Subtract(a Value) (Value, error) {
	ai, err := decimal.NewFromString(v.value)
	if err != nil {
		return Value{"0"}, err
	}
	bi, err := decimal.NewFromString(a.value)
	if err != nil {
		return Value{"0"}, err
	}
	ci := ai.Sub(bi)
	return Value{ci.String()}, nil
}

func (v Value) Add(a Value) (Value, error) {
	ai, err := decimal.NewFromString(v.value)
	if err != nil {
		return Value{"0"}, err
	}
	bi, err := decimal.NewFromString(a.value)
	if err != nil {
		return Value{"0"}, err
	}
	ci := ai.Add(bi)
	return Value{ci.String()}, nil
}

func (v Value) Less(a Value) (bool, error) {
	ai, err := decimal.NewFromString(v.value)
	if err != nil {
		return false, err
	}
	bi, err := decimal.NewFromString(a.value)
	if err != nil {
		return false, err
	}
	return ai.LessThan(bi), nil
}

func (v *Value) IsNegative() (bool, error) {
	vi, err := decimal.NewFromString(v.value)
	if err != nil {
		return false, err
	}
	return vi.IsNegative(), nil
}

func (v *Value) IsPositive() (bool, error) {
	vi, err := decimal.NewFromString(v.value)
	if err != nil {
		return false, err
	}
	return vi.IsPositive(), nil
}

func (v *Value) IsZero() (bool, error) {
	vi, err := decimal.NewFromString(v.value)
	if err != nil {
		return false, err
	}
	return vi.IsZero(), nil
}

func (v *Value) String() string {
	return v.value
}

func (v *Value) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatUint(uint64(t.Uint32()), 10)), nil
}

func (t *Time) UnmarshalJSON(b []byte) error {
	n, err := strconv.ParseUint(string(b), 10, 32)
	if err != nil {
		return err
	}
	t.SetUint32(uint32(n))
	return nil
}

func (t HumanTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.String() + `"`), nil
}

func (t *HumanTime) UnmarshalJSON(b []byte) error {
	return t.SetString(string(b[1 : len(b)-1]))
}

func CheckValueInteger(key string, item interface{}) (bool, error) {
	if item != nil {
		switch item.(type) {
		case float64:
			v, iss := item.(float64)
			return iss && float64(int64(v)) == v, nil
		case int64:
			_, iss := item.(int64)
			return iss, nil
		case string:
			s, iss := item.(string)
			if !iss {
				return false, nil
			}
			z := big.NewInt(0)
			_, ok := z.SetString(s, 10)
			if !ok {
				return false, nil
			}

			maxV := big.NewInt(MAX_VALUE)
			if z.Cmp(maxV) > 0 {
				err := errors.New("error " + key + ", out of range")
				return false, err
			}

			minV := big.NewInt(-MAX_VALUE)
			if z.Cmp(minV) < 0 {
				err := errors.New("error " + key + ", out of range")
				return false, err
			}

			return true, nil
		}
	}
	return false, nil
}

func CheckTotal(key string, s string, decimals int64) (int64, error) {
	a, err := decimal.NewFromString(s)
	if err != nil {
		return 0, err
	}

	b, err := decimal.NewFromString("10")
	if err != nil {
		return 0, err
	}

	p := decimal.NewFromInt(decimals)
	z := a.Mul(b.Pow(p))

	if !z.IsInteger() {
		err := errors.New("error " + key)
		return 0, err
	}

	maxV := decimal.NewFromInt(MAX_VALUE)
	if !z.LessThanOrEqual(maxV) {
		err := errors.New("error " + key + ", out of range")
		return 0, err
	}

	minV := decimal.NewFromInt(-MAX_VALUE)
	if z.LessThan(minV) {
		err := errors.New("error " + key + ", out of range")
		return 0, err
	}

	return z.IntPart(), nil
}

func GetVersionKey(name string, version uint64, seperator string) string {
	return fmt.Sprintf("%s%s%d", name, seperator, version)
}

func GetTransactionKey(account libcore.Address, sequence uint64, seperator string) string {
	return fmt.Sprintf("%s%s%d", account.String(), seperator, sequence)
}

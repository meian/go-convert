package convert

import (
	"errors"
	"fmt"
)

var (
	ErrNil          = errors.New("nil cannot convert to number")
	ErrNoNumConvert = errors.New("cannot convert to number")
)

type RangeError struct {
	from  string
	to    string
	value interface{}
}

func (e RangeError) Error() string {
	return fmt.Sprintf("out of range value: from=%s, to=%s value=%v", e.from, e.to, e.value)
}

func (e *RangeError) Is(target error) bool {
	if _, ok := target.(RangeError); ok {
		return true
	}
	if _, ok := target.(*RangeError); ok {
		return true
	}
	return false
}

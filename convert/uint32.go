// Code generated by "gen-int -t uint32 -o convert/uint32.go"; DO NOT EDIT.
package convert

import (
	"errors"
	"strconv"
)

func ToUint32(v interface{}) (uint32, error) {
	if v == nil {
		return 0, ErrNil
	}
	switch vt := v.(type) {

	case int8:
		if vt < 0 {
			return 0, &RangeError{"int8", "uint32", vt}
		}
		return uint32(vt), nil
	case *int8:
		if vt == nil {
			return 0, ErrNil
		}
		if *vt < 0 {
			return 0, &RangeError{"*int8", "uint32", *vt}
		}
		return uint32(*vt), nil

	case int16:
		if vt < 0 {
			return 0, &RangeError{"int16", "uint32", vt}
		}
		return uint32(vt), nil
	case *int16:
		if vt == nil {
			return 0, ErrNil
		}
		if *vt < 0 {
			return 0, &RangeError{"*int16", "uint32", *vt}
		}
		return uint32(*vt), nil

	case int32:
		if vt < 0 {
			return 0, &RangeError{"int32", "uint32", vt}
		}
		return uint32(vt), nil
	case *int32:
		if vt == nil {
			return 0, ErrNil
		}
		if *vt < 0 {
			return 0, &RangeError{"*int32", "uint32", *vt}
		}
		return uint32(*vt), nil

	case int64:
		if vt < 0 || int64(maxUint32) < vt {
			return 0, &RangeError{"int64", "uint32", vt}
		}
		return uint32(vt), nil
	case *int64:
		if vt == nil {
			return 0, ErrNil
		}
		if *vt < 0 || int64(maxUint32) < *vt {
			return 0, &RangeError{"*int64", "uint32", *vt}
		}
		return uint32(*vt), nil

	case int:
		if vt < 0 || (!is64 || int64(maxUint32) < int64(vt)) {
			return 0, &RangeError{"int", "uint32", vt}
		}
		return uint32(vt), nil
	case *int:
		if vt == nil {
			return 0, ErrNil
		}
		if *vt < 0 || (!is64 || int64(maxUint32) < int64(*vt)) {
			return 0, &RangeError{"*int", "uint32", *vt}
		}
		return uint32(*vt), nil

	case uint8:
		return uint32(vt), nil
	case *uint8:
		if vt == nil {
			return 0, ErrNil
		}
		return uint32(*vt), nil

	case uint16:
		return uint32(vt), nil
	case *uint16:
		if vt == nil {
			return 0, ErrNil
		}
		return uint32(*vt), nil

	case uint32:
		return vt, nil
	case *uint32:
		if vt == nil {
			return 0, ErrNil
		}
		return *vt, nil

	case uint64:
		if uint64(maxUint32) < vt {
			return 0, &RangeError{"uint64", "uint32", vt}
		}
		return uint32(vt), nil
	case *uint64:
		if vt == nil {
			return 0, ErrNil
		}
		if uint64(maxUint32) < *vt {
			return 0, &RangeError{"*uint64", "uint32", *vt}
		}
		return uint32(*vt), nil

	case uint:
		if uint(maxUint32) < vt {
			return 0, &RangeError{"uint", "uint32", vt}
		}
		return uint32(vt), nil
	case *uint:
		if vt == nil {
			return 0, ErrNil
		}
		if uint(maxUint32) < *vt {
			return 0, &RangeError{"*uint", "uint32", *vt}
		}
		return uint32(*vt), nil

	case string:
		n, err := strconv.ParseUint(vt, 10, 32)
		if err != nil {
			if errors.Is(err, strconv.ErrRange) {
				err = &RangeError{"string", "uint32", vt}
			}
			return 0, err
		}
		r, err := ToUint32(n)
		if err, ok := err.(*RangeError); ok && err != nil {
			err.from = "string"
		}
		return r, err
	case *string:
		if vt == nil {
			return 0, ErrNil
		}
		n, err := strconv.ParseUint(*vt, 10, 32)
		if err != nil {
			if errors.Is(err, strconv.ErrRange) {
				err = &RangeError{"*string", "uint32", *vt}
			}
			return 0, err
		}
		r, err := ToUint32(n)
		if err, ok := err.(*RangeError); ok && err != nil {
			err.from = "*string"
		}
		return r, err

	}
	return 0, ErrNoNumConvert
}

func ToUint32P(v interface{}) (*uint32, error) {
	if v == nil {
		return nil, nil
	}
	switch vt := v.(type) {

	case int8:
		if vt < 0 {
			return nil, &RangeError{"int8", "uint32", vt}
		}
		casted := uint32(vt)
		return &casted, nil
	case *int8:
		if vt == nil {
			return nil, nil
		}
		if *vt < 0 {
			return nil, &RangeError{"*int8", "uint32", *vt}
		}
		casted := uint32(*vt)
		return &casted, nil

	case int16:
		if vt < 0 {
			return nil, &RangeError{"int16", "uint32", vt}
		}
		casted := uint32(vt)
		return &casted, nil
	case *int16:
		if vt == nil {
			return nil, nil
		}
		if *vt < 0 {
			return nil, &RangeError{"*int16", "uint32", *vt}
		}
		casted := uint32(*vt)
		return &casted, nil

	case int32:
		if vt < 0 {
			return nil, &RangeError{"int32", "uint32", vt}
		}
		casted := uint32(vt)
		return &casted, nil
	case *int32:
		if vt == nil {
			return nil, nil
		}
		if *vt < 0 {
			return nil, &RangeError{"*int32", "uint32", *vt}
		}
		casted := uint32(*vt)
		return &casted, nil

	case int64:
		if vt < 0 || int64(maxUint32) < vt {
			return nil, &RangeError{"int64", "uint32", vt}
		}
		casted := uint32(vt)
		return &casted, nil
	case *int64:
		if vt == nil {
			return nil, nil
		}
		if *vt < 0 || int64(maxUint32) < *vt {
			return nil, &RangeError{"*int64", "uint32", *vt}
		}
		casted := uint32(*vt)
		return &casted, nil

	case int:
		if vt < 0 || (!is64 || int64(maxUint32) < int64(vt)) {
			return nil, &RangeError{"int", "uint32", vt}
		}
		casted := uint32(vt)
		return &casted, nil
	case *int:
		if vt == nil {
			return nil, nil
		}
		if *vt < 0 || (!is64 || int64(maxUint32) < int64(*vt)) {
			return nil, &RangeError{"*int", "uint32", *vt}
		}
		casted := uint32(*vt)
		return &casted, nil

	case uint8:
		casted := uint32(vt)
		return &casted, nil
	case *uint8:
		if vt == nil {
			return nil, nil
		}
		casted := uint32(*vt)
		return &casted, nil

	case uint16:
		casted := uint32(vt)
		return &casted, nil
	case *uint16:
		if vt == nil {
			return nil, nil
		}
		casted := uint32(*vt)
		return &casted, nil

	case uint32:
		return &vt, nil
	case *uint32:
		return vt, nil

	case uint64:
		if uint64(maxUint32) < vt {
			return nil, &RangeError{"uint64", "uint32", vt}
		}
		casted := uint32(vt)
		return &casted, nil
	case *uint64:
		if vt == nil {
			return nil, nil
		}
		if uint64(maxUint32) < *vt {
			return nil, &RangeError{"*uint64", "uint32", *vt}
		}
		casted := uint32(*vt)
		return &casted, nil

	case uint:
		if uint(maxUint32) < vt {
			return nil, &RangeError{"uint", "uint32", vt}
		}
		casted := uint32(vt)
		return &casted, nil
	case *uint:
		if vt == nil {
			return nil, nil
		}
		if uint(maxUint32) < *vt {
			return nil, &RangeError{"*uint", "uint32", *vt}
		}
		casted := uint32(*vt)
		return &casted, nil

	case string:
		n, err := strconv.ParseUint(vt, 10, 32)
		if err != nil {
			if errors.Is(err, strconv.ErrRange) {
				err = &RangeError{"string", "*uint32", vt}
			}
			return nil, err
		}
		r, err := ToUint32P(n)
		if err, ok := err.(*RangeError); ok && err != nil {
			err.from = "string"
		}
		return r, err
	case *string:
		if vt == nil {
			return nil, nil
		}
		n, err := strconv.ParseUint(*vt, 10, 32)
		if err != nil {
			if errors.Is(err, strconv.ErrRange) {
				err = &RangeError{"*string", "*uint32", *vt}
			}
			return nil, err
		}
		r, err := ToUint32P(n)
		if err, ok := err.(*RangeError); ok && err != nil {
			err.from = "*string"
		}
		return r, err

	}
	return nil, ErrNoNumConvert
}

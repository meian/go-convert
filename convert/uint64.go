// Code generated by "gen-int -t uint64 -o convert/uint64.go"; DO NOT EDIT.
package convert

import (
	"errors"
	"strconv"
)

func ToUint64(v interface{}) (uint64, error) {
	if v == nil {
		return 0, ErrNil
	}
	switch vt := v.(type) {

	case int8:
		if vt < 0 {
			return 0, &RangeError{"int8", "uint64", vt}
		}
		return uint64(vt), nil
	case *int8:
		if vt == nil {
			return 0, ErrNil
		}
		if *vt < 0 {
			return 0, &RangeError{"*int8", "uint64", *vt}
		}
		return uint64(*vt), nil

	case int16:
		if vt < 0 {
			return 0, &RangeError{"int16", "uint64", vt}
		}
		return uint64(vt), nil
	case *int16:
		if vt == nil {
			return 0, ErrNil
		}
		if *vt < 0 {
			return 0, &RangeError{"*int16", "uint64", *vt}
		}
		return uint64(*vt), nil

	case int32:
		if vt < 0 {
			return 0, &RangeError{"int32", "uint64", vt}
		}
		return uint64(vt), nil
	case *int32:
		if vt == nil {
			return 0, ErrNil
		}
		if *vt < 0 {
			return 0, &RangeError{"*int32", "uint64", *vt}
		}
		return uint64(*vt), nil

	case int64:
		if vt < 0 {
			return 0, &RangeError{"int64", "uint64", vt}
		}
		return uint64(vt), nil
	case *int64:
		if vt == nil {
			return 0, ErrNil
		}
		if *vt < 0 {
			return 0, &RangeError{"*int64", "uint64", *vt}
		}
		return uint64(*vt), nil

	case int:
		if vt < 0 {
			return 0, &RangeError{"int", "uint64", vt}
		}
		return uint64(vt), nil
	case *int:
		if vt == nil {
			return 0, ErrNil
		}
		if *vt < 0 {
			return 0, &RangeError{"*int", "uint64", *vt}
		}
		return uint64(*vt), nil

	case uint8:
		return uint64(vt), nil
	case *uint8:
		if vt == nil {
			return 0, ErrNil
		}
		return uint64(*vt), nil

	case uint16:
		return uint64(vt), nil
	case *uint16:
		if vt == nil {
			return 0, ErrNil
		}
		return uint64(*vt), nil

	case uint32:
		return uint64(vt), nil
	case *uint32:
		if vt == nil {
			return 0, ErrNil
		}
		return uint64(*vt), nil

	case uint64:
		return vt, nil
	case *uint64:
		if vt == nil {
			return 0, ErrNil
		}
		return *vt, nil

	case uint:
		if uint(maxUint64) < vt {
			return 0, &RangeError{"uint", "uint64", vt}
		}
		return uint64(vt), nil
	case *uint:
		if vt == nil {
			return 0, ErrNil
		}
		if uint(maxUint64) < *vt {
			return 0, &RangeError{"*uint", "uint64", *vt}
		}
		return uint64(*vt), nil

	case string:
		n, err := strconv.ParseUint(vt, 10, 64)
		if err != nil {
			if errors.Is(err, strconv.ErrRange) {
				err = &RangeError{"string", "uint64", vt}
			}
			return 0, err
		}
		r, err := ToUint64(n)
		if err, ok := err.(*RangeError); ok && err != nil {
			err.from = "string"
		}
		return r, err
	case *string:
		if vt == nil {
			return 0, ErrNil
		}
		n, err := strconv.ParseUint(*vt, 10, 64)
		if err != nil {
			if errors.Is(err, strconv.ErrRange) {
				err = &RangeError{"*string", "uint64", *vt}
			}
			return 0, err
		}
		r, err := ToUint64(n)
		if err, ok := err.(*RangeError); ok && err != nil {
			err.from = "*string"
		}
		return r, err

	}
	return 0, ErrNoNumConvert
}

func ToUint64P(v interface{}) (*uint64, error) {
	if v == nil {
		return nil, nil
	}
	switch vt := v.(type) {

	case int8:
		if vt < 0 {
			return nil, &RangeError{"int8", "uint64", vt}
		}
		casted := uint64(vt)
		return &casted, nil
	case *int8:
		if vt == nil {
			return nil, nil
		}
		if *vt < 0 {
			return nil, &RangeError{"*int8", "uint64", *vt}
		}
		casted := uint64(*vt)
		return &casted, nil

	case int16:
		if vt < 0 {
			return nil, &RangeError{"int16", "uint64", vt}
		}
		casted := uint64(vt)
		return &casted, nil
	case *int16:
		if vt == nil {
			return nil, nil
		}
		if *vt < 0 {
			return nil, &RangeError{"*int16", "uint64", *vt}
		}
		casted := uint64(*vt)
		return &casted, nil

	case int32:
		if vt < 0 {
			return nil, &RangeError{"int32", "uint64", vt}
		}
		casted := uint64(vt)
		return &casted, nil
	case *int32:
		if vt == nil {
			return nil, nil
		}
		if *vt < 0 {
			return nil, &RangeError{"*int32", "uint64", *vt}
		}
		casted := uint64(*vt)
		return &casted, nil

	case int64:
		if vt < 0 {
			return nil, &RangeError{"int64", "uint64", vt}
		}
		casted := uint64(vt)
		return &casted, nil
	case *int64:
		if vt == nil {
			return nil, nil
		}
		if *vt < 0 {
			return nil, &RangeError{"*int64", "uint64", *vt}
		}
		casted := uint64(*vt)
		return &casted, nil

	case int:
		if vt < 0 {
			return nil, &RangeError{"int", "uint64", vt}
		}
		casted := uint64(vt)
		return &casted, nil
	case *int:
		if vt == nil {
			return nil, nil
		}
		if *vt < 0 {
			return nil, &RangeError{"*int", "uint64", *vt}
		}
		casted := uint64(*vt)
		return &casted, nil

	case uint8:
		casted := uint64(vt)
		return &casted, nil
	case *uint8:
		if vt == nil {
			return nil, nil
		}
		casted := uint64(*vt)
		return &casted, nil

	case uint16:
		casted := uint64(vt)
		return &casted, nil
	case *uint16:
		if vt == nil {
			return nil, nil
		}
		casted := uint64(*vt)
		return &casted, nil

	case uint32:
		casted := uint64(vt)
		return &casted, nil
	case *uint32:
		if vt == nil {
			return nil, nil
		}
		casted := uint64(*vt)
		return &casted, nil

	case uint64:
		return &vt, nil
	case *uint64:
		return vt, nil

	case uint:
		if uint(maxUint64) < vt {
			return nil, &RangeError{"uint", "uint64", vt}
		}
		casted := uint64(vt)
		return &casted, nil
	case *uint:
		if vt == nil {
			return nil, nil
		}
		if uint(maxUint64) < *vt {
			return nil, &RangeError{"*uint", "uint64", *vt}
		}
		casted := uint64(*vt)
		return &casted, nil

	case string:
		n, err := strconv.ParseUint(vt, 10, 64)
		if err != nil {
			if errors.Is(err, strconv.ErrRange) {
				err = &RangeError{"string", "*uint64", vt}
			}
			return nil, err
		}
		r, err := ToUint64P(n)
		if err, ok := err.(*RangeError); ok && err != nil {
			err.from = "string"
		}
		return r, err
	case *string:
		if vt == nil {
			return nil, nil
		}
		n, err := strconv.ParseUint(*vt, 10, 64)
		if err != nil {
			if errors.Is(err, strconv.ErrRange) {
				err = &RangeError{"*string", "*uint64", *vt}
			}
			return nil, err
		}
		r, err := ToUint64P(n)
		if err, ok := err.(*RangeError); ok && err != nil {
			err.from = "*string"
		}
		return r, err

	}
	return nil, ErrNoNumConvert
}
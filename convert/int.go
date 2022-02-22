package convert

func ToInt(v interface{}) (int, error) {
	if v == nil {
		return 0, ErrNil
	}
	if is64 {
		vt, err := ToInt64(v)
		if err != nil {
			if err, ok := err.(*RangeError); ok && err != nil {
				err.to = "int"
			}
			return 0, err
		}
		return int(vt), nil
	}
	vt, err := ToInt32(v)
	if err != nil {
		if err, ok := err.(*RangeError); ok && err != nil {
			err.to = "int"
		}
		return 0, err
	}
	return int(vt), nil
}

func ToIntP(v interface{}) (*int, error) {
	if v == nil {
		return nil, nil
	}
	if is64 {
		vt, err := ToInt64P(v)
		if err != nil {
			if err, ok := err.(*RangeError); ok && err != nil {
				err.to = "*int"
			}
			return nil, err
		}
		if vt == nil {
			return nil, nil
		}
		casted := int(*vt)
		return &casted, nil
	}
	vt, err := ToInt32P(v)
	if err != nil {
		if err, ok := err.(*RangeError); ok && err != nil {
			err.to = "*int"
		}
		return nil, err
	}
	if vt == nil {
		return nil, nil
	}
	casted := int(*vt)
	return &casted, nil
}

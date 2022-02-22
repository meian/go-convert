package convert

func ToUint(v interface{}) (uint, error) {
	if v == nil {
		return 0, ErrNil
	}
	if is64 {
		vt, err := ToUint64(v)
		if err != nil {
			if err, ok := err.(*RangeError); ok && err != nil {
				err.to = "uint"
			}
			return 0, err
		}
		return uint(vt), nil
	}
	vt, err := ToUint32(v)
	if err != nil {
		if err, ok := err.(*RangeError); ok && err != nil {
			err.to = "uint"
		}
		return 0, err
	}
	return uint(vt), nil
}

func ToUintP(v interface{}) (*uint, error) {
	if v == nil {
		return nil, nil
	}
	if is64 {
		vt, err := ToUint64P(v)
		if err != nil {
			if err, ok := err.(*RangeError); ok && err != nil {
				err.to = "*uint"
			}
			return nil, err
		}
		if vt == nil {
			return nil, nil
		}
		casted := uint(*vt)
		return &casted, nil
	}
	vt, err := ToUint32P(v)
	if err != nil {
		if err, ok := err.(*RangeError); ok && err != nil {
			err.to = "*uint"
		}
		return nil, err
	}
	if vt == nil {
		return nil, nil
	}
	casted := uint(*vt)
	return &casted, nil
}

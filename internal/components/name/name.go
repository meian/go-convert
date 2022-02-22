package name

import (
	"strconv"
	"strings"
)

type Name string

var (
	Signeds = []Name{
		Name("int8"),
		Name("int16"),
		Name("int32"),
		Name("int64"),
		Name("int"),
	}
	Unsigneds = []Name{
		Name("uint8"),
		Name("uint16"),
		Name("uint32"),
		Name("uint64"),
		Name("uint"),
	}
)

func ByName(name string) (Name, bool) {
	for _, n := range Signeds {
		if string(n) == name {
			return n, true
		}
	}
	for _, n := range Unsigneds {
		if string(n) == name {
			return n, true
		}
	}
	return Name(""), false
}

func (n Name) IsSigned() bool {
	return strings.HasPrefix(string(n), "int")
}

func (n Name) BitSize() int {
	bss := strings.TrimPrefix(strings.TrimPrefix(string(n), "u"), "int")
	bs, _ := strconv.Atoi(bss)
	return bs
}

func (n Name) UpperFirst() string {
	return strings.ToUpper(string(n)[:1]) + string(n)[1:]
}

func (n Name) MaxVar() string {
	return "max" + n.UpperFirst()
}

func (n Name) MinVar() string {
	return "min" + n.UpperFirst()
}

func (n Name) ToMethod() string {
	return "To" + n.UpperFirst()
}

func (n Name) ToPointerMethod() string {
	return n.ToMethod() + "P"
}

func (n Name) SizeEq(other Name) bool {
	if n == other {
		return true
	}
	return n.BitSize() == other.BitSize()
}

func (n Name) SizeLess(other Name) bool {
	if n == other {
		return false
	}
	nsn := n.BitSize()
	otn := other.BitSize()
	if nsn == 0 {
		return false
	}
	if otn == 0 {
		return true
	}
	return nsn < otn
}

func (n Name) SizeGreater(other Name) bool {
	if n == other {
		return false
	}
	nsn := n.BitSize()
	otn := other.BitSize()
	if nsn == 0 {
		return true
	}
	if otn == 0 {
		return false
	}
	return nsn > otn
}

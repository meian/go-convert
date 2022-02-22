package convert_test

import (
	"testing"

	"github.com/meian/go-convert/convert"
	"github.com/stretchr/testify/assert"
)

func TestToInt32(t *testing.T) {
	p := (*int)(nil)
	v, err := convert.ToInt32(p)
	assert.Error(t, err)
	assert.Zero(t, v)
}

package golang

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Age int

type Number interface {
	~int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

func Min[T Number](a, b T) T {
	if a < b {
		return a
	} else {
		return b
	}

}

func TestMin(t *testing.T) {
	assert.Equal(t, 3, Min(3, 5))
	assert.Equal(t, 2.5, Min(2.5, 3.7))
	assert.Equal(t, uint(10), Min(10, 20))
	assert.Equal(t, int16(-5), Min(-5, 0))
	assert.Equal(t, Age(30), Min(Age(30), Age(40)))
}
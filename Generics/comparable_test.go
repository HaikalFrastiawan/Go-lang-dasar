package golang

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func isSame[T comparable](value1, value2 T) bool {
	if value1 == value2 {
		return true
	} else {
		return false
	}
}

func TestIsSame(t *testing.T) {
	assert.True(t, isSame[string]("Hello", "Hello"))
	assert.True(t, isSame[int](42, 42))
	assert.True(t, isSame[bool](true,true))
}
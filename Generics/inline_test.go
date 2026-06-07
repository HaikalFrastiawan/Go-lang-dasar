package golang

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func FindMin[T interface{ int | int64 | float64 }](frist, second T) T {
	if frist < second {
		return frist
	} else {
		return second
	}
}

func TestFindMin(t *testing.T) {
	assert.Equal(t, 3, FindMin(3, 5))
	assert.Equal(t, 2.5, FindMin(2.5, 3.7))
	assert.Equal(t, int64(10), FindMin(int64(10), int64(20)))
}

func GetFirst[T []E, E any](slice T) E {
	first := slice[0]
	return first
}

func TestGetFirst(t *testing.T) {
	name := []string{"Alice", "Bob", "Charlie"}

	first := GetFirst(name)
	assert.Equal(t, "Alice", first)
}
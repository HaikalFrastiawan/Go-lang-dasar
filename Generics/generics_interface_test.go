package golang

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type GetterSetter[T any] interface {
	GetValue() T
	SetValue(value T)
}

func ChangeValue[T any](param GetterSetter[T], value T) T {
	param.SetValue(value)
	return param.GetValue()
}

type MyData[T any] struct {
	value T
}

func (d *MyData[T]) GetValue() T {
	return d.value
}

func (d *MyData[T]) SetValue(value T)  {
	d.value = value
}

func TestGenericInterface(t *testing.T) {
	data := MyData[string]{}
	restult := ChangeValue[string](&data, "Hello")
	assert.Equal(t, "Hello", restult)
	assert.Equal(t, "Hello", data.GetValue())
}

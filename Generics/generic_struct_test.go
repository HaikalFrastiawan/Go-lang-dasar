package golang

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Data[T any] struct {
	Frist  T
	Second T
}

func (d *Data[_])  SayHello(name string)string {
	return "Hello " + name
}

func (d *Data[T]) ChangeFirst(first T) T {
	d.Frist = first
	return d.Frist
}


func TestData(t *testing.T) {
	data := Data[string]{
		Frist:  "Hello",
		Second: "World",
	}
	fmt.Println(data)
}

func TestGenericMethods(t *testing.T) {
	data := Data[string]{
		Frist:  "Hello",
		Second: "World",
	}
	assert.Equal(t, "Hello John", data.SayHello("John"))
	assert.Equal(t, "Hi", data.ChangeFirst("Hi"))

}
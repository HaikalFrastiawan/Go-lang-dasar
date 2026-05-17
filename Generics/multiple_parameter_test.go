package golang

import (
	"fmt"
	"testing"
)

func MultipleParameter[T any, U any](param1 T, param2 U) {
	fmt.Println(param1)
	fmt.Println(param2)
}

func TestMultipleParameter(t *testing.T) {
	MultipleParameter[string, int]("Hello, World!", 42)

}
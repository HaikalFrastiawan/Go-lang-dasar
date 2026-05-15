package golang

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

//jika tidak menggunakan generic harus 2 kali membuat function dengan tipe data yang berbeda, namun dengan generic kita hanya perlu membuat 1 function saja yang bisa digunakan untuk berbagai tipe data

//func Length(param string) string{
//	fmt.Println(param)
//	return param
//}

// funct Length(param int) int{
//	fmt.Println(param)
//	return param
//}

//func Lenggth(param interface{}) interface{}{
//	fmt.Println(param)
//	return param
//}

func Length[T any](param T) T{
	fmt.Println(param)
	return param
}

func TestLength(t *testing.T) {

	var result string = Length[string]("Hello, World!")
	fmt.Println(result)

	var resultNumber int = Length[int](42)
	fmt.Println(resultNumber)

	assert.Equal(t, "Hello, World!", result)
	assert.Equal(t, 42, resultNumber)

}
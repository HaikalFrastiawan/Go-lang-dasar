package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//skip
func TestSkip(t *testing.T){
	if runtime.GOOS == "windows" {
		t.Skip("can not run on Windows")
	}
	//ini hanya esek
	result := HelloWorld("Haikal")
	require.Equal(t, "Hello Haikal", result , "Result must be 'Hello Haikal'")
	
}


//Require -> sama seperti failNow() ketika gagal kode bawahnya ga di eksekusi
func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Haikal")
	require.Equal(t, "Hello Haikal", result , "Result must be 'Hello Haikal'")
	//Require -> sama seperti failNow() ketika gagal kode bawahnya ga di eksekusi
	fmt.Println("Test HelloWorld with Rquire DOne")
}

//assertion
func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Haikal")
	assert.Equal(t, "Hello Haikal", result , "Result must be 'Hello Haikal'")
	fmt.Println("Test HelloWorld with Assert DOne")
	//kalo gagal lebih detail error nya
}

func TestHelloWorldHaikal(t *testing.T) {
	result := HelloWorld("Haikal")

	if result != "Hello Haikal" {
		t.Error("Result Must be 'Hello Haikal'")

	}

	fmt.Println("TestHelloWorld Done")
}

func TestHelloWorldFrastiawan(t *testing.T) {
	result := HelloWorld("Frastiawan")

	if result != "Hello Frastiawan" {
		t.Fatal("Result Must be 'Hello Frastiawan'")
	}
	fmt.Println("TestHelloWorldFrastiawan Done")
}

//runing test codenya
//go test : untuk menjalakan unit test
//go test -v : keseluruhan code
//go test -v -run (nama functionnya -> TestHelloWorld): ini hanya  1 fungsi

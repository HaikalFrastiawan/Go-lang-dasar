package helper

import "testing"


func TestHelloWorl(t *testing.T){
	result := HelloWorld("Haikal")

	if result != "Hello Haikal" {
		panic("Result is not 'Hello Haikal'")
	}
}

//runing test codenya
//go test : untuk menjalakan unit test
//go test -v : keseluruhan code 
//go test -v -run (nama functionnya -> TestHelloWorld): ini hanya  1 fungsi
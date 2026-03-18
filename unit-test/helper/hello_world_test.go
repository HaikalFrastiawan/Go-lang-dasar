package helper

import (
	"fmt"
	"testing"
)

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

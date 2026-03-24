package main_test

import (
	_ "embed"
	"fmt"
	"io/fs"
	"os"
	"testing"
)

//go:embed version.txt
var version string


func TestString(t *testing.T) {
	fmt.Println(version)
	
}

//go:embed logo.png
var logo []byte

func TestByte(t *testing.T) {
	err := os.WriteFile("logo_new.png", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}
	
}

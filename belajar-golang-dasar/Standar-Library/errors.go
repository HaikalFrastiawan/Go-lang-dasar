package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("Validation error")
	NotFoundError = errors.New("not found error")
)

func GetById(id string)error{
	if id == ""{
		return ValidationError
	}
	if id != "haikal"{
		return NotFoundError
	}

	//sukses
	return nil
}


func main() {
	err := GetById("")
	
	if err  != nil {
		if errors.Is(err, ValidationError){
			fmt.Println("Validation Error")
		} else if errors.Is(err, NotFoundError){
			fmt.Println("Not Found Error")
		} else {
			fmt.Println("unknow Error ")
		}
	}
}
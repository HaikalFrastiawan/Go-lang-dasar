package golangjson

import (

	"encoding/json"
	"fmt"

	"testing"
)

type Address struct {
	Street string
	Country string
	PostalCode string
}

type Customer struct {
	FirstName string
	MiddleName string
	LastName string
	Age int
	Married bool
	Hobbies []string
	Address []Address	
}



func TestJSONObject(t *testing.T) {
	customer := Customer{
		FirstName: "Haikal",
		MiddleName: "Frastiawan",
		LastName: "Putra",
		Age: 30,
		Married: true,
	}

	byte, _ := json.Marshal(customer)
	fmt.Println(string(byte))
}
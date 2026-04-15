package golangjson

import (

	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	FristName string
	MiddleName string
	LastName string
	Age int
	Married bool
}

func TestJSONObject(t *testing.T) {
	customer := Customer{
		FristName: "Haikal",
		MiddleName: "Frastiawan",
		LastName: "Putra",
		Age: 30,
		Married: true,
	}

	byte, _ := json.Marshal(customer)
	fmt.Println(string(byte))
}
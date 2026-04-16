package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray(t *testing.T) {
	customer := Customer{
		FirstName:  "Haikal",
		MiddleName: "Frastiawan",
		LastName:   "Putra",
		Age:        30,
		Married:    true,
		Hobbies:    []string{"Coding", "Gaming", "Traveling"},
	}

	byte, _ := json.Marshal(customer)
	fmt.Println(string(byte))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonSting := `{"FirstName":"Haikal","MiddleName":"Frastiawan","LastName":"Putra","Age":30,"Married":true,"Hobbies":["Coding","Gaming","Traveling"]}`

	jsonBytes := []byte(jsonSting)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.MiddleName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Age)
	fmt.Println(customer.Married)
	fmt.Println(customer.Hobbies)
}

func TestJSONArrayCompleks(t *testing.T) {
	customer := Customer{
		FirstName:  "Haikal",
		Address: []Address{
			{ Street: "Jl. Merdeka No. 123", Country: "Indonesia", PostalCode: "12345",
		},{
			Street: "Jl. Sudirman No. 456", Country: "Indonesia", PostalCode: "67890",
		},
	},
		
	
	}

	byte, _ := json.Marshal(customer)
	fmt.Println(string(byte))
}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonSting := `{"FirstName":"Haikal","MiddleName":"Frastiawan","LastName":"Putra","Age":30,"Married":true,"Hobbies":["Coding","Gaming","Traveling"], "Address":[{"Street":"Jl. Merdeka No. 123","Country":"Indonesia","PostalCode":"12345"},{"Street":"Jl. Sudirman No. 456","Country":"Indonesia","PostalCode":"67890"}]}`

	jsonBytes := []byte(jsonSting)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.MiddleName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Age)
	fmt.Println(customer.Married)
	fmt.Println(customer.Hobbies)
	fmt.Println(customer.Address)
}

func TestJSONOnlyArrayComplexDecode(t *testing.T) {
	jsonSting := `[{"Street":"Jl. Merdeka No. 123","Country":"Indonesia","PostalCode":"12345"},{"Street":"Jl. Sudirman No. 456","Country":"Indonesia","PostalCode":"67890"}]`

	jsonBytes := []byte(jsonSting)

	address := []Address{}

	err := json.Unmarshal(jsonBytes, &address)
	if err != nil {
		panic(err)
	}

	fmt.Println(address)

}
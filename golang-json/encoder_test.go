package golangjson

import (
	"encoding/json"
	"os"
	"testing"
)

func TestEncoder(t *testing.T) {
	writer, _ := os.Create("CustomerOut.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName: "Haikal",
		MiddleName: "Frastiawan",
		LastName: "Putra",
		Age: 30,
		Married: true,
	}
	encoder.Encode(customer)

}
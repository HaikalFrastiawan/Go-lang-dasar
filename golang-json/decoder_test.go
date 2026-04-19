package golangjson

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStringDecoder(t *testing.T) {
	reader, _  := os.Open("Customer.json")
	decoder := json.NewDecoder(reader)

	customer := &Customer{}
	decoder.Decode(customer)

	fmt.Println(customer)
}
package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

func TestJSONTag(t *testing.T) {
	product := Product {
		Id: "P001",
		Name: "Laptop",
		ImageUrl: "https://example.com/laptop.jpg",
	}
	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}

func TestJSONTagDecode(t *testing.T) {
	jsonString := `{"id":"P001","name":"Laptop","image_url":"https://example.com/laptop.jpg"}`

	jsonBytes := []byte(jsonString)
	product := &Product{}
	err := json.Unmarshal(jsonBytes, product)
	if err != nil {
		panic(err)
	}
	fmt.Println(product)
	fmt.Println(product.Id)
	fmt.Println(product.Name)
	fmt.Println(product.ImageUrl)
}
package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapDecode(t *testing.T) {
	jsonString :=  `{"id":"123", "name":"Haikal", "age":30, "married":true, "hobbies":["Coding", "Gaming"]}`
	jsonBytes := []byte(jsonString)
	
	var result map[string]interface{}
	err := json.Unmarshal(jsonBytes, &result)
	if err != nil {
		t.Errorf("Error unmarshaling JSON: %v", err)
	}

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["age"])
	fmt.Println(result["married"])
	fmt.Println(result["hobbies"])

}

func TestMapEncode(t *testing.T) {
	Product := map[string]interface{}{
		"id": "P001",
		"name": "Laptop",
		"price": 1500.00,
		"available": true,
		"tags": []string{"electronics", "computers"},
	}
	bytes, err := json.Marshal(Product)
	if err != nil {
		t.Errorf("Error marshaling JSON: %v", err)
	}
	fmt.Println(string(bytes))	

}
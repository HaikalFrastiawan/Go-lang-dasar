package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func LogJson(data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestEncode(t *testing.T) {
	LogJson("Haikal")
	LogJson(1)
	LogJson(true)
	LogJson([]string{"Haikal", "Frastiawan"})
}
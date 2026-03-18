package main

import (
    "encoding/json"
    "fmt"
)

type User struct {
    FullName string `json:"full_name"`
    Age      int    `json:"age"`
}

func main() {
    // 1. Marshalling (Struct -> JSON)
    user := User{"Haikal", 20}
    jsonData, _ := json.Marshal(user)
    fmt.Println(string(jsonData)) // {"full_name":"Haikal","age":20}

    // 2. Unmarshalling (JSON -> Struct)
    jsonString := `{"full_name":"Budi","age":25}`
    var userBaru User
    json.Unmarshal([]byte(jsonString), &userBaru)
    fmt.Println(userBaru.FullName) // Budi
}
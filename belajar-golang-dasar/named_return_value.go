package main

import "fmt"

func getCompleteName() (firstName , midleName, lastName string) {
	firstName = "Haikal"
	midleName = "Fras"
	lastName = "tiawan"

	return firstName, midleName, lastName
}

func main() {
	a, b, c := getCompleteName()
	fmt.Println(a,b,c)
}
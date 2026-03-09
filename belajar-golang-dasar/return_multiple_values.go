package main

import "fmt"

func getFullName()(string, string){
	return "Eko", "khanedy"
}


func main(){
	// fristname, lastName := getFullName()
	// fmt.Println(fristname,lastName)

	fristname, _ := getFullName()
	fmt.Println(fristname)
}
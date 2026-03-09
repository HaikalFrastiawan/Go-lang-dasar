package main

import "fmt"

//function parameter
func sayHelloTo(fristName string, lastname string){
	fmt.Println("Hello", fristName, lastname)
}


//function return value
func getHello(name string)string{
	hello := "Helo " + name
	return hello
}

func main (){
	sayHelloTo("Haikal","Frastiawan")
	result := getHello("Adit")
	fmt.Println(result)


	//kalo gamau nambah variabel
	fmt.Println(getHello("eko"))
}
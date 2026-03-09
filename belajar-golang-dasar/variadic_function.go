package main

import "fmt"


func sumAll(numbers ...int)int{
	total := 0

	for _, number := range numbers{
		total += number
	}
	return total
}

func main(){
	fmt.Println(sumAll(10,10,101,10,10))


	numbers := []int{10,10,10,10}
	//untuk konversi slash ke variadic
	fmt.Println(sumAll(numbers...))
}
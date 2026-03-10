package main

import "fmt"

func main(){
	counter := 0

	increment := func() {
		fmt.Println("Increment")
		counter++
	}

	increment()
	increment() 
	//jadi si clouser bukan 0 tapi jdi nambah sesuai jumlah fungi increment yg di panggil
	fmt.Println(counter)
}
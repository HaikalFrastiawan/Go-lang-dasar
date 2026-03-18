package main

import (
	"container/list"
	"fmt"
)

func main() {
	var data  = list.New()

	data.PushBack("Haikal")
	data.PushBack("Frastiawan")

	var head = data.Front()
	fmt.Println(head.Value) // haikal

	next := head.Next() //Frastiawan
	fmt.Println(next.Value)


	for e := data.Front(); e != nil; e = e.Next() {	
		fmt.Println(e.Value)
	}
}
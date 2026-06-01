package golang

import "testing"

type Bag[T any] []T

func PrintBag[T any](bag Bag[T]) {
	for _, item := range bag {
		println(item)
	}
}

func TestBagString(t *testing.T){
	names := Bag[string]{"Alice", "Bob", "Charlie"}
	PrintBag(names)
}

func TestBagInt(t *testing.T){
	numbers := Bag[int]{1, 2, 3, 4, 5}
	PrintBag(numbers)
}


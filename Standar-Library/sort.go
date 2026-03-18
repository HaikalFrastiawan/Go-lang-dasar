package main

import (
    "fmt"
    "sort"
)

func main() {
    //int
    angka := []int{5, 2, 9, 1, 4}
    sort.Ints(angka)
    fmt.Println("Angka Terurut:", angka) 

    // String
    nama := []string{"Eko", "Budi", "Joko", "Andi"}
    sort.Strings(nama)
    fmt.Println("Nama Terurut:", nama) 
}
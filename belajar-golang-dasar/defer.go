package main

import "fmt"

func looging(){
	fmt.Println("selesai memanggil function")

}

func RunAplication(){
	//menunda ekseusi program hingga semua program yg di jalankan selesai di eksekusi
	defer looging()
	fmt.Println("Run Aplication")
}

func main(){	
	RunAplication()
}
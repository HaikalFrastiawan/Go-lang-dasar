package main

import "fmt"

func tanganiBenca(){
	if r := recover(); r != nil {
		fmt.Println("Aplikasi Selamat dari error", r)
	}

}

func main(){
	defer tanganiBenca()

	fmt.Println("menjalankan sistem...")	
	panic("koneksi terputus") //triger panic

	fmt.Println("code tidak berjalan")
}
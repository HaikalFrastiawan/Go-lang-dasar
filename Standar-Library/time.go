package main

import (
	"fmt"
	"time"
)

func main() {
	
	now := time.Now()
	bukaPuasa := time.Date(now.Year(), now.Month(), now.Day(), 18, 14, 0, 0, now.Location())

	selisih := bukaPuasa.Sub(now)

	fmt.Println("Waktu Sekarang:", now.Format("15:04:05"))
	fmt.Println("Waktu Buka   :", bukaPuasa.Format("15:04:05"))

	if selisih.Seconds() <= 0 {
		fmt.Println("Selamat berbuka puasa!")
	} else {

		fmt.Printf("Sisa waktu: %.0f jam, %.0f menit, %.0f detik lagi\n", 
			selisih.Hours(), selisih.Minutes(), selisih.Seconds())
		

		fmt.Println("Durasi persis:", selisih.String())
	}
}
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Haikal Fraastiawan","Haikal")) //Mengecek apakah sebuah kata ada di dalam teks.

	fmt.Println(strings.Split("Haikal Fraastiawan", "Haikal")) //Memotong string berdasarkan pemisah tertentu.

	fmt.Println(strings.ToLower("Haikal Fraastiawan")) //Mengubah semua huruf menjadi kecil.

	fmt.Println(strings.ToUpper("Haikal Fraastiawan"))//Mengubah semua huruf menjadi kapital.

	fmt.Println(strings.Trim("Haikal Fraastiawan", "Haikal")) //Menghapus karakter tertentu di awal dan akhir string.

	fmt.Println(strings.ReplaceAll("Haikal Fraastiawan", "Haikal", "fras")) //Mengganti semua kata tertentu dengan kata baru.
}
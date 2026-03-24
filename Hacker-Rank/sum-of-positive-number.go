package main

import (
    "fmt"
)

/*
 * Lengkapi fungsi di bawah ini.
 */
func sumPositive(arr []int32) {

	
    // 1. Deklarasikan variabel untuk menampung total (Gunakan int64 agar muat banyak)
	var total int64 = 0
    
    // 2. Buat perulangan (loop) untuk mengecek setiap angka di dalam arr
	for i := 0; i < len(arr); i++ {
		if arr[i] > 0 {
			total += int64(arr[i])
		}
		
	}
	fmt.Println(total)
    
    // 3. Di dalam loop, buat kondisi (if) jika angka > 0
    
    // 4. Tambahkan angka tersebut ke variabel total
    
    // 5. Cetak total menggunakan fmt.Println()
}

func main() {
    // Ini untuk mengetes kodemu
    testCase := []int32{1, -4, 7, 12}
    sumPositive(testCase) // Seharusnya muncul angka 20
}
package main

import "fmt"

func main() {
	// 1. CLOSURE: Saldo disimpan di dalam closure agar aman (Encapsulation)
	saldo := 1000000
	transfer := func(jumlah int) {
		
		// 2. DEFER & RECOVER: Jaring pengaman
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("⚠️ SISTEM TERHENTI: %v. Transaksi dibatalkan.\n", r)
			}
			fmt.Println("📢 Laporan: Sesi transaksi selesai.")
		}()

		fmt.Printf("Mencoba transfer: Rp%d...\n", jumlah)

		// 3. PANIC: Simulasi kondisi fatal (Misal: Deteksi Fraud)
		if jumlah > 5000000 {
			panic("DETEKSI FRAUD! Nominal terlalu besar")
		}

		if jumlah > saldo {
			fmt.Println("❌ Gagal: Saldo tidak cukup.")
			return
		}

		saldo -= jumlah
		fmt.Printf("✅ Berhasil! Sisa saldo: Rp%d\n", saldo)
	}

	// --- UJI COBA ---
	
	// Transaksi Normal
	transfer(200000) 
	fmt.Println("---------------------------")

	// Transaksi yang memicu PANIC (Percobaan Hack/Nominal Gila)
	transfer(9999999) 
	fmt.Println("---------------------------")

	// Program masih tetap jalan karena ada RECOVER!
	transfer(100000) 
}
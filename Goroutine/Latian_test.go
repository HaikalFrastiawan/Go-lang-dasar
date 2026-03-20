package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestSimulasiKopi(t *testing.T) {
	// 1. Meja Pick-up (Channel)
	// Kita pakai buffer 1 supaya Barista bisa naruh kopi meski pelanggan belum datang
	mejaPickup := make(chan string, 1)

	fmt.Println("Pelanggan: Saya pesan 1 Kopi Susu ya!")

	// 2. Barista mulai bekerja di Background (Goroutine)
	go func() {
		fmt.Println("Barista: Oke, pesanan diterima. Sedang dibuat...")
		
		// Simulasi meracik kopi (butuh waktu 2 detik)
		time.Sleep(2 * time.Second)
		
		hasil := "Kopi Susu Gula Aren"
		
		// 3. Taruh di meja pickup (Kirim ke Channel)
		fmt.Println("Barista: Kopi sudah jadi! Saya taruh di meja ya.")
		mejaPickup <- hasil
	}()

	// 4. Pelanggan melakukan hal lain dulu sambil nunggu
	fmt.Println("Pelanggan: Sambil nunggu, saya main HP dulu...")
	
	// 5. Mengambil kopi dari meja (Menerima dari Channel)
	// Baris ini akan MENUNGGU sampai Barista selesai mengirim data
	kopi := <-mejaPickup

	fmt.Printf("Pelanggan: Akhirnya! Saya terima %s. Mantap!\n", kopi)
}
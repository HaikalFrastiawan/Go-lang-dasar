package goroutine

import (
	"fmt"
	"strconv"
	"testing"
	"time"

)

//Buffer Channel
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

//Range channel Latihan 1
func TestLatihanRange(t *testing.T){
	channel := make(chan string)

	go func(){
		for i := 0; i < 10; i ++ {
			channel <- "Nomor telepon ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Mengirim SMS Ke: ",data)
	}
	fmt.Println("sekesai")
}

//Range Channel + Buffered Latihan 2
func TestGudangLogistik(t *testing.T) {
	// 1. Inisialisasi Channel dengan BUFFER (Kapasitas Gudang)
	// Silakan ganti angka 3 ini untuk melihat perbedaannya
	gudang := make(chan string, 3) 

	// 2. PABRIK (Producer Goroutine) - Kerjanya Sangat Cepat
	go func() {
		for i := 1; i <= 10; i++ {
			barang := "Barang-" + strconv.Itoa(i)
			
			// Pabrik langsung kirim barang ke gudang tanpa nunggu
			gudang <- barang 
			fmt.Println("🏭 PABRIK: Berhasil kirim ke gudang:", barang)
		}
		// Tutup gudang kalau produksi hari ini selesai
		close(gudang)
	}()

	// 3. TRUK (Consumer - Main Program) - Kerjanya Lambat
	fmt.Println("🚛 TRUK: Mulai mengangkut barang dari gudang...")
	
	for barang := range gudang {
		// Simulasi bongkar muat 1 detik per barang
		time.Sleep(1 * time.Second) 
		
		fmt.Printf("🚛 TRUK: Selesai angkut %s (Gudang sekarang ada ruang kosong)\n", barang)
	}

	fmt.Println("✅ SEMUA LOGISTIK SELESAI")
}
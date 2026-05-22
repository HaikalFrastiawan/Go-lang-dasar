package golang

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// Membuat custom type untuk keperluan testing
type Umur int
type Suhu float64

func TestDouble_TypeApproximation(t *testing.T) {
	// Skenario 1: Menggunakan tipe data bawaan (int murni)
	t.Run("Tipe Bawaan: int", func(t *testing.T) {
		var input int = 15
		expected := 30
		result := Double(input)
		
		assert.Equal(t, expected, result)
	})

	// Skenario 2: Menggunakan tipe data bawaan (float64 murni)
	t.Run("Tipe Bawaan: float64", func(t *testing.T) {
		var input float64 = 5.5
		expected := 11.0
		result := Double(input)
		
		assert.Equal(t, expected, result)
	})

	// Skenario 3: Menggunakan Custom Type (Umur -> ~int)
	t.Run("Custom Type: Umur berbasis int", func(t *testing.T) {
		var input Umur = 25
		expected := Umur(50) // Hasilnya juga harus berupa custom type Umur
		result := Double(input)
		
		assert.Equal(t, expected, result)
		// Memastikan tipe datanya tidak berubah menjadi int biasa
		assert.IsType(t, Umur(0), result, "Tipe kembalian harus tetap 'Umur'")
	})

	// Skenario 4: Menggunakan Custom Type (Suhu -> ~float64)
	t.Run("Custom Type: Suhu berbasis float64", func(t *testing.T) {
		var input Suhu = 36.5
		expected := Suhu(73.0)
		result := Double(input)
		
		assert.Equal(t, expected, result)
		assert.IsType(t, Suhu(0.0), result, "Tipe kembalian harus tetap 'Suhu'")
	})
}
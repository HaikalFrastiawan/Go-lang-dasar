package helper

import (
	"fmt"
	"os"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// benchmark = kecepatan func
// runing benchmark :
// go test -v -bench=. (runing seluruh)
// go test -v -run=NotMatchUnitTest -bench=BenchMarkTest (runing banchmark tanpa runing unit test)
func BenchmarkHelloWorldHaikal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Haikal")
	}
}

func BenchmarkHelloWorldFrastiawan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Frastiawan")
	}
}

// Table test
func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "HaikalTable",
			request:  "HaikalTable",
			expected: "Hello HaikalTable",
		},
		{
			name:     "FrastiawanTable",
			request:  "FrastiawanTable",
			expected: "Hello FrastiawanTable",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}

}

// SubTest
func TestSubtest(t *testing.T) {

	//jika ingin menjalankan salah satu testnya saja
	//go test TestSubtest/Haikal
	//go test TestSubtest/Frastiawan

	t.Run("Haikal", func(t *testing.T) {
		result := HelloWorld("Haikal")
		require.Equal(t, "Hello Haikal", result, "Result must be 'Hello Haikal'")
	})

	t.Run("Frastiawan", func(t *testing.T) {
		result := HelloWorld("Frastiawan")
		require.Equal(t, "Hello Frastiawan", result, "Result must be 'Hello Frastiawan'")
	})

}

// TestMain -> before after tesst
func TestMain(m *testing.M) {
	// --- BEFORE TESTING ---
	fmt.Println("Before Unit Test")
	fmt.Println("Buka koneksi database...")
	fmt.Println("Nyalakan server Redis...")

	// Jalankan semua unit test (fungsi Test... lainnya)
	exitCode := m.Run()

	// --- AFTER TESTING ---
	fmt.Println(" After Unit Test")
	fmt.Println("Hapus data sampah di database...")
	fmt.Println("Tutup koneksi database...")

	// Keluar dari proses test dengan kode yang sesuai
	os.Exit(exitCode)
}

// skip
func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("can not run on Windows")
	}

	result := HelloWorld("Haikal")
	require.Equal(t, "Hello Haikal", result, "Result must be 'Hello Haikal'")

}

// Require -> sama seperti failNow() ketika gagal kode bawahnya ga di eksekusi
func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Haikal")
	require.Equal(t, "Hello Haikal", result, "Result must be 'Hello Haikal'")
	//Require -> sama seperti failNow() ketika gagal kode bawahnya ga di eksekusi
	fmt.Println("Test HelloWorld with Rquire DOne")
}

// assertion
func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Haikal")
	assert.Equal(t, "Hello Haikal", result, "Result must be 'Hello Haikal'")
	fmt.Println("Test HelloWorld with Assert DOne")
	//kalo gagal lebih detail error nya
}

func TestHelloWorldHaikal(t *testing.T) {
	result := HelloWorld("Haikal")

	if result != "Hello Haikal" {
		t.Error("Result Must be 'Hello Haikal'")

	}

	fmt.Println("TestHelloWorld Done")
}

func TestHelloWorldFrastiawan(t *testing.T) {
	result := HelloWorld("Frastiawan")

	if result != "Hello Frastiawan" {
		t.Fatal("Result Must be 'Hello Frastiawan'")
	}
	fmt.Println("TestHelloWorldFrastiawan Done")
}

//runing test codenya
//go test : untuk menjalakan unit test
//go test -v : keseluruhan code
//go test -v -run (nama functionnya -> TestHelloWorld): ini hanya  1 fungsi

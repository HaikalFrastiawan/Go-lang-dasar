package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello world")
}

func TestCreateGoroutine(t *testing .T) {
	go RunHelloWorld()
	fmt.Println("Upss")

	time.Sleep(1 *time.Second)
}
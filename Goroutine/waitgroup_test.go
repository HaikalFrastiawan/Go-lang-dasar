package goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsycnchronous(group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)

	fmt.Println("Hello")
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsycnchronous(group)
	}

	group.Wait()
	fmt.Println("selesai")

}

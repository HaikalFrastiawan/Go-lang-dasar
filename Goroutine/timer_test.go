package goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)

	fmt.Println("Mulai pada: ", time.Now())

	time := <-timer.C
	fmt.Println("Selesai Pada: ", time)
}

func TestAfterTimer(t *testing.T) {
	channel := time.After(5 * time.Second)
	fmt.Println("Mulai pada: ", time.Now())

	time := <-channel
	fmt.Println("Selesai Pada: ", time)
}

func TestAfterFunction(t*testing.T){
	group := sync.WaitGroup{}
	group.Add(1)

	time.AfterFunc(5 * time.Second, func ()  {
		fmt.Println("Waktu selesai AfterFunc",time.Now())
		group.Done()
	})
	fmt.Println("Menunggu AfterFunc selesai...")
	fmt.Println(time.Now())
	group.Wait()
}

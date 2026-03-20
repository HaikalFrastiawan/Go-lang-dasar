package goroutine

import (
	"fmt"
	"testing"
	"time"

)

func TestCreateChannel(t *testing.T){
	channel := make(chan string)
	defer close(channel)

	go func () {
		time.Sleep(2 * time.Second)

		//mengirim data
		channel <- "Haikal Frastiawan"
		fmt.Println("Selesai mengirim data  ke channel")
	}()

	//menerima data
	data := <- channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

//channel as parameter
func  GiveMeResponse(channel chan string){
	time.Sleep(2 * time.Second)
	channel <- "Haikal Frastiawan"
}

func TestChannelAsParameter(t *testing.T){
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <- channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

//channel in and out
func OnlyIn (channel chan<- string){
	time.Sleep(2 * time.Second)
	channel <- "Haikal Frastiawan"
}

func OnlyOut (channel <-chan string){
	data := <- channel
	fmt.Println(data)
}

func TestInAndOutChannel(t *testing.T){
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

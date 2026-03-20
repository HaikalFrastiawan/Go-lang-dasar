package goroutine

import (
	"fmt"
	"testing"
)

func TestCreateChannel(t *testing.T){
	channel := make(chan string)

	//mengirim data channel
	channel <- "Haikal"


	//menerima data channel
	data := <- channel
	fmt.Println(<- channel)


	//untuk mastiin udah di close channel 
	defer close(channel)
	//jangan lupa di close
	close(channel)
}
package goroutine

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	var x int64 = 0

	group := sync.WaitGroup{}
	//total 100.000
	for i := 1; i <= 1000; i++ {
		go func() {
			group.Add(1)
			for j := 1; j <= 100; j++ {
				//untuk data primity / number race condition
				atomic.AddInt64(&x, 1)
			}
			group.Done()
		}()
	}
	group.Wait()
	fmt.Println("counter = ", x)
}

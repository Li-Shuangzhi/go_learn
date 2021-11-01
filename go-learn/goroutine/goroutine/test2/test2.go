package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counter int64
	wg      sync.WaitGroup
)

func main() {
	wg.Add(2)
	go incCounter(1)
	go incCounter(2)
	wg.Wait()
	fmt.Println("final counter:", counter)
}

func incCounter(id int) {
	defer wg.Done()

	for i := 0; i < 2; i++ {
		atomic.AddInt64(&counter, 1)

		//当前goroutine从线程退出, 并放回队列
		runtime.Gosched()

	}
}

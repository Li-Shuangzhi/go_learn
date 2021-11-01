package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	//一定要先make
	numChan := make(chan int, 10)
	//strChan := make(chan string)

	for i := 0; i < 10; i++ {
		for i := 0; i < 10; i++ {
			numChan <- i
			fmt.Println("写入", i)
		}
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			data := <-numChan
			fmt.Println("读到", data)
		}
	}()

	wg.Wait()

}

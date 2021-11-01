package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//var memoryAccess sync.Mutex
	//data := 0
	//go func() {
	//	memoryAccess.Lock()
	//	data++
	//	memoryAccess.Unlock()
	//}()
	//memoryAccess.Lock()
	//if data == 0 {
	//	fmt.Println("the value is 0")
	//} else {
	//	fmt.Printf("the value is %v", data)
	//}
	//memoryAccess.Unlock()
	test1()
}

func test() {
	channel := make(chan int)
	data := 0
	go func() {
		time.Sleep(time.Second)
		data++
		<-channel
		fmt.Println("func end")
	}()
	channel <- 1
	fmt.Println("end")
	//if data == 0 {
	//	fmt.Println("the value is 0")
	//} else {
	//	fmt.Printf("the value is %v", data)
	//}
}

func demo() {
	channel := make(chan int)
	go func() {
		time.Sleep(time.Second)
		<-channel
		fmt.Println("func end")
	}()
	channel <- 1
	fmt.Println("end")
}

type value struct {
	mu    sync.Mutex
	value int
}

var wg sync.WaitGroup

func test1() {
	printSum := func(v1, v2 *value) {
		defer wg.Done()
		v1.mu.Lock()
		defer v1.mu.Unlock()
		v2.mu.Lock()
		defer v2.mu.Unlock()
		fmt.Printf("sum=%v\n", v1.value+v2.value)
	}
	var a, b value
	wg.Add(2)
	go printSum(&a, &b)
	go printSum(&b, &a)
	wg.Wait()
}

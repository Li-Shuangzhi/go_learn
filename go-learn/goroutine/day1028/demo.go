package main

import (
	"fmt"
	"runtime"
	"sync"
)

func hello() {
	defer wg.Done()
	fmt.Println("hello")
}

type Stu struct {
	name string
}

var wg sync.WaitGroup

func main() {
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.GOMAXPROCS(0))
	wg.Add(1)
	go hello()
	wg.Wait()
	fmt.Println("adapter")
}

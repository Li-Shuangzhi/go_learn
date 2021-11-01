package main

import (
	"fmt"
)

func main() {
	//	c := make(chan struct{})
	//	go func(i chan struct{}) {
	//		sum := 0
	//		for i := 0; i < 10000; i++ {
	//			sum += i
	//		}
	//		println(sum)
	//		c <- struct{}{}
	//	}(c)
	//	println("Num=", runtime.NumCgoCall())
	//	<- c

	//c := make(chan struct{})
	//ci := make(chan int, 100)
	//go func() {
	//	close(ci)
	//	for i := 0; i < 10; i++ {
	//		ci <- i
	//	}
	//	c <- struct{}{}
	//}()
	//
	//fmt.Println(runtime.NumGoroutine())
	//
	//<-c
	//
	//fmt.Println(runtime.NumGoroutine())
	//
	//for v := range ci {
	//	println(v)
	//}
	test()
}

func test() {
	fmt.Println("++++++++++++++++++++")
	c := make(chan int, 5)
	c <- 1
	c <- 3
	c <- 5
	c <- 1
	c <- 9
	close(c)

	for i := range c {
		<-c
		fmt.Println(i)
		fmt.Println("len:", len(c))
	}
}

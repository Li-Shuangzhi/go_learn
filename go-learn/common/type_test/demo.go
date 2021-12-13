package main

import "fmt"

//自定义类型有两种 : 1. 结构体 2. 类型别名
// go语言中的引用类型 : slice map channel interface func

type LSZ []int

func (l LSZ) notify() {
	fmt.Println(l)
}

func main() {
	var l LSZ
	l = append(l, 10)
	l.notify()
}

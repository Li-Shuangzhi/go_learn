package main

import (
	"fmt"
	"runtime"
	"strconv"
)

func main() {
	f1()
}

func f1() {
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		fmt.Println("runtime.Caller() failed")
	}
	fmt.Println("name:" + runtime.FuncForPC(pc).Name())
	fmt.Println("file:" + file)
	fmt.Println("line", strconv.Itoa(line))
}

package main

import (
	"fmt"
	"testing"
)

func TestSingleton(t *testing.T) {
	for i := 0; i < 10; i++ {
		go getInstance()
	}
	fmt.Scanln()
}

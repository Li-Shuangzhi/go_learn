package _struct

import "fmt"

type Student struct {
	Id   int
	Name string
}

func demo() {
	s1 := &Student{
		Id:   100,
		Name: "lsz",
	}
	s2 := &Student{
		Id:   100,
		Name: "lsz",
	}
	fmt.Println(*s1 == *s2)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

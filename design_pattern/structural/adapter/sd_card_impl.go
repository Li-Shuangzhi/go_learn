package main

import "fmt"

type SDCardImpl struct {
	info string
}

func (s *SDCardImpl) readSD() {
	fmt.Printf("从sdcard读出: %v\n", s.info)
}

func (s *SDCardImpl) writeSD(str string) {
	fmt.Println("写入sdcard成功")
	s.info += str
}

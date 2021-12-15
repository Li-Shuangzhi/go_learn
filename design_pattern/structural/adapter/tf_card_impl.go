package main

import "fmt"

type TFCardImpl struct {
	info string
}

func (t *TFCardImpl) readTF() {
	fmt.Printf("从tfcard读出: %v\n", t.info)
}

func (t *TFCardImpl) writeTF(str string) {
	t.info += str
	fmt.Println("写入tfcard成功")
}

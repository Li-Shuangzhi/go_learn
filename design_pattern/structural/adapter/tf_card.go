package main

type TFCard interface {
	writeTF(str string)
	readTF()
}

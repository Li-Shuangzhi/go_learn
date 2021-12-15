package main

type SDCard interface {
	readSD()
	writeSD(str string)
}

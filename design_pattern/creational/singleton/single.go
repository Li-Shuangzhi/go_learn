package main

/*
Singleton is a creational design pattern,
which ensures that only one object of its kind exists
and provides a single point of access to it for any other code.
*/
import (
	"fmt"
	"sync"
)

type single struct{}

var singleton *single
var lock = &sync.Mutex{}

func getInstance() *single {
	if singleton == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleton == nil {
			singleton = &single{}
			fmt.Println("single create now")
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}
	return singleton
}

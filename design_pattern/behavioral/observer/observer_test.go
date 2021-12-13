package observer

import (
	"fmt"
	"testing"
)

func TestObserver(t *testing.T) {
	item := &Item{value: 10}
	observer1 := &ConcreteObserver{id: 1001, high: 1}
	item.register(observer1)
	item.value++
	item.notify()
	fmt.Println(observer1.high)
}

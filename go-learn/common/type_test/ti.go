package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex
var wg sync.WaitGroup

type Ban struct {
	visitIPs map[string]time.Time
}

func NewBan() *Ban {
	return &Ban{visitIPs: make(map[string]time.Time)}
}
func (o *Ban) visit(ip string) bool {
	defer mutex.Unlock()
	mutex.Lock()
	{
		if _, ok := o.visitIPs[ip]; ok {
			return true
		}
		o.visitIPs[ip] = time.Now()
		return false
	}

}
func main() {
	success := 0
	ban := NewBan()
	for i := 0; i < 1000; i++ {
		for j := 0; j < 100; j++ {
			wg.Add(1)
			t := j
			go func() {
				defer wg.Done()
				ip := fmt.Sprintf("192.168.1.%d", t)
				if !ban.visit(ip) {
					success++
				}
			}()
		}
	}
	wg.Wait()
	{

	}
	fmt.Println("success:", success)
}

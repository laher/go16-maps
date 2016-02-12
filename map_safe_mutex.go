package main

import (
	"fmt"
	"sync"
	"time"
)

var l = sync.RWMutex{}
var m = map[string]time.Time{}

func main() {

	go func() {
		for {
			time.Sleep(1 * time.Second)
			l.Lock()
			m["a"] = time.Now()
			l.Unlock()
		}
	}()

	for {
		time.Sleep(time.Second / 2)
		l.RLock()
		fmt.Printf("time 'a' is %v\n", m["a"])
		l.RUnlock()
	}
}

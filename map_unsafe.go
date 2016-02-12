package main

import (
	"fmt"
	"time"
)

var m = map[string]time.Time{}

func main() {

	go func() {
		for {
			time.Sleep(1 * time.Second)
			m["a"] = time.Now()
		}
	}()

	for {
		time.Sleep(time.Second / 2)
		fmt.Printf("time 'a' is %v\n", m["a"])
	}
}

package main

import (
	"fmt"
	"time"
)

var m = map[string]time.Time{}

func main() {

	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered from panic", r)
			}
		}()
		for {
			time.Sleep(10 * time.Millisecond)
			m["a"] = time.Now()
		}
	}()

	for {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered from panic", r)
			}
		}()
		time.Sleep(time.Millisecond)
		_ = m["a"]
	}
}

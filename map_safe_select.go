package main

import (
	"fmt"
	"time"
)

var m = map[string]time.Time{}
var r = make(chan time.Time)
var w = make(chan time.Time)

func main() {

	go func() {
		for {
			select {
			case t := <-w:
				m["a"] = t
			case r <- m["a"]:

			}
		}
	}()
	go func() {
		for {
			time.Sleep(1 * time.Second)
			w <- time.Now()
		}
	}()

	for {
		time.Sleep(time.Second / 2)
		fmt.Printf("time 'a' is %v\n", <-r)
	}
}

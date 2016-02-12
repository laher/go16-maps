package main

import (
	"fmt"
	"time"
)

var m = map[string]time.Time{}

type writer struct {
	k string
	v time.Time
}

var w = make(chan writer)

func write(k string, v time.Time) {
	w <- writer{k: "a", v: v}
}

type reader struct {
	k  string
	vc chan time.Time
}

func read(k string) time.Time {
	rdr := reader{
		k:  k,
		vc: make(chan time.Time),
	}
	r <- rdr
	return <-rdr.vc
}

var r = make(chan reader)

func main() {

	go func() {
		for {
			select {
			case kv := <-w:
				m[kv.k] = kv.v
			case kvc := <-r:
				kvc.vc <- m[kvc.k]
			}
		}
	}()
	go func() {
		for {
			time.Sleep(1 * time.Second)
			write("a", time.Now())
		}
	}()

	for {
		time.Sleep(time.Second / 2)
		fmt.Printf("time 'a' is %v\n", read("a"))
	}
}

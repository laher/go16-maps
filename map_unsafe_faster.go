package main

import "time"

var m = map[string]time.Time{}

func main() {

	go func() {
		for {
			time.Sleep(10 * time.Millisecond)
			m["a"] = time.Now()
		}
	}()

	for {
		time.Sleep(time.Millisecond)
		_ = m["a"]
	}
}

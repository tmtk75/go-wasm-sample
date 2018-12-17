// +build js,wasm

package main

import (
	"fmt"
	"time"
)

var Version string

func main() {
	c := make(chan struct{}, 0)

	f := func(t time.Duration) {
		go (func() {
			i := 0
			for {
				time.Sleep(t * time.Second)
				fmt.Printf("sleep in %v: %v\n", t*time.Second, i)
				i++
			}
		})()
	}

	f(1)
	f(2)

	<-c
}

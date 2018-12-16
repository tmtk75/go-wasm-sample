// +build js,wasm

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"syscall/js"
)

func request(i []js.Value) {
	c := http.Client{}
	r, err := c.Get("http://example.com")
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Printf("%v", string(b))
}

func registerCallbacks() func() {
	cb := js.NewCallback(request)
	js.Global().Set("request", cb)
	return func() {
		cb.Release()
	}
}

func main() {
	c := make(chan struct{}, 0)

	println("WASM Go Initialized")
	r := registerCallbacks()
	defer r()
	<-c
}

// +build js,wasm

package main

import (
	"fmt"
	"strconv"
	"syscall/js"
)

func add(i []js.Value) {
	fmt.Printf("%v\n", i)
	value1 := js.Global().Get("document").Call("getElementById", i[0].String()).Get("value").String()
	value2 := js.Global().Get("document").Call("getElementById", i[1].String()).Get("value").String()

	int1, _ := strconv.Atoi(value1)
	int2, _ := strconv.Atoi(value2)

	a := int1 + int2
	js.Global().Set("output", a)
	fmt.Println(a)

	js.Global().Get("document").Call("getElementById", i[2].String()).Set("value", a)
}

func registerCallbacks() func() {
	cb := js.NewCallback(add)
	js.Global().Set("add", cb)
	return func() {
		cb.Release()
	}
}

func main() {
	c := make(chan struct{}, 0)

	println("WASM Go Initialized")
	// register functions
	r := registerCallbacks()
	<-c
	r()
}

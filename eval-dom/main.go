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
}

func registerCallbacks() {
	js.Global().Set("add", js.NewCallback(add))
}

func main() {
	c := make(chan struct{}, 0)

	println("WASM Go Initialized")
	// register functions
	registerCallbacks()
	<-c
}

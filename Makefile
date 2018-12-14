.PHONY: build
build: init test.wasm

test.wasm: main.go
	GOOS=js GOARCH=wasm go build -o test.wasm main.go

.PHONY: run
run: server.go
	go run ./server.go

.PHONY: init
init: wasm_exec.html wasm_exec.js

wasm_exec.html:
	curl -sO https://raw.githubusercontent.com/golang/go/master/misc/wasm/wasm_exec.html

wasm_exec.js:
	curl -sO https://raw.githubusercontent.com/golang/go/master/misc/wasm/wasm_exec.js

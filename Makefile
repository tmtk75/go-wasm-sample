.PHONY: build
build: init test.wasm

test.wasm: main.go
	GOOS=js GOARCH=wasm go build -o test.wasm main.go

.PHONY: run
run: server.go
	go run ./server.go

.PHONY: init
init: wasm_exec.html wasm_exec.js go_js_wasm_exec

wasm_exec.html:
	curl -sO https://raw.githubusercontent.com/golang/go/master/misc/wasm/wasm_exec.html

wasm_exec.js:
	#curl -sO https://raw.githubusercontent.com/golang/go/master/misc/wasm/wasm_exec.js
	cp `go env GOROOT`/misc/wasm/wasm_exec.js .

go_js_wasm_exec:
	curl -sO https://raw.githubusercontent.com/golang/go/master/misc/wasm/go_js_wasm_exec
	chmod +x go_js_wasm_exec

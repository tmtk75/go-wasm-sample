.PHONY: build
build: init test.wasm

.PHONY: serve
serve: server.go
	go run ./server.go

.PHONY: init
init: wasm_exec.html wasm_exec.js go_js_wasm_exec

wsm_exec.js:
	#curl -sO https://raw.githubusercontent.com/golang/go/master/misc/wasm/wasm_exec.js
	cp `go env GOROOT`/misc/wasm/wasm_exec.js .

go_js_wasm_exec:
	curl -sO https://raw.githubusercontent.com/golang/go/master/misc/wasm/go_js_wasm_exec
	chmod +x go_js_wasm_exec

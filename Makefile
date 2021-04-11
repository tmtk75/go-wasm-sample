.PHONY: build
build: init
	(cd calc; make)
	(cd eval-dom; make)
	(cd goroutine; make)
	(cd hello; make)
	(cd http; make)

.PHONY: serve
serve: server.go
	go run ./server.go

.PHONY: init
init: wasm_exec.js go_js_wasm_exec

wasm_exec.js:
	#curl -sO https://raw.githubusercontent.com/golang/go/master/misc/wasm/wasm_exec.js
	cp `go env GOROOT`/misc/wasm/wasm_exec.js .

go_js_wasm_exec:
	#curl -sO https://raw.githubusercontent.com/golang/go/master/misc/wasm/go_js_wasm_exec
	cp `go env GOROOT`/misc/wasm/go_js_wasm_exec .
	chmod +x go_js_wasm_exec

clean:
	rm -f wasm_exec.js go_js_wasm_exec


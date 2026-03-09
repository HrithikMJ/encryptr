go-wasm:
	cd wasm \
		&& GOOS=js GOARCH=wasm go build -a -o main.wasm ./wasm/encryptr \
		&& mv main.wasm ../web/static/main.wasm \
		&& cp "$$(go env GOROOT)/lib/wasm/wasm_exec.js" ../web/static/wasm_exec.js
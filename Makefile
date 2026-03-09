wasm:
	GOOS=js GOARCH=wasm go build -o wasm/main.wasm wasm/cmd/encryptr/main.go
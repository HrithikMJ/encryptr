go-wasm:
	cd wasm && GOOS=js GOARCH=wasm go build -a -o main.wasm ./cmd/encryptr && cp ./main.wasm ../web/static/main.wasm && cd ..
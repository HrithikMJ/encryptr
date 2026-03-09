package main

import (
	"crypto/sha256"
	"encoding/base64"
	"encrptr/internal/helpers"
	"syscall/js"
)

func main() {

	jsGlobal := js.Global()
	jsGlobal.Set("wasmEncodeImage", js.FuncOf(wasmEncodeImage))
	select {}
}
func hash32Bytes(s string) []byte {
	h := sha256.Sum256([]byte(s))
	return h[:]
}

func wasmEncodeImage(this js.Value, args []js.Value) any {
	promiseConstructor := js.Global().Get("Promise")

	return promiseConstructor.New(js.FuncOf(func(this js.Value, pArgs []js.Value) any {
		resolve := pArgs[0]
		reject := pArgs[1]

		if len(args) != 2 {
			reject.Invoke("invalid arguments")
			return nil
		}

		message := args[0].String()
		key := args[1].String()

		go func() {
			encrypted, err := helpers.Encrypt([]byte(message), hash32Bytes(key))
			if err != nil {
				reject.Invoke(err.Error())
				return
			}

			rawImage, err := helpers.EncodeImage(encrypted)
			if err != nil {
				reject.Invoke(err.Error())
				return
			}

			resolve.Invoke(base64.StdEncoding.EncodeToString(rawImage))
		}()

		return nil
	}))
}

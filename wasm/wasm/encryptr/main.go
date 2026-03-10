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
	jsGlobal.Set("wasmDecodeImage", js.FuncOf(wasmDecodeImage))
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

func wasmDecodeImage(this js.Value, args []js.Value) any {
	promiseConstructor := js.Global().Get("Promise")

	return promiseConstructor.New(js.FuncOf(func(this js.Value, pArgs []js.Value) any {
		resolve := pArgs[0]
		reject := pArgs[1]

		if len(args) != 2 {
			reject.Invoke("invalid arguments: need base64 image and key")
			return nil
		}

		b64Image := args[0].String()
		key := args[1].String()

		go func() {
			rawImage, err := base64.StdEncoding.DecodeString(b64Image)
			if err != nil {
				reject.Invoke("failed to decode base64: " + err.Error())
				return
			}

			data, err := helpers.DecodeImage(rawImage)
			if err != nil {
				reject.Invoke("failed to decode image: " + err.Error())
				return
			}

			hashedKey := hash32Bytes(key)
			plaintext, err := helpers.Decrypt(data, hashedKey)
			if err != nil {
				reject.Invoke("decryption failed: invalid key or corrupted image: " + err.Error())
				return
			}
			resolve.Invoke(string(plaintext))
		}()

		return nil
	}))
}

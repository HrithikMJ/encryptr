package main

import (
	"encrptr/internal/helpers"
	"hash/fnv"
	"syscall/js"
)

func main() {

	jsGlobal := js.Global()
	jsGlobal.Set("wasmEncodeImage", js.FuncOf(wasmEncodeImage))
	select {}
}
func hash32(s string) []byte {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum(nil)
}

func wasmEncodeImage(this js.Value, args []js.Value) any {
	if len(args) != 2 {
		return js.ValueOf("Invalid arguments")
	}
	message := args[0].String()
	key := args[1].String()
	encrypted, err := helpers.Encrypt([]byte(message), hash32(key))
	if err != nil {
		return js.ValueOf(err.Error())
	}
	return helpers.EncodeImage(encrypted)
}

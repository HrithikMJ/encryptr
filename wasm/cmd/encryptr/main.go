package main

import (
	"crypto/sha256"
	"encrptr/internal/helpers"
	"fmt"
)

func main() {
	message := "grow flee opera uncle track book mule six cattle bachelor wait liberty mind second fall churn badge hello erase uncover seven enrich flock favorite"
	key := "grow flee opera uncle track book mule six cattle bachelor wait liberty mind second fall churn badge hello erase uncover seven enrich flock favorite"

	hashedKey := hash32Bytes(key)
	fmt.Println("Hashed key:", hashedKey)
	// Encrypt the message
	encrypted, err := helpers.Encrypt([]byte(message), hashedKey)
	if err != nil {
		fmt.Println("encryption error:", err)
		return
	}
	fmt.Println("Encrypted:", len(encrypted))

	// Encode encrypted bytes into an image (PNG bytes)
	rawImage, err := helpers.EncodeImage(encrypted)
	if err != nil {
		fmt.Println("encode image error:", err)
		return
	}
	fmt.Println("Encoded PNG bytes:", len(rawImage))

	// Decode the image back into encrypted bytes
	text, err := helpers.DecodeImage(rawImage)
	if err != nil {
		fmt.Println("decode image error:", err)
		return
	}

	fmt.Println("Text:", len(text))
	// Decrypt the bytes back to original message
	plaintext, err := helpers.Decrypt(text, hashedKey)
	if err != nil {
		fmt.Println("decryption error:", err)
		return
	}
	fmt.Println("Decrypted plaintext:", string(plaintext))
}
func hash32Bytes(s string) []byte {
	h := sha256.Sum256([]byte(s))
	return h[:]
}

package main

import (
	"crypto/sha256"
	"encrptr/internal/helpers"
	"fmt"
)

func main() {
	message := "grow flee opera uncle track book mule six cattle bachelor wait liberty mind second fall churn badge hello erase uncover seven enrich flock favorite"
	key := "grow flee opera uncle track book mule six cattle bachelor wait liberty mind second fall churn badge hello erase uncover seven enrich flock favorite"
	encrypted, err := helpers.Encrypt([]byte(message), hash32Bytes(key))
	if err != nil {
		fmt.Println(err)
	}
	rawImage, err := helpers.EncodeImage(encrypted)
	fmt.Println(rawImage)
	if err != nil {
		fmt.Println(err)
	}
}
func hash32Bytes(s string) []byte {
	h := sha256.Sum256([]byte(s))
	return h[:]
}

package main

import (
	"encoding/hex"
	"fmt"
)

func encryption(text_string []byte, key []byte) []byte {
	new_string := make([]byte, len(text_string))
	j := 0
	for i := 0; i < len(text_string); i++ {
		new_string[i] = text_string[i] ^ key[j]
		if j+1 == len(key) {
			j = 0
		} else {
			j++
		}
	}
	return new_string
}

func main() {
	plaintext := []byte("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal")
	key := []byte("ICE")

	fmt.Println(hex.EncodeToString(encryption(plaintext, key)))
}

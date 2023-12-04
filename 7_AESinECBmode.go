package main

import (
	"encoding/base64"
    "os"
    "log"
    "fmt"
    "crypto/aes"
)

func decryptAESECB(data, key []byte) []byte {
    cipher, _ := aes.NewCipher([]byte(key))
    text := make([]byte, len(data))
    blockSize := 16

    for b0, bn := 0, blockSize; b0 < len(data); b0, bn = b0+blockSize, bn+blockSize {
        cipher.Decrypt(text[b0:bn], data[b0:bn])
    }

    return text
}

func main(){

    key := "YELLOW SUBMARINE"

    file, err := os.Open("7.txt")
	if err != nil {
		log.Fatal(err)
	}

    allBytes := make([]byte, 4000)
    file.Read(allBytes)
	content, _ := base64.StdEncoding.DecodeString(string(allBytes))

    fmt.Printf("Key: %s\n", string(key))
    text := decryptAESECB(content, []byte(key))
    fmt.Printf("Message:\n%s\n", string(text))
}

package main

import (
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	const hex_string = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	chars := []byte("!\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~")
	hex_chars := hex.EncodeToString(chars)
	bytes_chars, err := hex.DecodeString(hex_chars)
	if err != nil {
		log.Fatal(err)
	}
	bytes_string, err := hex.DecodeString(hex_string)
	if err != nil {
		log.Fatal(err)
	}
	var countE = make([]int, len(bytes_chars))
	var countA = make([]int, len(bytes_chars))
	var new_string string
	var biggerCountE int
	var biggerCountA int
	for i := 0; i < len(bytes_chars); i++ {
		invalid := false
		new_bytes := make([]byte, len(bytes_string))
		for j := 0; j < len(bytes_string); j++ {
			new_bytes[j] = bytes_chars[i] ^ bytes_string[j]
			//new_string := hex.EncodeToString(new_bytes)
			//fmt.Printf("%s\n",new_string)
		}
		for k := 0; k < len(new_bytes); k++ {
			specialChars := "_*%$#^~|{}`()[]"
			for l := 0; l < len(specialChars); l++ {
				char := specialChars[l]
				if string(new_bytes[k]) == string(char) {
					invalid = true
				}
			}
			if string(new_bytes[k]) == "e" || string(new_bytes[k]) == "E" {
				countE[i]++
			}
			if string(new_bytes[k]) == "a" || string(new_bytes[k]) == "A" {
				countA[i]++
			}
		}
		if invalid == true {
			continue
		}
		if countE[i] > biggerCountE && countA[i] > biggerCountA {
			biggerCountE = countE[i]
			biggerCountA = countA[i]
			new_string = string(new_bytes)
		}
	}
	fmt.Println(new_string)
}

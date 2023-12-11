package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"log"
	"os"
)

func analyze_string(conv_string []byte) bool {
	symbols := "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz "
	special := false
	var counter int
	var letter_count int
	for i := 0; i < len(conv_string); i++ {
		for j := 0; j < len(symbols); j++ {
			if conv_string[i] == symbols[j] {
				special = true
			}
		}
	}
	for i := 0; i < len(conv_string); i++ {
		for j := 0; j < len(letters); j++ {
			if conv_string[i] == letters[j] {
				letter_count++
			}
		}
	}
	for i := 0; string(conv_string) != "\n" && i < len(conv_string); i++ {
		counter++
	}
	if special == true {
		return false
	} else if counter == (letter_count + 1) {
		return true
	} else {
		return false
	}
}

func main() {
	file, err := os.Open("4.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	hex_string := bufio.NewScanner(file)
	for hex_string.Scan() {
		chars := []byte("!\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~")
		hex_chars := hex.EncodeToString(chars)
		bytes_chars, err := hex.DecodeString(hex_chars)
		if err != nil {
			log.Fatal(err)
		}
		bytes_string, err := hex.DecodeString(hex_string.Text())
		if err != nil {
			log.Fatal(err)
		}
		for i := 0; i < len(bytes_chars); i++ {
			new_bytes := make([]byte, len(bytes_string))
			for j := 0; j < len(bytes_string); j++ {
				new_bytes[j] = bytes_chars[i] ^ bytes_string[j]
				//new_string := hex.EncodeToString(new_bytes)
				//fmt.Printf("%s\n",new_string)
			}
			if analyze_string(new_bytes) {
				fmt.Println(string(new_bytes))
				fmt.Println("XORed with: " + string(bytes_chars[i]))
			}
		}
	}
	err = hex_string.Err()
	if err != nil {
		log.Fatal(err)
	}
}

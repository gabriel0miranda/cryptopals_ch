package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"math"
)

func HammingDistance(a, b []byte) int {
	distance := 0
	for i := range a {
		r := a[i] ^ b[i]
		for r > 0 {
			if r&1 == 1 {
				distance++
			}
			r = r >> 1
		}
	}
	return distance
}

func dotProduct(a, b []float64) float64 {
	sum := 0.0
	for i := range a {
		sum += a[i] * b[i]
	}
	return sum
}

func magnitude(a []float64) float64 {
	return math.Sqrt(dotProduct(a, a))
}

func cosineSimilarity(a, b []float64) float64 {
	return dotProduct(a, b) / (magnitude(a) * magnitude(b))
}

func analyze_string(conv_string []byte) float64 {
	var english_scores = []float64{.0817, .0149, .0278, .0425, .1270, .0223, .0202, .0609, .0697, .0015, .0077, .0402, .0241, .0675, .0751, .0193, .0009, .0599, .0633, .0906, .0276, .0098, .0236, .0015, .0197, .0007}

	cnt := make([]int, 26)
	for _, ch := range conv_string {
		if 'A' <= ch && ch <= 'Z' {
			ch -= 32
		}
		if 'a' <= ch && ch <= 'z' {
			cnt[int(ch)-'a']++
		}
	}
	total := float64(len(conv_string))
	score := make([]float64, 26)
	for i, c := range cnt {
		score[i] = float64(c) / total
	}
	return cosineSimilarity(score, english_scores)
}

func breakXOR(byteBlock []byte) byte {
	var english_score float64
	var key byte
	var block_score float64
	for i := 0; i <= 255; i++ {
		new_bytes := make([]byte, len(byteBlock))
		for j := 0; j < len(byteBlock); j++ {
			new_bytes[j] = byte(i) ^ byteBlock[j]
			//new_string := hex.EncodeToString(new_bytes)
			//fmt.Printf("%s\n",new_string)
		}
		block_score = analyze_string(new_bytes)
		if block_score > english_score {
			english_score = block_score
			key = byte(i)
		}
	}
	return key
}

func breakXORRepeated(allBytes []byte, keyLen int) []byte {
	fullKey := make([]byte, keyLen)
	blockSize := len(allBytes) / keyLen
	for i := 0; i < keyLen; i++ {
		block := make([]byte, blockSize)
		for j := 0; j < blockSize; j++ {
			block[j] = allBytes[i+j*keyLen]
		}
		key := breakXOR(block)
		fullKey[i] = key
	}
	return fullKey
}

func findKeySize(allBytes []byte) int {
	keyLen := 0
	smallerHammDist := 1000.0
	blockSize := len(allBytes) / 40
	for i := 1; i < 40; i++ {
		hammDist := 0.0
		for j := 0; j < blockSize; j++ {
			a := j * i
			b := (j + 1) * i
			c := (j + 2) * i
			hammDist += float64(HammingDistance(allBytes[a:b], allBytes[b:c]) / i)
		}
		hammDist /= float64(blockSize)
		if hammDist < smallerHammDist {
			smallerHammDist = hammDist
			keyLen = i
		}
	}
	fmt.Printf("Probable key size: %d bits, Smaller Hamming Distance: %.2f\n", keyLen, smallerHammDist)
	return keyLen
}

func decodeXOR(text_string []byte, key []byte) []byte {
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
	file, err := ioutil.ReadFile("6.txt")
	if err != nil {
		log.Fatal(err)
	}

	allBytes, _ := base64.StdEncoding.DecodeString(string(file))
	keyLen := findKeySize(allBytes)
	key := breakXORRepeated(allBytes, keyLen)
	fmt.Printf("Key: %s\n", string(key))
	text := decodeXOR(allBytes, key)
	fmt.Printf("Message:\n%s\n", string(text))
}

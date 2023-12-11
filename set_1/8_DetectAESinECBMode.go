package main

import (
    "os"
    "log"
    "fmt"
    "encoding/hex"
    "bufio"
    "bytes"
)

func detectECB(file  []byte) int {
    var decodedFile  [][]byte
    scanner := bufio.NewScanner(bytes.NewReader(file))

    for scanner.Scan(){
        decodedBytes := make([]byte,hex.DecodedLen(len(scanner.Bytes())))
        _, err := hex.Decode(decodedBytes, scanner.Bytes())
        if err != nil {
            log.Fatal(err)
        }
        
        decodedFile = append(decodedFile, decodedBytes)
    }

    detection := detectLine(decodedFile)
    return detection
}

func detectLine(data [][]byte) int {
    for i, ln := range data {
        block := make([][]byte, 0)
        for j := 0; j<len(ln); j += 16 {
            blocks := ln[j:min(j+15, len(ln))]
            for _, c := range block {
                if bytes.Equal(c,blocks) {
                    return i+1
                }
            }
            block = append(block, blocks)
        }
    }
    return 0
}

func min(a, b int) int {
    if a <= b {
        return a
    }
    return b
}

func main(){
    file, err := os.ReadFile("8.txt")
	if err != nil {
		log.Fatal(err)
	}
    
    fmt.Printf("Line with ECB: %d\n", detectECB(file))
}

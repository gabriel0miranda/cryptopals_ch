package main

import (
    "fmt"
    )

//Defining constants for the block to be padded and the desired block lenght
const Block = "I AM NOT PADDED YET"
const BlockLenght = 8

//Padding function according to PKCS7 padding
func PKCS7Padding(block []byte) (paddedBlock []byte){
    var paddedBlockSize int
    var paddingValue int
    if len(block)%BlockLenght == 0 {
        paddedBlockSize = len(block)+BlockLenght
        paddingValue = BlockLenght
    } else {
        paddedBlockSize = len(block)+BlockLenght-(len(block)%BlockLenght)
        paddingValue = len(block)%BlockLenght
    }

    paddedBlock = make([]byte,paddedBlockSize)

    fmt.Printf("padding value: %s\n",rune(byte(paddingValue)))

    for cursor := 0; cursor < paddedBlockSize; cursor++ {
        if cursor < len(block) {
            paddedBlock[cursor] = block[cursor]
        } else {
            paddedBlock[cursor] = byte(paddingValue)
        }
    }

    return paddedBlock
}

//Main function
func main(){
    fmt.Printf("Block given: %s\n",string(Block))
    fmt.Printf("Block lenght given: %d\n",BlockLenght)
    fmt.Printf("Padded block: %s\n",string(PKCS7Padding([]byte(Block))))
}

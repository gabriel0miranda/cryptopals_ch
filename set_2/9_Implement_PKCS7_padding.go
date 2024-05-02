package main

import (
    "fmt"
    "log"
    )

//Defining constants for the block to be padded and the desired block lenght
const Block = "I AM NOT PADDED YET"
const BlockLenght = 8

//Padding function according to PKCS7 padding
func PKCS7Padding(block []byte){
    var paddedBlockSize int
    var paddingValue int
    if BlockLenght%8 == 0 {
        paddedBlockSize = len(Block)+BlockLenght
        paddingValue = BlockLenght
    } else {
        paddedBlockSize = len(Block)+BlockLenght-(len(Block)%BlockLenght)
        paddingValue = len(Block)%BlockLenght
    }

    paddedBlock := make([]byte,paddedBlockSize)

    for 
}

//Main function
func main(){
    fmt.Printf("Block given: %s\n",string(Block))
    fmt.Printf("Block lenght given: %d\n",BlockLenght)
    fmt.Printf("Padded block: %s\n",string(PKCS7Padding([]byte(Block))))
}

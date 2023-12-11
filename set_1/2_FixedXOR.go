package main

import "fmt"
import "encoding/hex"
import "log"

func main(){
    const string1 = "1c0111001f010100061a024b53535009181c"
    const string2 = "686974207468652062756c6c277320657965"

    hex_bytes1, err := hex.DecodeString(string1)
    if err != nil {
        log.Fatal(err)
    }
    hex_bytes2, err := hex.DecodeString(string2)
    if err != nil {
        log.Fatal(err)
    }
    
    xored_bytes := make([]byte,len(hex_bytes1))

    for i:=0; i<len(hex_bytes1); i++ {
        xored_bytes[i] = hex_bytes1[i] ^ hex_bytes2[i]
    }
    hex_string := hex.EncodeToString(xored_bytes)
    fmt.Println(hex_string)
}

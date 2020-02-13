package main

import "fmt"
import "encoding/hex"
import "log"
import "encoding/base64"

func main(){
    const str = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
    hex_bytes, err := hex.DecodeString(str)
    if err != nil {
        log.Fatal(err)
    }
    b64_string := base64.StdEncoding.EncodeToString(hex_bytes)
    fmt.Println(b64_string)
}


package main

import (
    "fmt"
)

func main() {

	text := "Hello World"

    encodedBytes :=  xorEncodeDecode([]byte(text), uint64(1343123213123434))
    fmt.Println(encodedBytes)
    decodedBytes := xorEncodeDecode(encodedBytes, uint64(1343123213123434))
    fmt.Println(string(decodedBytes[:]))
    // fmt.Println(secret_key)
    // decodedTest := xorEncodeDecode([]byte(text), uint64(secret_key))
    // fmt.Println(decodedTest)
    // fmt.Println(secret_key)

}

package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"net"
	"os"
	"project1/helper"
	"time"
)

func xorEncodeDecode(text []byte, key *uint64) []byte {
	encryptedLi := make([]byte, len(text))
	for i := 0; i < len(text); i++ {
		encryptedLi[i] = text[i] ^ byte(*key)
	}

	*key = helper.XorShift(*key)
	return encryptedLi
}

func sendDataToClient(key *uint64, conn net.Conn, bufferSize int) {
	msg := helper.GenerateRandomString(bufferSize)
	msgEncoded := xorEncodeDecode([]byte(msg), key)

	conn.Write(msgEncoded)

	// 8-byte acknowledgement
	returnDataFromServer := make([]byte, 8)
	conn.Read(returnDataFromServer)
	hash := sha256.Sum256([]byte("Data recieved"))
	if bytes.Equal(hash[:8], returnDataFromServer) {
		fmt.Println("Acknowledgment recieved: data are equal")
	} 

}

func main() {
	// Input argument for host address
	arguments := os.Args[1:]
	if len(arguments) > 1 {
		fmt.Println("Error: only one argument accepted which is the server address and port")
		return
	}

	conn, err := net.Dial("udp", "127.0.0.1:1234")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	defer conn.Close()
	bufferSizes := []int{8, 64, 256, 512}

	for _, value := range bufferSizes {
		key := uint64(1343123213123434)
		start := time.Now()
		sendDataToClient(&key, conn, value)
		elapsed := time.Since(start)
		fmt.Println("Elapsed time:", elapsed, "for the bufferSize of", value)
	}
	fmt.Println("--------------------------------------")
}

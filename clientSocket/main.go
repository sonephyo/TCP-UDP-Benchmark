package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

func xorShift(r uint64) uint64 {
	r ^= r << 13
	r ^= r >> 7
	r ^= r << 17
	return r
}

func xorEncodeDecode(text []byte, key *uint64) []byte {
	encryptedLi := make([]byte, len(text))
	for i := 0; i < len(text); i++ {
		encryptedLi[i] = text[i] ^ byte(*key)
	}
	*key = xorShift(*key)
	fmt.Println("key isssss: ", *key)
	return encryptedLi
}

func main() {

	// Input argument for host address
	arguments := os.Args[1:]
	if len(arguments) > 1 {
		fmt.Println("Error: only one argument accepted which is the server address and port")
		return
	}

	hostAddress := "localhost:8080"
	if len(arguments) == 1 {
		hostAddress = arguments[0]
	}

	// Creating connection with the server
	scanner := bufio.NewScanner(os.Stdin)
	conn, err := net.Dial("tcp", hostAddress)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	key := uint64(1343123213123434)

	defer conn.Close()
	for {
		fmt.Print("Enter message: ")
		if !scanner.Scan() {
			break
		}
		msg := scanner.Text()

		if msg == "exit" {
			break
		}

		start := time.Now()

		msgEncoded := xorEncodeDecode([]byte(msg), &key)
		fmt.Println(msgEncoded)
		_, err = conn.Write(msgEncoded)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		elapsed := time.Since(start)
		fmt.Println("Elapsed time:", elapsed)
	}
}

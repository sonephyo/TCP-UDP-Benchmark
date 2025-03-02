package main

import (
	// "bufio"
	"encoding/binary"
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
	return encryptedLi
}

func sendDataToClient(msg string, key *uint64, conn net.Conn, bufferSize int) {
	msgEncoded := xorEncodeDecode([]byte(msg), key)

	// Sending message length data
	lengthBuffer := make([]byte, 4)
	binary.BigEndian.PutUint32(lengthBuffer, uint32(len(msgEncoded)))
	conn.Write(lengthBuffer)

	// Sending ChuckSize data
	lengthChuck := make([]byte, 4)
	binary.BigEndian.PutUint32(lengthChuck, uint32(bufferSize))
	conn.Write(lengthChuck)


	for i := 0; i < len(msgEncoded); i += bufferSize {
		end := i + bufferSize
		if end > len(msgEncoded) {
			end = len(msgEncoded)
		}
		chunk := msgEncoded[i:end]
		conn.Write(chunk)
	}

	returnDataFromServer := make([]byte, len(msgEncoded))
	conn.Read(returnDataFromServer)
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
	// scanner := bufio.NewScanner(os.Stdin)
	conn, err := net.Dial("tcp", hostAddress)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	key := uint64(1343123213123434)

	defer conn.Close()
	bufferSizes := []int{8, 64, 256, 512}

	for _, longMessage := range LongMessages {
		fmt.Println("# Sending Message of : ", cropString(longMessage, 20))
		for _, value := range bufferSizes {
			start := time.Now()
			sendDataToClient(longMessage, &key, conn, value)
			elapsed := time.Since(start)
			fmt.Println("Elapsed time:", elapsed, "for the bufferSize of", value)
		}
		fmt.Println("--------------------------------------")
	}
	
	// for {
	// 	fmt.Print("Enter message: ")
	// 	if !scanner.Scan() {
	// 		break
	// 	}
	// 	msg := scanner.Text()

	// 	if msg == "exit" {
	// 		break
	// 	}

	// 	bufferSizes := []int{8, 64, 256, 512}

	// 	for _, value := range bufferSizes {
	// 		start := time.Now()
	// 		sendDataToClient(msg, &key, conn, value)
	// 		elapsed := time.Since(start)
	// 		fmt.Println("Elapsed time:", elapsed)
	// 		fmt.Println("----")
	// 	}

	// }
}

// Helper Functions
func cropString(s string, size int) string {
	if len(s) <= size {
		return s
	}
	return s[:size] + "..."
}
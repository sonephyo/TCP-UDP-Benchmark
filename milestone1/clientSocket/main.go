package main

import (
	// "bufio"
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/json"
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

	// 8-byte acknowledgement
	returnDataFromServer := make([]byte, 8)
	conn.Read(returnDataFromServer)
	hash := sha256.Sum256(msgEncoded)
	if bytes.Equal(hash[:8], returnDataFromServer) {
		fmt.Println("Acknowledgment recieved: data are equal")
	} else {
		fmt.Println("Warning: data recieved are inequal")
	}

}

func main() {

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
	conn, err := net.Dial("tcp", hostAddress)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	key := uint64(1343123213123434)

	defer conn.Close()
	bufferSizes := []int{8, 64, 256, 512}

	hashmap := make(map[string]map[int][]float64)

	for i := 1; i <= 100; i++ {
		for _, longMessage := range LongMessages {
			fmt.Println("# Sending Message of : ", helper.CropString(longMessage, 20))
			for _, value := range bufferSizes {
				start := time.Now()
				sendDataToClient(longMessage, &key, conn, value)
				elapsed := time.Since(start)
				fmt.Println("Elapsed time:", elapsed, "for the bufferSize of", value)
				if _, exists := hashmap[longMessage]; !exists {
					hashmap[longMessage] = make(map[int][]float64)
				}

				hashmap[longMessage][value] = append(hashmap[longMessage][value], elapsed.Seconds())

			}
			fmt.Println("--------------------------------------")

		}
	}

	file, err := os.Create("hashmap_data.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Convert the hashmap to JSON
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Optional: Pretty-print the JSON output
	err = encoder.Encode(hashmap)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}
}

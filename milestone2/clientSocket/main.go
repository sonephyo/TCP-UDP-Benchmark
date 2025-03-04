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
	"strconv"
	"strings"
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

func sendDataToClient(msgLi []string, key *uint64, conn net.Conn, bufferSize int) {

	// Sending message length data
	msgLength := make([]byte, 4)
	binary.BigEndian.PutUint32(msgLength, uint32(len(msgLi[0])))
	conn.Write(msgLength)

	// Sending ChuckSize data
	msgCount := make([]byte, 4)
	binary.BigEndian.PutUint32(msgCount, uint32(len(msgLi)))
	conn.Write(msgCount)

	for _, msg := range msgLi {
		msgEncoded := xorEncodeDecode([]byte(msg), key)
		conn.Write(msgEncoded)
	}

	// 8-byte acknowledgement
	returnDataFromServer := make([]byte, 8)
	conn.Read(returnDataFromServer)
	messageCombined := strings.Join(msgLi, "")
	hash := sha256.Sum256([]byte(messageCombined))
	if bytes.Equal(hash[:8], returnDataFromServer) {
		fmt.Println("Acknowledgment recieved: data are equal")
	} else {
		fmt.Println("Warning: Acknowledgement data recieved are inequal")
	}
}

func generate1MBStr(messageCount int, messageSize int) []string {
	var stringList []string
	for i := 0; i < messageCount; i++ {
		stringList = append(stringList, helper.GenerateRandomString(messageSize))
	}
	return stringList
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
	conn, err := net.Dial("tcp", hostAddress)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	key := uint64(1343123213123434)

	defer conn.Close()
	messageCounts := []int{1024, 2048, 4096}
	messageSizes := []int{1024, 512, 256}
	hashmap := make(map[string][]float64)

	for i := 1; i <= 100; i++ {
		for i := 0; i < len(messageCounts); i++ {
			messageLi := generate1MBStr(messageCounts[i], messageSizes[i])
			fmt.Println("First Data Snippet :", helper.CropString(messageLi[0], 20))

			start := time.Now()
			sendDataToClient(messageLi, &key, conn, messageSizes[i])

			elapsed := time.Since(start)
			fmt.Println("Elapsed time:", elapsed)

			str := strconv.Itoa(messageCounts[i]) + "x" + strconv.Itoa(messageSizes[i])
			if _, exists := hashmap[str]; !exists {
				hashmap[str] = make([]float64, 0)
			}
			hashmap[str] = append(hashmap[str], elapsed.Seconds())

		}
	}
	fmt.Println("--------------------------------------")
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

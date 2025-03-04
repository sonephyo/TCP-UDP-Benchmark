package main

import (
	"encoding/binary"
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

	lengthBuffer := make([]byte, 4)
	binary.BigEndian.PutUint32(lengthBuffer, uint32(len(msgEncoded)))
	_, err := conn.Write(lengthBuffer)
	if err != nil {
		fmt.Println("asdfasdfasdfasdf")
		fmt.Println("Error: ", err)
		return 
	}
	fmt.Println("ya recieved")

}


func main() {
	// Input argument for host address
	arguments := os.Args[1:]
	if len(arguments) > 1 {
		fmt.Println("Error: only one argument accepted which is the server address and port")
		return
	}

	// hostAddress := "127.0.0.1:63745"
	// if len(arguments) == 1 {
	// 	hostAddress = arguments[0]
	// }


	// serverAddr, err := net.ResolveUDPAddr("udp", hostAddress)
	// if err != nil {
	// 	fmt.Println("Error resolving UDP address:", err)
	// 	return
	// }

	// // Creating connection with the server
	// conn, err := net.DialUDP("udp", nil, serverAddr)
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// 	return
	// }
	conn, err := net.Dial("udp", "127.0.0.1:1234")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	key := uint64(1343123213123434)

	defer conn.Close()
	start := time.Now()
	sendDataToClient("hasdfasdfasdfasdfasdfi", &key, conn, 8)
	elapsed := time.Since(start)
	fmt.Println("Elapsed time:", elapsed, "for the bufferSize of", 8)
	// bufferSizes := []int{8, 64, 256, 512}

	// for _, longMessage := range LongMessages {
	// 	fmt.Println("# Sending Message of : ", helper.CropString(longMessage, 20))
	// 	for _, value := range bufferSizes {
	// 		start := time.Now()
	// 		sendDataToClient(longMessage, &key, conn, serverAddr, value)
	// 		elapsed := time.Since(start)
	// 		fmt.Println("Elapsed time:", elapsed, "for the bufferSize of", value)
	// 	}
	// 	fmt.Println("--------------------------------------")
	// }
}
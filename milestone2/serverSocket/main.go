package main

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"os"
	"project1/helper"
	"strings"
)


func xorEncodeDecode(text []byte, key *uint64) []byte {
    encryptedLi := make([]byte, len(text))
    for i:=0; i< len(text); i++ {
        encryptedLi[i] = text[i]^byte(*key)
    }
	*key = helper.XorShift(*key)

    return encryptedLi
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	key := uint64(1343123213123434)

	for {
		
		msgLength := make([]byte, 4)
		_, err := io.ReadFull(conn, msgLength)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Client closed Connection")
				return
			}
			fmt.Println("Read full Error: ", err)
			return 
		}

		messageLength := binary.BigEndian.Uint32(msgLength)
		if messageLength == 0 {
			fmt.Println("Empty message/ Invalid")
			return
		}

		msgCount := make([]byte, 4)
		_, err = io.ReadFull(conn, msgCount)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Client closed Connection")
				return
			}
			fmt.Println("Read full Error: ", err)
			return 
		}

		messageCount := binary.BigEndian.Uint32(msgCount)
		if messageCount == 0 {
			fmt.Println("Invalid chuckSize")
			return
		}

		var fullMessages []string
		fmt.Println("Message Length: ", messageLength)
		fmt.Println("Message Count: ", messageCount)

		for i := 0; i < int(messageCount); i++ {
			msgRecieved := make([]byte, messageLength)
			_, err := io.ReadFull(conn, msgRecieved)
			if err != nil {
				if err == io.EOF {
					fmt.Println("EOF reached before reading all messages")
				} else {
					fmt.Println("Error reading message:", err)
				}
				return
			}

			decodedMsg := xorEncodeDecode(msgRecieved, &key)

			fullMessages = append(fullMessages, string(decodedMsg))
		}
	
		// Sending 8-byte acknowledgement back to the client
		messageCombined := strings.Join(fullMessages, "")
		hash := sha256.Sum256([]byte(messageCombined))
		checksum := []byte(hash[:8])
		conn.Write(checksum)
	}
}


func main() {

	// Input argment for the host address
	arguments := os.Args[1:]
	if len(arguments) > 1 {
		fmt.Println("Error: only one argument accepted which is the server address and port")
		return
	}

	hostAddress := "localhost:8080"
	if len(arguments) == 1 {
		hostAddress = arguments[0]
	}

	// Creating a host
	listener, err := net.Listen("tcp", hostAddress)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	defer listener.Close()

	fmt.Println("Server is listening on port ", hostAddress)

	// Listening to incoming data
	for {
		conn, err := listener.Accept()
		
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}

		go handleClient(conn)
	}
}

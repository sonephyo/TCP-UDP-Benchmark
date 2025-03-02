package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"os"
)

func xorShift(r uint64) uint64 {
	r ^= r << 13
	r ^= r >> 7
	r ^= r << 17
	return r
}

func xorEncodeDecode(text []byte, key *uint64) []byte {
    encryptedLi := make([]byte, len(text))
    for i:=0; i< len(text); i++ {
        encryptedLi[i] = text[i]^byte(*key)
    }
	*key = xorShift(*key)

    return encryptedLi
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	key := uint64(1343123213123434)

	for {
		
		lengthBuffer := make([]byte, 4)
		_, err := io.ReadFull(conn, lengthBuffer)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Client closed Connection")
				return
			}
			fmt.Println("Read full Error: ", err)
			return 
		}

		messageLength := binary.BigEndian.Uint32(lengthBuffer)
		if messageLength == 0 {
			fmt.Println("Empty message/ Invalid")
			return
		}

		lengthChuck := make([]byte, 4)
		_, err = io.ReadFull(conn, lengthChuck)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Client closed Connection")
				return
			}
			fmt.Println("Read full Error: ", err)
			return 
		}

		chuckLength := binary.BigEndian.Uint32(lengthChuck)
		if chuckLength == 0 {
			fmt.Println("Invalid chuckSize")
			return
		}

		var fullMessage []byte
		fmt.Println("Message Length: ", messageLength)
		fmt.Println("ChuckLength: ", chuckLength)

		for uint32(len(fullMessage)) < messageLength {
			remaining := messageLength - uint32(len(fullMessage))
			
			currentChunkSize := chuckLength
			if currentChunkSize > remaining {
				currentChunkSize = remaining
			}

			chunk := make([]byte, currentChunkSize)
			n, err:= conn.Read(chunk)
			if err != nil {
				if err == io.EOF && uint32(len(fullMessage)) == messageLength {
					break
				}
				fmt.Println("Error reading chuck: ", err)
				break
			}

			fullMessage = append(fullMessage, chunk[:n]...)
		}
		
		decodedBytes := xorEncodeDecode(fullMessage, &key)
		fmt.Printf("Received Letter: %s \n", string(decodedBytes[:]))


		// Sending encrypted data back to the client
		conn.Write(fullMessage)

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


package main

import (
	"fmt"
	"net"
	"os"
	"io"
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

	// Creating buffer to read data
	buffer := make([]byte, 1024)
	key := uint64(1343123213123434)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Client disconnected")
			} else {
				fmt.Println("Read error:", err)
			}
			return
		}
		decodedBytes := xorEncodeDecode(buffer[:n], &key)
		fmt.Printf("Bytes read: %d, Buffer length: %d \n", n, len(buffer))
		fmt.Printf("Received Letter: %s \n", string(decodedBytes[:]))
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


package main

import (
	// "crypto/sha256"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"os"
	"project1/helper"
	"strconv"
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

func handleClient(conn *net.UDPConn, clientAddress *net.UDPAddr, data []byte) {
	// buffer := make([]byte, 2048)

	// key := uint64(1343123213123434)

	// for {
	

		messageLength := binary.BigEndian.Uint32(data)
		if messageLength == 0 {
			fmt.Println("Empty message/ Invalid")
			return
		}

		fmt.Println("LengthBuffer:", messageLength)

		// messageLength := binary.BigEndian.Uint32(lengthBuffer)
		// if messageLength == 0 {
		// 	fmt.Println("Empty message/ Invalid")
		// 	return
		// }

		// lengthChuck := make([]byte, 4)
		// _, err = io.ReadFull(conn, lengthChuck)
		// if err != nil {
		// 	if err == io.EOF {
		// 		fmt.Println("Client closed Connection")
		// 		return
		// 	}
		// 	fmt.Println("Read full Error: ", err)
		// 	return 
		// }

		// chuckLength := binary.BigEndian.Uint32(lengthChuck)
		// if chuckLength == 0 {
		// 	fmt.Println("Invalid chuckSize")
		// 	return
		// }

		// var fullMessage []byte
		// fmt.Println("Message Length: ", messageLength)
		// fmt.Println("ChuckLength: ", chuckLength)

		// for uint32(len(fullMessage)) < messageLength {
		// 	remaining := messageLength - uint32(len(fullMessage))
			
		// 	currentChunkSize := chuckLength
		// 	if currentChunkSize > remaining {
		// 		currentChunkSize = remaining
		// 	}

		// 	chunk := make([]byte, currentChunkSize)
		// 	n, _, err:= conn.ReadFromUDP(chunk)
		// 	if err != nil {
		// 		if err == io.EOF && uint32(len(fullMessage)) == messageLength {
		// 			break
		// 		}
		// 		fmt.Println("Error reading chuck: ", err)
		// 		break
		// 	}

		// 	fullMessage = append(fullMessage, chunk[:n]...)
		// }
		
		// decodedBytes := xorEncodeDecode(fullMessage, &key)
		// fmt.Printf("Received Letter: %s \n", string(decodedBytes[:]))


		// // Sending 8-byte acknowledgement back to the client
		// hash := sha256.Sum256(fullMessage)
		// checksum := []byte(hash[:8])
		// _, err = conn.Write(checksum)
		// if err != nil {
		// 	fmt.Println("Error sending acknowledgement: ", err)
		// 	return
		// }

	// }
}


func main() {

	// Input argment for the host address
	arguments := os.Args[1:]
	if len(arguments) > 1 {
		fmt.Println("Error: only one argument accepted which is the server address and port")
		return
	}

	hostAddress := "127.0.0.1:1234"
	if len(arguments) == 1 {
		hostAddress = arguments[0]
	}

	hostAddressLi := strings.Split(hostAddress, ":")
	if len(hostAddressLi) != 2 {
		fmt.Println("Error: Invalid hostAddressFormat")
		return
	}

	port, err := strconv.Atoi(hostAddressLi[1])
	if err != nil {
		fmt.Println("Error: Invalid Port Number")
	}

	address := net.UDPAddr {
		Port: port,
		IP: net.ParseIP(hostAddressLi[0]),
	}

	listener, err := net.ListenUDP("udp", &address)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	defer listener.Close()

	fmt.Println("Server is listening on port ", hostAddress)

	for {
		buffer := make([]byte, 65535)
		n, clientAddress, err  := listener.ReadFromUDP(buffer)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Client closed Connection")
				return
			}
			fmt.Println("Read full Error: ", err)
			return 
		}

		data := make([]byte, n)
		copy(data, buffer[:n])
		go handleClient(listener, clientAddress, data)
	}
}

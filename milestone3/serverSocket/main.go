// package main

// import (
// 	"encoding/binary"
// 	"fmt"
// 	"io"
// 	"net"
// 	"os"
// 	"project1/helper"
// 	"strconv"
// 	"strings"
// 	"sync"
// )

// type ClientData struct {
// 	messageLength uint32
// 	lengthChuck   uint32
// 	fullMessage   []byte
// 	key           uint64
// }

// var clientDataMap = make(map[string]*ClientData)
// var mapLock sync.Mutex
// var i int = 0

// func xorEncodeDecode(text []byte, key *uint64) []byte {
// 	encryptedLi := make([]byte, len(text))
// 	for i := 0; i < len(text); i++ {
// 		encryptedLi[i] = text[i] ^ byte(*key)
// 	}
// 	*key = helper.XorShift(*key)
// 	return encryptedLi
// }

// func handleClient(conn *net.UDPConn, clientAddress *net.UDPAddr, data []byte) {
// 	keyStr := clientAddress.String()
// 	mapLock.Lock()
// 	if _, exists := clientDataMap[keyStr]; !exists {
// 		clientData := &ClientData{
// 			key: uint64(1343123213123434),
// 		}
// 		clientDataMap[keyStr] = clientData
// 	}
// 	clientData := clientDataMap[keyStr]
// 	fmt.Println(clientData)
// 	fmt.Println("Current i : ", i)
// 	i++
// 	fmt.Println("----")
// 	if clientData.messageLength == 0 {
// 		fmt.Println("lengthBuffer: ", data)

// 		ml := binary.BigEndian.Uint32(data)
// 		if ml == 0 {
// 			fmt.Println("Empty message/ Invalid")
// 			mapLock.Unlock()
// 			return
// 		}
// 		clientData.messageLength = ml
// 		fmt.Println("Message Length: ", clientData.messageLength)
// 		mapLock.Unlock()
// 		return
// 	} else if clientData.lengthChuck == 0 {
// 		fmt.Println("lengthChuck: ", data)

// 		lc := binary.BigEndian.Uint32(data)
// 		if lc == 0 {
// 			fmt.Println("Empty chunk length/ Invalid")
// 			mapLock.Unlock()
// 			return
// 		}
// 		clientData.lengthChuck = lc
// 		fmt.Println("Length Chuck: ", clientData.lengthChuck)
// 		mapLock.Unlock()
// 		return
// 	} else if uint32(len(clientData.fullMessage)) < clientData.messageLength {
// 		clientData.fullMessage = append(clientData.fullMessage, data...)
// 		fmt.Println(len(clientData.fullMessage))
// 		if uint32(len(clientData.fullMessage)) >= clientData.messageLength {
// 			decodedBytes := xorEncodeDecode(clientData.fullMessage, &clientData.key)
// 			fmt.Printf("Received Letter: %s \n", string(decodedBytes))
// 		}
// 		mapLock.Unlock()
// 		return
// 	}
// 	mapLock.Unlock()
// }

// func main() {
// 	arguments := os.Args[1:]
// 	if len(arguments) > 1 {
// 		fmt.Println("Error: only one argument accepted which is the server address and port")
// 		return
// 	}
// 	hostAddress := "127.0.0.1:1234"
// 	if len(arguments) == 1 {
// 		hostAddress = arguments[0]
// 	}
// 	hostAddressLi := strings.Split(hostAddress, ":")
// 	if len(hostAddressLi) != 2 {
// 		fmt.Println("Error: Invalid hostAddressFormat")
// 		return
// 	}
// 	port, err := strconv.Atoi(hostAddressLi[1])
// 	if err != nil {
// 		fmt.Println("Error: Invalid Port Number")
// 		return
// 	}
// 	address := net.UDPAddr{
// 		Port: port,
// 		IP:   net.ParseIP(hostAddressLi[0]),
// 	}
// 	listener, err := net.ListenUDP("udp", &address)
// 	if err != nil {
// 		fmt.Println("Error: ", err)
// 		return
// 	}
// 	defer listener.Close()
// 	fmt.Println("Server is listening on port ", hostAddress)
// 	for {
// 		buffer := make([]byte, 65535)
// 		n, clientAddress, err := listener.ReadFromUDP(buffer)
// 		if err != nil {
// 			if err == io.EOF {
// 				fmt.Println("Client closed Connection")
// 				return
// 			}
// 			fmt.Println("Read full Error: ", err)
// 			return
// 		}
// 		data := make([]byte, n)
// 		copy(data, buffer[:n])
// 		go handleClient(listener, clientAddress, data)
// 	}
// }

package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"os"
	"project1/helper"
	"strconv"
	"strings"
	"sync"
)

type ClientData struct {
	messageLength uint32
	lengthChuck   uint32
	fullMessage   []byte
	key           uint64
}

var clientDataMap = make(map[string]*ClientData)
var mapLock sync.Mutex
var i int = 0

func xorEncodeDecode(text []byte, key *uint64) []byte {
	encryptedLi := make([]byte, len(text))
	for i := 0; i < len(text); i++ {
		encryptedLi[i] = text[i] ^ byte(*key)
	}
	*key = helper.XorShift(*key)
	return encryptedLi
}

func handleClient(conn *net.UDPConn, clientAddress *net.UDPAddr, data []byte) {
	keyStr := clientAddress.String()
	mapLock.Lock()
	if _, exists := clientDataMap[keyStr]; !exists {
		clientData := &ClientData{
			key: uint64(1343123213123434),
		}
		clientDataMap[keyStr] = clientData
	}
	clientData := clientDataMap[keyStr]
	fmt.Println(clientData)
	fmt.Println("Current i : ", i)
	i++
	fmt.Println("----")

	// Read message length
	if clientData.messageLength == 0 {
		fmt.Println("lengthBuffer: ", data)
		ml := binary.BigEndian.Uint32(data)
		if ml == 0 {
			fmt.Println("Empty message/ Invalid")
			mapLock.Unlock()
			return
		}
		clientData.messageLength = ml
		fmt.Println("Message Length: ", clientData.messageLength)
		mapLock.Unlock()
		return
	}

	// Read chunk length
	if clientData.lengthChuck == 0 {
		fmt.Println("lengthChuck: ", data)
		lc := binary.BigEndian.Uint32(data)
		if lc == 0 {
			fmt.Println("Empty chunk length/ Invalid")
			mapLock.Unlock()
			return
		}
		clientData.lengthChuck = lc
		fmt.Println("Length Chuck: ", clientData.lengthChuck)
		mapLock.Unlock()
		return
	}

	// Append chunks to the full message
	if uint32(len(clientData.fullMessage)) < clientData.messageLength {
		clientData.fullMessage = append(clientData.fullMessage, data...)
		fmt.Println("Full message size:", len(clientData.fullMessage))
		fmt.Println(len(clientData.fullMessage), clientData.fullMessage)
		if uint32(len(clientData.fullMessage)) >= clientData.messageLength {
			decodedBytes := xorEncodeDecode(clientData.fullMessage, &clientData.key)
			fmt.Printf("Received Letter: %s \n", string(decodedBytes))
		}
	}
	mapLock.Unlock()
}

func main() {
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
		return
	}
	address := net.UDPAddr{
		Port: port,
		IP:   net.ParseIP(hostAddressLi[0]),
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
		n, clientAddress, err := listener.ReadFromUDP(buffer)
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

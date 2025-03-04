package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"net"
	"os"
	"project1/helper"
	"strconv"
	"strings"
	"sync"
)

var mapLock sync.Mutex

func xorEncodeDecode(text []byte, key *uint64) []byte {
	encryptedLi := make([]byte, len(text))
	for i := 0; i < len(text); i++ {
		encryptedLi[i] = text[i] ^ byte(*key)
	}
	*key = helper.XorShift(*key)
	return encryptedLi
}

func handleClient(conn *net.UDPConn, clientAddress *net.UDPAddr, data []byte) {
	mapLock.Lock()
	key := uint64(1343123213123434)
	decodedBytes := xorEncodeDecode(data, &key)
	fmt.Printf("Received Letter: %s \n", helper.CropString(string(decodedBytes[:]), 20))
	fmt.Println("Size : ", len(decodedBytes))


	hash := sha256.Sum256([]byte("Data recieved"))
	checksum := []byte(hash[:8])
	conn.WriteToUDP(checksum, clientAddress)
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

package main

import (
	"fmt"

	"golang.org/x/sys/windows"
)

func main() {
	// Initialize Winsock
	wsadata := new(windows.WSAData)
	err := windows.WSAStartup(uint32(0x202), wsadata)
	if err != nil {
		panic(fmt.Errorf("WSAStartup failed: %v", err))
	}
	defer windows.WSACleanup()

	// Create a UDP socket
	fd, err := windows.Socket(windows.AF_INET, windows.SOCK_DGRAM, windows.IPPROTO_UDP)
	if err != nil {
		panic(fmt.Errorf("Socket creation failed: %v", err))
	}
	defer windows.Closesocket(fd)

	serverAddr := &windows.SockaddrInet4{Port: 9999}
	copy(serverAddr.Addr[:], []byte{127, 0, 0, 1}) // 127.0.0.1

	message := make([]byte, 1024)
	fmt.Scan(&message)

	err = windows.Sendto(fd, message, 0, serverAddr)
	if err != nil {
		panic(fmt.Errorf("Sendto failed: %v", err))
	}
	fmt.Println("Message sent to server")

	// Receive a response
	buffer := make([]byte, 1024)
	n, _, err := windows.Recvfrom(fd, buffer, 0)
	if err != nil {
		panic(fmt.Errorf("Recvfrom failed: %v", err))
	}

	fmt.Printf("Server response: %s\n", buffer[:n])
}

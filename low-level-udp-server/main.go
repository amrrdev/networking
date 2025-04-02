package main

import (
	"fmt"
	"strings"

	"golang.org/x/sys/windows"
)

func main() {

	wsadata := new(windows.WSAData)
	err := windows.WSAStartup(uint32(0x202), wsadata)
	if err != nil {
		panic(fmt.Errorf("WSAStartup failed: %v", err))
	}
	defer windows.WSACleanup()

	fd, err := windows.Socket(windows.AF_INET, windows.SOCK_DGRAM, windows.IPPROTO_UDP)
	if err != nil {
		panic(err)
	}
	defer windows.Close(fd)

	addr := windows.SockaddrInet4{
		Port: 9999,
		Addr: [4]byte{127, 0, 0, 1},
	}

	if err := windows.Bind(fd, &addr); err != nil {
		panic(err)
	}

	fmt.Println("UDP server listening on :9999")

	buff := make([]byte, 1024)
	n, fromAddress, err := windows.Recvfrom(fd, buff, 0)
	if err != nil {
		panic(err)
	}

	// clientAddr := fromAddress.(*windows.SockaddrInet4)
	fmt.Printf("Received from %v: %s\n", fromAddress, buff[:n])

	resp := []byte(strings.ToUpper(string(buff[:n])))

	if err := windows.Sendto(fd, resp, 0, fromAddress); err != nil {
		panic(err)
	}

}

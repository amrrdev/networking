package main

import (
	"fmt"

	"golang.org/x/sys/windows"
)

func main() {
	wsadata := new(windows.WSAData)
	err := windows.WSAStartup(uint32(0x202), wsadata)
	if err != nil {
		panic(fmt.Errorf("WSAStartup failed: %v", err))
	}
	defer windows.WSACleanup()

	fd, err := windows.Socket(windows.AF_INET, windows.SOCK_STREAM, windows.IPPROTO_TCP)
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

	if err := windows.Listen(fd, 10); err != nil {
		panic(err)
	}

	fmt.Println("Server listening on http://127.0.0.1:9999")

	for {
		// https://github.com/golang/go/issues/54212
		connFd, _, err := windows.Accept(fd)
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(connFd)
	}
}

func handleConnection(connFd windows.Handle) {
	defer windows.Close(connFd)
	buff := make([]byte, 10000000)
	n, err := windows.Read(connFd, buff)
	if err != nil {
		fmt.Println("Error reading from connection:", err)
		return
	}

	_, err = windows.Write(connFd, buff[:n])
	if err != nil {
		fmt.Println("Error writing response:", err)
	}
}

package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

const (
	network = "tcp"
	adress  = ":8080"
)

func main() {
	listener, err := net.Listen(network, adress)
	if err != nil {
		fmt.Printf("Error starting TCP server: %v", err)
		return
	}
	defer listener.Close()
	fmt.Printf("Listening on %v", adress)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection: %v", err)
			continue
		}

		go handlerConnection(conn)
	}
}

func handlerConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading from connection: %v", err)
			return
		}

		fmt.Printf("Message received: %v", message)
		newMessage := strings.ToUpper(message)

		_, err = conn.Write([]byte(newMessage))
		if err != nil {
			fmt.Printf("Error writing to connection: %v", err)
		}
	}
}

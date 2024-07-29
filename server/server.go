package main

import (
	"bufio"
	"log"
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
		log.Fatalf("error starting TCP server: %v", err)
	}
	defer listener.Close()
	log.Printf("Listening on %v", adress)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("error accepting connection: %v", err)
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
			log.Printf("error reading from connection: %v", err)
			return
		}

		log.Printf("Message received: %v", message)
		newMessage := strings.ToUpper(message)

		_, err = conn.Write([]byte(newMessage))
		if err != nil {
			log.Printf("error writing to connection: %v", err)
			return
		}
	}
}

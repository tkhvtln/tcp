package main

import (
	"log"
	"net"
	"strings"
)

const (
	network = "udp"
	adress  = ":8080"
)

func main() {
	listener, err := net.ListenPacket(network, adress)
	if err != nil {
		log.Fatalf("error starting TCP server: %v", err)
	}
	defer listener.Close()
	log.Printf("Listening on %v", adress)

	buffer := make([]byte, 1024)
	for {
		n, addr, err := listener.ReadFrom(buffer)
		if err != nil {
			log.Printf("error accepting connection: %v", err)
			continue
		}

		go handlerConnection(listener, addr, buffer[:n])
	}
}

func handlerConnection(conn net.PacketConn, addr net.Addr, message []byte) {
	log.Printf("Message received from %v: %v", addr, string(message))
	newMessage := strings.ToUpper(string(message))

	_, err := conn.WriteTo([]byte(newMessage), addr)
	if err != nil {
		log.Printf("error writing to connection: %v", err)
	}
}

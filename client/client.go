package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

const (
	network      = "udp"
	adress       = ":8080"
	readTimeout  = 10 * time.Second
	writeTimeout = 10 * time.Second
)

func main() {
	conn, err := net.Dial(network, adress)
	if err != nil {
		log.Fatalf("error connecting to server: %v", err)
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	buffer := make([]byte, 1024)

	for {
		fmt.Print("Enter text: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("error reading input: %v\n", err)
			continue
		}

		conn.SetWriteDeadline(time.Now().Add(writeTimeout))
		_, err = conn.Write([]byte(text))
		if err != nil {
			log.Printf("error sending to server: %v\n", err)
			continue
		}

		conn.SetReadDeadline(time.Now().Add(readTimeout))
		n, err := conn.Read(buffer)
		if err != nil {
			log.Printf("error reading from server: %v\n", err)
			continue
		}

		message := string(buffer[:n])
		log.Printf("Message from server: %v", message)
	}
}

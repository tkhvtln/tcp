package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

const (
	network      = "tcp"
	adress       = ":8080"
	readTimeout  = 10 * time.Second
	writeTimeout = 10 * time.Second
)

func main() {
	conn, err := net.Dial(network, adress)
	if err != nil {
		fmt.Printf("Error connecting to server: %v", err)
		return
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	connReader := bufio.NewReader(conn)

	for {
		fmt.Print("Enter text: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading input: %v", err)
			return
		}

		conn.SetWriteDeadline(time.Now().Add(writeTimeout))
		_, err = fmt.Fprintf(conn, text)
		if err != nil {
			fmt.Printf("Error sending to server: %v", err)
			return
		}

		conn.SetReadDeadline(time.Now().Add(readTimeout))
		message, err := connReader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading from server: %v", err)
			return
		}

		fmt.Printf("Message from server: %v", message)
	}
}

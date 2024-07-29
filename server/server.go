package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

const (
	network = "tcp"
	adress  = ":8080"
)

func main() {
	defer closeDB()
	err := initDB()
	if err != nil {
		log.Fatal(err)
	}

	listener, err := net.Listen(network, adress)
	if err != nil {
		log.Fatalf("error starting TCP server: %v\n", err)
	}
	defer listener.Close()
	log.Printf("Listening on %v\n", adress)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("error accepting connection: %v\n", err)
			continue
		}

		go func(c net.Conn) {
			defer c.Close()
			showAllUsers(c)
			handlerConnection(c)
		}(conn)
	}
}

func handlerConnection(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("error reading from connection: %v\n", err)
			return
		}

		message = strings.TrimSpace(message)
		if message == "" {
			log.Println("Received empty message")
			conn.Write([]byte("Received empty message\n"))
			continue
		}
		log.Printf("Message received: %v", message)

		idUser, err := strconv.Atoi(message)
		if err != nil {
			log.Printf("Invalid input: %v\n", err)
			conn.Write([]byte("Invalid input\n"))
			continue
		}

		userInfo, err := getUserInfo(idUser)
		if err != nil {
			log.Print(err)
			userInfo = err.Error() + "\n"
		}

		_, err = conn.Write([]byte(userInfo))
		if err != nil {
			fmt.Printf("error writing to connection: %v\n", err)
			return
		}
	}
}

func showAllUsers(conn net.Conn) {
	users, err := getAllUsers()
	if err != nil {
		log.Println(err)
		users = err.Error()
	}

	_, err = conn.Write([]byte(users + "\n"))
	if err != nil {
		log.Printf("error writing to connection: %v\n", err)
	}

}

package main

import (
	"fmt"
	"log"
	"net"
)

func handleConnection(connection net.Conn) {

	buf := make([]byte, 1024)
	_, err := connection.Read(buf)
	if err != nil {
			fmt.Println("Error reading connection", err.Error())
			return
	}

  fmt.Println(string(buf))

	connection.Write(buf)
	connection.Close()
}

func main() {
  listener, err := net.Listen("tcp", ":8080")
  
  if err != nil {
		log.Fatal("could not create the listener", err.Error())
	}

	fmt.Println("Server started")

	for {
		fmt.Println("listening")
		connection, err := listener.Accept()
		if err != nil {
			fmt.Println("Connection error", err.Error())
		}

		go handleConnection(connection)
	}
}

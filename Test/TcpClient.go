package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")

	if err != nil {
		log.Fatal("connection error", err.Error())
	}

	fmt.Println("Connection established")

	_, err = conn.Write([]byte("Hello, server!"))
	if err != nil {
			fmt.Println(err)
			return
	}

	buf := make([]byte, 1024)
	
	_, err = conn.Read(buf)

	fmt.Println(string(buf))

	conn.Close()
}

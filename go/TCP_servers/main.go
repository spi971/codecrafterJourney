package main

import (
	"log"
	"net"
	"time"
)

func do(conn net.Conn) {
	buffer := make([]byte, 1024)

	_, err := conn.Read(buffer)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("processing the request")
	time.Sleep(8 * time.Second)

	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\nWah gwan\r\n"))
	conn.Close()
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		log.Println("waiting for a client to connect")
		conn, err := listener.Accept()

		if err != nil {
			log.Fatal(err)
		}
		log.Println("client connected")
		go do(conn)
	}
}

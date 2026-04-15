package main

import (
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	go Broadcaster()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go Handler(conn)
	}

}

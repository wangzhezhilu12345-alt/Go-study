package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		log.Print("done")
		done <- struct{}{} //指示主goroutine
	}() //匿名函数立马调用

	mustCopy(conn, os.Stdin)
	<-done
}

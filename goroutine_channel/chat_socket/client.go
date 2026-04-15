package main

import (
	"bufio"
	"fmt"
	"net"
)

func Handler(conn net.Conn) {
	ch := make(chan string)

	go Clientwriter(conn, ch)
	who := conn.RemoteAddr().String()
	//这个who大胆猜测应该是远端ip地址。
	ch <- "you are" + who
	messages <- who + " has arrived"
	entering <- ch

	input := bufio.NewScanner(conn) //包装成一个扫描器
	for input.Scan() {              //意思是一直发的话就一直读
		messages <- who + input.Text() //把消息发到共享的管道里
	}
	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}

func Clientwriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

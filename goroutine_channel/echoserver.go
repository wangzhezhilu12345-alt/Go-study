package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, msg string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(msg))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", msg)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(msg))
}

func handler(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		echo(c, input.Text(), 1*time.Second)
	}
	c.Close()
}

func mustCopy(dst io.Writer, src io.Reader) {
	_, err := io.Copy(dst, src)
	if err != nil {
		log.Fatal(err)
	}

}

// func main() {
// 	conn, err := net.Dial("tcp", "loaclhost:8080")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer conn.Close()
// 	go mustCopy(os.Stdout, conn) //这个参数是tm的接口....
// 	mustCopy(conn, os.Stdin)
// }

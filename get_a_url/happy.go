package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url) //利用GET产生一个HTTP请求
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch:%v\n", err)
		}
		b, err := ioutil.ReadAll(resp.Body) //把整个响应结果读取并且存入b
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetching %s %v", url, err)
			os.Exit(0)
		}
		fmt.Print(b)
	}
}

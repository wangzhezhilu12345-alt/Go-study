package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func fetch(url string, ch chan string) {
	start := time.Now() //计算耗时
	rescult, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) //发送通道到ch
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, rescult.Body) //这个函数的作用是把值从后面拷贝到前面
	rescult.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) //启动第一个go routine
	}
	fmt.Printf("%.2fs\n", time.Since(start).Seconds())
}

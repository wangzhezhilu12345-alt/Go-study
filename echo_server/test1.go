package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8888", nil)) //nb了，这一行顶c++ 100多行....
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.PATH= %q\n", r.URL.Path) //w是http响应 r 是http请求
}

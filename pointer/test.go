package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "commit error")
var sep = flag.String("broken", "", "error") //这是开一个指针

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}

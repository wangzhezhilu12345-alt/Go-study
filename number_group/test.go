package main

import "fmt"

func main() {
	var a [3]int
	fmt.Print(a[0])
	for i, v := range a {
		fmt.Print(i, v)
	}
	for _, v := range a {
		fmt.Print(v)
	}
}

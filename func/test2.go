package main

import "fmt"

func func_test2() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {
	f := func_test2()
	fmt.Println(f())
}

package main

import "fmt"

func square(n int) int {
	return n * n
}

func func_test1() {
	f := square
	fmt.Print(f(3))
}

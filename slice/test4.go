package main

import "fmt"

func Rotate(a []int) {
	first := a[0]
	for i := 0; i < len(a)-1; i++ {
		a[i] = a[i+1]
	}
	a[len(a)-1] = first
}

func demoRotate() {
	var b = []int{1, 2, 3, 4}
	Rotate(b)
	fmt.Print(b)
}

package main

import (
	"fmt"
)

func test3() {
	funcs := []func(){}

	for i := 0; i < 3; i++ {
		i := i
		funcs = append(funcs, func() {
			fmt.Println(i)
		})
	}

	for _, f := range funcs {
		f()
	}
}

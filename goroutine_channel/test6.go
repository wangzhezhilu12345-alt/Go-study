package main

import "fmt"

func test6() {
	ch := make(chan int, 1)

	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch<-i:
		}
	}
}
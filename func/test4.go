package main

import "fmt"

func test4() {
	fmt.Println("start")

	x := 1

	// defer 1（值在这里就确定了）
	defer fmt.Println("defer1:", x)

	// defer 2（闭包，会读取最新的 x）
	defer func() {
		fmt.Println("defer2:", x)
	}()

	x = 2

	fmt.Println("end")
}

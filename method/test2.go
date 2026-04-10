package main

import "fmt"

func test2() {
	q := Point{1, 2}
	p := Point{3, 4}
	distancefromp := p.Distance
	fmt.Println(distancefromp(q))

}

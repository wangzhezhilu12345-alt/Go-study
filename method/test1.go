package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

// a common func
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// a method to Point
func (p *Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Print(Distance(p, q))
	fmt.Print(q.Distance(p))
}

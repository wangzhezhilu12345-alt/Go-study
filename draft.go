package draft

type Point struct {
	x, y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

var w Wheel

//w.Radius=5
//这样可以直接用，就是一种匿名

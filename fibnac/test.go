package main

func main() {
	x, y := 0, 1
	for i := 0; i < 10; i++ {
		x, y = y, x+y
	}
}

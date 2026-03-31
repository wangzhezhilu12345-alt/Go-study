package main

type A float64
type B float64

const (
	absolute  A = -273.15
	Freezingc A = 0
)

func Ctof(c A) B {
	return B(c*9/5 + 32)
}

func ftoc(c B) A {
	return A(c - 32*5/9)
}

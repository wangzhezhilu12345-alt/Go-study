package main

var global *int

func f() {
	x := 15
	global = &x
}

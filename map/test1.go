package main

func equal(x, y map[string]int) bool {

	if len(x) != len(y) {
		return false
	}
	for k, v := range x {
		yv, ok := y[k] //这个值赋给了yv
		if !ok || yv != v {
			return false
		}

	}
	return true
}

func demoEqual() {
	var a, b map[string]int
	equal(a, b)
}

package main

import "fmt"

// Clear removes adjacent duplicates from a sorted slice in place.
func Clear(a []int) []int {
	if len(a) == 0 {
		return a
	}

	j := 0
	for i := 1; i < len(a); i++ {
		if a[i] != a[j] {
			j++
			a[j] = a[i]
		}
	}

	return a[:j+1]
}

func demoClear() {
	b := []int{1, 1, 2, 2, 3, 3, 4, 4}
	fmt.Print(Clear(b))
}

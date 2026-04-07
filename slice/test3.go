package main

func reverse(tmp *[]int) []int {
	var s = *tmp
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
	return s
}

func main() {
	var a []int
	reverse(&a)
}

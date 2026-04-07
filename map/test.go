package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := make(map[string]int)
	var names []string
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Println(name)
	}

}

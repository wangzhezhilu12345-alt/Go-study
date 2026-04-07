package main

var graph = make(map[string]map[string]bool)

func addegde(from, to string) {
	graph[from][to] = true
	//这里比较流氓，你注意第二个map是否存在
}

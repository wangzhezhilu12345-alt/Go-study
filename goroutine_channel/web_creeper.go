package main

import (
	"fmt"
	"its_mygo/links"
	"os"
)

var tokens = make(chan struct{}, 20)

func crawl(rawURL string) []string {
	fmt.Println(rawURL)

	tokens <- struct{}{}
	list, err := links.Extract(rawURL)
	<-tokens

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return list
}

func main() {
	worklist := make(chan []string)
	unseelinks := make(chan string)

	go func() {
		worklist <- os.Args[1:]

	}()

	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseelinks {
				foundlinks := crawl(link)
				go func() {
					worklist <- foundlinks
				}()
			}
		}()
	}

	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseelinks <- link
			}
		}
	}
}

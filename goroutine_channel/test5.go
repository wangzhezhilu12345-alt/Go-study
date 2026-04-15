package main


// func test5(filenames <-chan string) int64 {

// 	sizes := make(chan int64)

// 	var wg sync.WaitGroup //工作group的个数
// 	for f := range filenames {
// 		wg.Add(1)
// 		go func(f string) {
// 			defer wg.Done()

// 		}(f)
// 	}
// 	return 0
// }

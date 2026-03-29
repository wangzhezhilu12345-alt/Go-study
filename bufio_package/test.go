package main

import (
	"bufio"
	"fmt"
	"os"
)

func countLine(f *os.File, cnts map[string][]string, filename string) {
	input := bufio.NewScanner(f) //从文件读取
	for input.Scan() {
		cnts[input.Text()] = append(cnts[input.Text()], filename) //gpt给出的解释是，用取出来的字符当作键
	}
}
func main() {
	counts := make(map[string][]string)
	files := os.Args[1:] //agrs命令行参数，这个在os中学到了
	if len(files) == 0 {
		countLine(os.Stdin, counts, "stdin")
	} else {
		for _, arg := range files { //用一个下标和值去遍历命令行
			f, err := os.Open(arg) //os系统接口
			if err != nil {        //如果错误码是我们内置好的，就说明当前没有出错
				fmt.Fprintf(os.Stderr, "dup: %v\n", err) //打印到标准错误
				continue
			}
			countLine(f, counts, arg) //从文件f里面读取
			f.Close()
		}

		for line, n := range counts {
			fmt.Printf("%s %s\n", n, line)
		}
	}
}

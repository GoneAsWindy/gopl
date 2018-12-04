package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	names := make(map[string]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, nil)
	} else {
		for _, arg := range files {
			file, e := os.Open(arg)
			if e != nil {
				fmt.Fprintf(os.Stderr, "dup2:%v\n", e)
				continue
			}
			countLines(file, counts, names)
			file.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, names[line])
		}

	}
}

// 这里的counts是map类型，map 作为为参数传递给某函数时，该函数
//接收这个引用的一份拷贝，类似于java的引用传递
func countLines(f *os.File, counts map[string]int, names map[string]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if names != nil {
			names[input.Text()] = f.Name()
		}
	}
	// 这里忽略了input.Err()可能产生的错误
}

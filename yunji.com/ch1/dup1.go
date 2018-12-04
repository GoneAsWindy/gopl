package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		// go 需要自己进行退出的逻辑判断
		if input.Text() == "exit" {
			break
		}
		counts[input.Text()]++
	}
	// range 用在map时其第一个参数应该是key，而不再是索引
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

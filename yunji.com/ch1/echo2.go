package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)

	for idx, arg := range os.Args[1:] {
		fmt.Println(strconv.Itoa(idx+1) + " " + arg)
	}
}

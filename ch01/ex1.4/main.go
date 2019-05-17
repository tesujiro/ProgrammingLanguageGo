package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	files := make(map[string][]string)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
			files[line] = append(files[line], filename)
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d(%v)\t%s\n", n, files[line], line)
		}
	}
}

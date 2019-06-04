package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	count := make(map[string]int) // a set of strings
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		count[word]++
	}
	for w := range count {
		fmt.Printf("%v\t%v\n", w, count[w])
	}
}

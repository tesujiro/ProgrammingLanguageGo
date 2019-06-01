package main

import "fmt"

// reverse reverses a slice of ints in place.
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverse2(ps *[]int) {
	s := *ps
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	list := []int{10, 3, 8, 2}
	fmt.Println(list)
	reverse(list)
	fmt.Println(list)
	reverse2(&list)
	fmt.Println(list)
}

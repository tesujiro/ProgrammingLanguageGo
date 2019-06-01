package main

import "fmt"

// rotate a slice in single pass
func rotate(s *[]int, n int) {
	// new variable
	temp := make([]int, len(*s))
	for i := 0; i < len(*s); i++ {
		temp[i] = (*s)[(i+n)%len(*s)]
	}
	*s = temp
}

func main() {
	list := []int{10, 3, 8, 2, 5}
	fmt.Println(list)
	rotate(&list, 2)
	fmt.Println(list)
}

package main

import (
	"fmt"
)

func main() {
	hs := []int{34, 89, 56, 73, 21, 65}
	fmt.Println("result: ", linear_search(5, hs))
	fmt.Println("result: ", linear_search(21, hs))
}

func linear_search(s int, a []int) int {
	for i, v := range a {
		if v == s {
			return i
		}
	}
	return -1
}

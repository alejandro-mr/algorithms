package main

import (
	"fmt"
)

func main() {
	a := [6]int{31, 41, 59, 26, 41, 58}
	fmt.Println("unsorted: ", a)
	aSorted := insertionSort(a)
	fmt.Println("sorted:   ", aSorted)
}

func insertionSort(a [6]int) [6]int {
	for j := 1; j < len(a); j++ {
		key := a[j]
		i := j - 1
		for i >= 0 && a[i] < key {
			a[i+1] = a[i]
			i--
		}
		a[i+1] = key
	}
	return a
}

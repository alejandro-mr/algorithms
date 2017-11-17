package main

import (
	"fmt"
)

func main() {
	nums := []int{25, 3, 2, 16, 1, 29, 3}
	fmt.Println("Unsorted: ", nums)
	fmt.Println("Sorted: ", selection_sort(nums))
}

func selection_sort(a []int) []int {
	n := len(a)
	for i := 0; i < n-1; i++ {
		min := i
		j := i + 1
		for ; j < n; j++ {
			if a[min] > a[j] {
				min = j
			}
		}
		if min != j {
			current := a[i]
			a[i] = a[min]
			a[min] = current
		}
	}

	return a
}

package main

import (
	"fmt"
)

func main() {
	nums := []int{25, 3, 2, 16, 1, 29, 3}
	fmt.Println("Unsorted: ", nums)
	bubble_sort(nums)
	fmt.Println("Sorted: ", nums)
}

func bubble_sort(a []int) {
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < len(a); i++ {
			if a[i-1] > a[i] {
				//Swap magic
				a[i] = a[i] + a[i-1]
				a[i-1] = a[i] - a[i-1]
				a[i] = a[i] - a[i-1]

				swapped = true
			}
		}
	}
}

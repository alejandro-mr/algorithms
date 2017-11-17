package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{25, 3, 2, 16, 1, 29, 3}
	fmt.Println("unsorted: ", nums)
	merge_sort(nums, 0, len(nums)-1)
	fmt.Println("sorted: ", nums)
}

func merge_sort(a []int, s int, e int) {
	if e > s {
		m := (e + s) / 2
		merge_sort(a, s, m)
		merge_sort(a, m+1, e)
		merge(a, s, m, e)
	}
}

func merge(a []int, s int, m int, e int) {
	l := make([]int, len(a[s:m+1]))
	copy(l, a[s:m+1])
	l = append(l, math.MaxUint32) //Sentinel left

	r := make([]int, len(a[m+1:e+1]))
	copy(r, a[m+1:e+1])
	r = append(r, math.MaxUint32) //Sentinel right

	i, j := 0, 0

	for k := s; k < e+1; k++ {
		if l[i] <= r[j] {
			a[k] = l[i]
			i++
		} else {
			a[k] = r[j]
			j++
		}

	}
}

package main

import "fmt"

func bin_search(intarr *[]int, find_v int, start int, end int) int {
	if end <= start {
		if find_v > (*intarr)[start] {
			return start + 1
		} else {
			return start
		}
	}
	mid := (start + end) >> 1
	if find_v == (*intarr)[mid] {
		return mid + 1
	}
	if find_v > (*intarr)[mid] {
		return bin_search(intarr, find_v, mid+1, end)
	}
	return bin_search(intarr, find_v, start, mid-1)
}

func insertion_sort(intarr *[]int) {
	arrlen := len(*intarr)
	for i := 1; i < arrlen; i++ {
		j := i - 1
		value := (*intarr)[i]
		idx := bin_search(intarr, value, 0, j)
		for j >= idx {
			(*intarr)[j+1] = (*intarr)[j]
			j--
		}
		(*intarr)[j+1] = value
	}
}

func int_slice_sort(intarr *[]int) {
	insertion_sort(intarr)
	for _, d := range *intarr {
		fmt.Printf("%d ", d)
	}
}

func main() {
	exvar := []int{3, 4, 1, 6, 2, 5}
	int_slice_sort(&exvar)
	// exvar1 := []int{1, 2, 3, 4, 5}
	// fmt.Printf("6 index: %d\n", bin_search(&exvar1, 6, 0, len(exvar1)-1))
}

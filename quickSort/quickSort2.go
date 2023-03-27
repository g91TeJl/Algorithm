package main

import "fmt"

func main() {
	arr := []int{5000, 5001, 7, 2, 6, 3}
	fmt.Println(quickSort(arr))
}

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivotIndex := len(arr) / 2
	pivot := arr[pivotIndex]
	var less []int
	var greater []int
	for i, v := range arr {
		if i != pivotIndex {
			if v < pivot {
				less = append(less, v)
			} else {
				greater = append(greater, v)
			}
		}
	}
	sortedLess := quickSort(less)
	sortedGreater := quickSort(greater)
	sorted := append(sortedLess, pivot)
	sorted = append(sorted, sortedGreater...)

	return sorted
}

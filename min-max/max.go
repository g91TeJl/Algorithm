package main

import (
	"fmt"
	"sort"
)

func max_min(n []int) []int {
	/*
		max := n[0]
		min := n[0]
		for _, num := range n {
			if num > max {
				max = num
			}
			if num < min {
				min = num
			}
		}
	*/
	sort.Ints(n)
	return []int{n[0], n[len(n)-1]}

	//return []int{min, max}
}

func main() {
	arr := []int{2334454, 5}
	fmt.Println(max_min(arr))
}

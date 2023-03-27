package main

import "fmt"

func main() {
	slice := []int{5, 3, 6, 2, 10, 1}
	fmt.Println(selectionSort(slice))
}

func selectionSort(slice []int) []int {
	var n = len(slice)
	for i := 0; i < n; i++ {
		minIdx := i
		for j := i; j < n; j++ {
			if slice[j] < slice[minIdx] {

				minIdx = j //смена минимального элемента (обязательна)
			}
			fmt.Println(i)
			fmt.Println(minIdx, "--")
		}
		slice[i], slice[minIdx] = slice[minIdx], slice[i]
		fmt.Println(slice)
	}
	return slice
}

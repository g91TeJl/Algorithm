package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var end []int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				end = append(end, i)
				end = append(end, j)
				return end
			}
		}
	}
	return end
}

func main() {
	nums := []int{3, 2, 3}
	fmt.Println(twoSum(nums, 6))
}

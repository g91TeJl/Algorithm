package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	slice := make([]int, 0)
	var bl bool
	if x < 10 && x > -1 {
		return true
	} else {
		for x > 0 {
			slice = append(slice, x%10)
			x = x / 10
		}
	}
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	for i := 0; i < (len(slice))/2; i++ {
		fmt.Println(slice[i], slice[len(slice)-1-i])
		if slice[i] == slice[len(slice)-1-i] {
			bl = true
		} else {
			return false
		}
	}
	return bl
}

func main() {
	s := 1000021
	fmt.Println(isPalindrome(s))
}

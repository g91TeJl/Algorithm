package main

import (
	"fmt"
	"strings"
)

func LongestConsec(strarr []string, k int) string {
	// if len(strarr) == 0 || k > len(strarr) || k <= 0 {
	// 	return ""
	// }
	// var maxStr string
	// maxLen := 0
	// for i := 0; i <= len(strarr)-1; i++ {
	// 	if k == 1 {
	// 		if len(strarr[i]) > maxLen {
	// 			maxLen = len(strarr[i])
	// 			maxStr = strarr[i]
	// 		}
	// 	} else {
	// 		if i+k-1 > len(strarr)-1 {
	// 			break
	// 		}
	// 		if len(strings.Join(strarr[i:i+k], "")) > maxLen {
	// 			maxLen = len(strings.Join(strarr[i:i+k], ""))
	// 			maxStr = strings.Join(strarr[i:i+k], "")
	// 		}
	// 	}
	// }

	// return maxStr
	var buffer string
	var largest string

	for i := 0; i <= len(strarr)-k; i++ {
		buffer = strings.Join(strarr[i:i+k], "")
		if len(buffer) > len(largest) {
			largest = buffer
		}
	}
	return largest
}

func main() {
	a := LongestConsec([]string{"itvayloxrp", "wkppqsztdkmvcuwvereiupccauycnjutlv", "vweqilsfytihvrzlaodfixoyxvyuyvgpck"}, 2)
	fmt.Println(a)
}

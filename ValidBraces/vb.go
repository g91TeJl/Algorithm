package main

import (
	"fmt"
)

func ValidBraces(str string) bool {
	var Parentheses = map[rune]rune{
		'{': '}',
		'(': ')',
		'[': ']',
	}
	stack := make([]rune, 0)
	for _, v := range str {
		parenthesis, isKey := Parentheses[v]
		if isKey {
			stack = append(stack, parenthesis)
			continue
		}
		lestack := len(stack)
		if lestack == 0 {
			return false
		} else if stack[lestack-1] != v {
			return false
		}
		stack = stack[:lestack-1]
	}
	return len(stack) == 0
}

func main() {
	str := "[({})](]"
	fmt.Println(ValidBraces(str))
}

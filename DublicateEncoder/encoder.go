package main

import (
	"fmt"
	"strings"
)

func DublicateEncode(word string) string {
	word = strings.ToLower(word)
	var end string
	for _, letter := range word {
		fmt.Println(strings.Index(word, string(letter)))
		if strings.Count(word, string(letter)) > 1 {
			end += ")"
		} else {
			end += "("
		}
	}
	return end
}

func main() {
	e := DublicateEncode("(( @")
	fmt.Println(e)
}

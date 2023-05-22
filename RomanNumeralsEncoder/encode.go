package main

/*
Create a function taking a positive integer between 1 and 3999 (both included) as its parameter and returning a string containing the Roman Numeral representation of that integer.
Modern Roman numerals are written by expressing each digit separately starting with the left most digit and skipping any digit with a value of zero.
In Roman numerals 1990 is rendered: 1000=M, 900=CM, 90=XC; resulting in MCMXC. 2008 is written as 2000=MM, 8=VIII; or MMVIII.
1666 uses each Roman symbol in descending order: MDCLXVI.
*/
import (
	"fmt"
	"strings"
)

// func Solution(number int) string {
// 	str := ""
// 	// Если число > 1000 то M * на остаток от деления на 1000 и т.д.
// 	if number >= 1000 {
// 		str += strings.Repeat("M", number/1000)
// 		number = number % 1000
// 	}
// 	if number >= 500 { //900
// 		// if number/10 >= 90 {
// 		// 	str += "СM"
// 		// 	number = (number - 500) % 100
// 		// } else {
// 		if number/100 == 9 {
// 			str += "CM"
// 		} else {

// 			str += "D"

// 			str += strings.Repeat("C", (number-500)/100)
// 		}
// 		number = (number - 500) % 100
// 	}

// 	if number >= 100 { //400
// 		if number/100 == 4 {
// 			str += "CD"
// 		} else {

// 			str += strings.Repeat("C", number/100)
// 		}
// 		number = number % 100
// 	}
// 	if number >= 50 { //90
// 		//fmt.Println(int(float64(number % 10)))
// 		if number/10 == 9 {
// 			str += "XC"
// 		} else {
// 			str += "L"
// 			str += strings.Repeat("X", (number-50)/10)
// 		}
// 		number = number % 10
// 	}

// 	if number >= 10 { //40
// 		if number/10 == 4 {
// 			str += "XL"
// 		} else {

// 			//str += "X"
// 			str += strings.Repeat("X", number/10)
// 		}
// 		number = number % 10
// 	}

// 	if number >= 5 { //4
// 		if number == 9 {
// 			str += "IX"
// 			number = 0
// 		} else {

// 			str += "V"
// 			str += strings.Repeat("I", (number-5)/10)
// 			number = (number - 5) % 10
// 		}

// 	}
// 	if number >= 1 {
// 		if number == 4 {
// 			str += "IV"
// 		} else {

// 			str += strings.Repeat("I", number/1)
// 		}
// 	}

// 	// for number/10 != 0 {
// 	// 	str += fmt.Sprint(number / 10)
// 	// 	number = number / 10
// 	// }
// 	return str
// }

func Solution(n int) string {
	if n <= 0 {
		return ""
	}
	var sb strings.Builder
	for i, v := range []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1} {
		for n >= v {
			n -= v
			sb.WriteString([]string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}[i])
		}
	}
	return sb.String()
}
func main() {
	a := Solution(2420)
	fmt.Println(a)
}

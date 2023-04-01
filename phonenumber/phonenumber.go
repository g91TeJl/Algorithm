package phonenumber

/*
Write a function that accepts an array of 10 integers (between 0 and 9), that returns a string of those numbers in the form of a phone number.
Example

CreatePhoneNumber([10]uint{1,2,3,4,5,6,7,8,9,0})  // returns "(123) 456-7890"

The returned format must be correct in order to complete this challenge.

Don't forget the space after the closing parentheses
*/

import (
	"fmt"
	"strconv"
)

func CreatePhoneNumber(numbers [10]uint) string {
	var phone string
	for _, num := range numbers {
		if num > 9 {
			return "Error"
		}
		phone += strconv.Itoa(int(num))
	}
	/*
		var test string = strings.Trim(strings.Replace(fmt.Sprint(numbers), " ", "", -1), "[]")
		return fmt.Sprintf("(%s) %s-%s", test[0:3], test[3:6], test[6:10])
	*/
	return fmt.Sprintf("(%s) %s-%s", phone[:3], phone[3:6], phone[6:])
}

/*
func main() {
	phone := [10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	s := fmt.Sprint(phone)
	fmt.Println(s)
	fmt.Println(CreatePhoneNumber(phone))
}
*/

package camelcase

/*
Complete the method/function so that it converts dash/underscore delimited words into camel casing. The first word within the output should be capitalized only if the original word was capitalized (known as Upper Camel Case, also often referred to as Pascal case). The next words should be always capitalized.
Examples

"the-stealth-warrior" gets converted to "theStealthWarrior"

"The_Stealth_Warrior" gets converted to "TheStealthWarrior"

"The_Stealth-Warrior" gets converted to "TheStealthWarrior"

*/

import (
	"strings"
)

func ToCamelCase(s string) string {
	// str := s
	// current := ""
	// for i := 0; i < len(str); {
	// 	if str[i] == '_' || str[i] == '-' {
	// 		if string(str[i+1]) != strings.ToUpper(string(str[i+1])) {
	// 			current += strings.ToUpper(string(str[i+1]))
	// 			i += 2
	// 			continue
	// 		}
	// 		i++
	// 		continue
	// 	}
	// 	current += string(str[i])
	// 	i++
	// }
	//return current

	// for idx, val := range str {
	// 	if val == '_' || val == '-' {
	// 		if idx == len(str)-1 {
	// 			break
	// 		}
	// 		if string(str[idx+1]) != strings.ToUpper(string(str[idx+1])) {
	// 			current += strings.ToUpper(string(str[idx+1]))
	// 		}
	// 		continue
	// 	}
	// 	current += string(str[idx])
	// }
	// return current

	if s == "" {
		return ""
	}
	result := strings.Title(strings.Replace(strings.Replace(s, "-", " ", -1), "_", " ", -1))
	result = s[:1] + result[1:]
	result = strings.Replace(result, " ", "", -1)
	return result
}

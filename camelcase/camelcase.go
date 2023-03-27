package camelcase

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

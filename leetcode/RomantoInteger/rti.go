package main

import "fmt"

func romanToInt(s string) int {
	end := 0

	for i := 0; i < len(s); i++ {

		if s[i] == 'M' {
			end += 1000
		}
		if s[i] == 'D' {
			end += 500
		}
		if s[i] == 'C' {

			if i+1 < len(s) && s[i+1] == 'D' {
				end += 400
				i++
				continue
			}

			if i+1 < len(s) && s[i+1] == 'M' {
				end += 900
				i++
				continue
			}

			end += 100
		}
		if s[i] == 'L' {
			end += 50
		}
		if s[i] == 'X' {

			if i+1 < len(s) && s[i+1] == 'L' {
				end += 40
				i++
				continue
			}

			if i+1 < len(s) && s[i+1] == 'C' {
				end += 90
				i++
				continue
			}

			end += 10
		}
		if s[i] == 'V' {
			end += 5
		}
		if s[i] == 'I' {

			if i+1 < len(s) && s[i+1] == 'V' {
				end += 4
				i++
				continue
			}

			if i+1 < len(s) && s[i+1] == 'X' {
				end += 9
				i++
				continue
			}

			end += 1
		}
	}
	return end

}

func main() {
	fmt.Println(romanToInt("MCMXCIV"))

}

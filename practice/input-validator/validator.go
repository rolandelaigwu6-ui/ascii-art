package inputvalidator

import (
	"strings"
	_ "strings"
)

func IsValidInput(input string, font map[byte][5]string) bool {
	if input == "" {
		return false
	}
	for _, line := range strings.Split(input, "\n") {
		if line == " " {
			return false
		}
		for char := 0; char < len(line); char++ {
			if char != len(font) {
				return false
			}
		}
	}
	// for char := 0; char < len(input); char++ {
	// 	if char == 0 {
	// 		return false
	// 	}
	// 	if input[char] == '\n' {
	// 		return true
	// 	}
	// 	_, ok := font[byte(char)] 
	// 	if !ok {
	// 		return true
	// 	}
	// 	if input[char] == '1' {
	// 		return false
	// 	}
	// 	if input[char] == 'a' {
	// 		return false
	// 	} else {
	// 		return true
	// 	}
		
	// }
	return true
}

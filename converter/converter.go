// Question 2: String Art Converter

// Create a function StringToArt that converts a string into ASCII art using only pipe |, underscore _, forward slash /, backslash \, and space characters. The output should look like simple line art.

// Requirements:

// Function signature: func StringToArt(input string) string

// Support digits 0-9 only

// Each digit should be 5 characters wide and 5 lines tall

// Digits should be joined without spacing

// Handle multiple lines in input separated by \n

// Return empty string for invalid input

package converter

import (
	"strings"

	//"honnef.co/go/tools/pattern"
)
var digits = map[byte][5]string{
	'0': {
		" ## ","|   |","|   |","#   |","|___|",
	},
	'1': {
		"  |  ","  |  ","  |  ","  |  ","  |  ",
	},
	'2': {
		" ___ ", "    |", " ___|", "|    ", "|___ ",
	},
}

func StringToArt(input string) string {
	// if input == "" {
	// 	return ""
	// }

	result := ""
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			return ""
		}

		for row := 0; row < 5; row++ {
			for i := range line {
				art, ok := digits[line[i]]
				if !ok {
					return ""
				}
				result += art[row]
			}
			result += "\n"
		}
	}

	return result
}



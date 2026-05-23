package stringart
import (
	"strings"
)

var digits = map[byte][5]string {
	'1': {
		"  |  ",
		"  |  ",
		"  |  ",
		"  |  ",
		"  |  ",
	},
	'2': {
		" ___ ",
		"    |",
		" ___|",
		"|    ",
		"|___ ",
	},
	'3': {
		" ___ ",
		"    |",
		" ___|",
		"    |",
		" ___|",
	},
}
func StringToArt(input string) string {
	if input == "" {
		return ""
	}
	result := ""
	
	for _, lines := range strings.Split(input,"\n") {
		if lines == "" {
			return ""
		}   
		// if char != len(map[byte][]string(digits)) {
		// 	return ""
		// }

		for row := 0; row < 5; row++ {
			for i := range lines {
				art,ok := digits[lines[i]]

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

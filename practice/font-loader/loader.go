package fontloader

func LoadFont(lines []string, chars []byte) map[byte][5]string {
	if len(lines) != len(chars)*5 {
		return nil
	}
	// ch := map[byte][5]string {
	// 	'1': {
	// 		"  |  ",
	// 		"  |  ",
	// 		"  |  ",
	// 		"  |  ",
	// 		"  |  ",
	// 	},
	// 	'2': {
	// 		" ___ ",
	// 		"    |",
	// 		" ___|",
	// 		"|    ",
	// 		"|___ ",
	// 	},
	// }

	font_char := make(map[byte][5]string, len(chars))
	for i, char := range chars {
		startline := i * 5
		var art [5]string
		for row := 0; row < 5; row++ {
			art[row] = lines[startline+row]
		}
		font_char[char] = art
	}
	return font_char
}

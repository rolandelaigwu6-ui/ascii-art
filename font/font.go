package font

func GenerateFont() map[rune][]string {
	result := make(map[rune][]string)

	for ch := rune(32); ch <= rune(126); ch++ {
		result[ch] = makeCharacter(ch)
	}

	return result
}

func makeCharacter(ch rune) []string {
	lines := make([]string, 8)

	if ch == ' ' {
		for i := 0; i < 8; i++ {
			lines[i] = "        "
		}
		return lines
	}

	for row := 0; row < 8; row++ {
		line := ""
		for col := 0; col < 8; col++ {
			if shouldDraw(ch, row, col) {
				line += "*"
			} else if shouldDot(ch, row, col) {
				line += "."
			} else {
				line += " "
			}
		}
		lines[row] = line
	}

	return lines
}

func shouldDraw(ch rune, row, col int) bool {
	if row == 0 || row == 7 || col == 0 || col == 7 {
		return true
	}

	code := int(ch)

	if isVowel(ch) {
		return row == col || row+col == 7
	}

	if isLetter(ch) {
		if row == 3 || col == 1 {
			return true
		}
		return (code+row+col)%5 == 0
	}

	if isDigit(ch) {
		return row == 1 || row == 6 || col == code%6+1
	}

	return (code+row*2+col*3)%7 == 0
}

func shouldDot(ch rune, row, col int) bool {
	return (int(ch)+row+col)%4 == 0
}

func isLetter(ch rune) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z')
}

func isDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

func isVowel(ch rune) bool {
	return ch == 'A' || ch == 'E' || ch == 'I' || ch == 'O' || ch == 'U' ||
		ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u'
}

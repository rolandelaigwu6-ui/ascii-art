package main

import (
	"fmt"
)

// func ValidateBanner(banner map[rune][]string) error {
// 	const (
// 		firstPrintable = rune(32)
// 		lastPrintable  = rune(126)
// 		glyphHeight    = 8
// 	)

// 	if len(banner) != int(lastPrintable-firstPrintable+1) {
// 		return fmt.Errorf("invalid banner: expected 95 characters, got %c", len(banner))
// 	}

// 	for char, glyph := range banner {
// 		if char < firstPrintable || char > lastPrintable {
// 			return fmt.Errorf("invalid banner: character %q is outside printable ASCII range", char)
// 		}

// 		if len(glyph) != glyphHeight {
// 			return fmt.Errorf("invalid banner: character %q has %d rows, expected %d", char, len(glyph), glyphHeight)
// 		}

// 		for rowIndex, row := range glyph {
// 			for _, value := range row {
// 				if value == '\n' || value == '\r' || value == '\t' || value < 32 || value > 126 {
// 					return fmt.Errorf("invalid banner: character %q row %d contains unsupported character %q", char, rowIndex+1, value)
// 				}
// 			}
// 		}
// 	}

// 	for char := firstPrintable; char <= lastPrintable; char++ {
// 		if _, ok := banner[char]; !ok {
// 			return fmt.Errorf("invalid banner: missing character %q", char)
// 		}
// 	}

// 	return nil
// }

// package main

// import (
// 	"fmt"
// )

// func ValidateBanner(banner map[rune][]string)error {
// 	const (
// 		firstprintable = 32
// 		lastprintable = 126
// 		glyphHeigth = 8

// 	)

// 	if banner == nil {
// 		return fmt.Errorf("invalid banner: banner is nil")
// 	}

// 	if len(banner) != int(lastprintable-firstprintable+1) {
// 		return fmt.Errorf("invalid banner: character out-of ASCII Printable range expected 95 unit got %d", len(banner))
// 	}
// 	// for _, gly := range banner {
// 	// 	if len(gly) != glyphHeigth {
// 	// 		return fmt.Errorf("invalid banner: missing %d: expected 8 lines", len(gly))
// 	// 	}
// 	// }
// 	for char := firstprintable; char <= lastprintable; char++ {

// 		gly,ok := banner[rune(char)]
// 		if !ok {
// 			return fmt.Errorf("Invalid banner: character %q missing", char)
// 		}
// 		if len(gly) != glyphHeigth {
// 			return fmt.Errorf("invalid banner: character %q is missing expected 8 line got %d",char, len(gly))
// 		}
// 	}
// 	return nil
// }

func ValidateBanner(banner map[rune][]string) error {
	const (
		firstprintable = 32
		lastprintable  = 126
		glyphHeight    = 8
	)

	if banner == nil {
		return fmt.Errorf("Invalid Banner: Banner == nil")
	}

	if len(banner) != int(lastprintable - firstprintable+1) {
		return fmt.Errorf("Invalid Banner: Got %d expected 95 unit", len(banner))
	}
	for char := firstprintable; char <= lastprintable; char++ {
		gly, ok := banner[rune(char)]
		if !ok { 
			return fmt.Errorf("Invalid Banner: missing line")
		}
		if len(gly) != glyphHeight {
			return fmt.Errorf("nvalid Banner:character %q has %d rows expected %d", rune(char), len(gly), glyphHeight)
		}
	}

	return nil
}

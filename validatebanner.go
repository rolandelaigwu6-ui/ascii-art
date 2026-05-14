package main

import (
	"fmt"
)

func ValidateBanner(banner map[rune][]string) error {
	const (
		firstPrintable = rune(32)
		lastPrintable  = rune(126)
		glyphHeight    = 8
	)

	if len(banner) != int(lastPrintable-firstPrintable+1) {
		return fmt.Errorf("invalid banner: expected 95 characters, got %c", len(banner))
	}

	for char, glyph := range banner {
		if char < firstPrintable || char > lastPrintable {
			return fmt.Errorf("invalid banner: character %q is outside printable ASCII range", char)
		}

		if len(glyph) != glyphHeight {
			return fmt.Errorf("invalid banner: character %q has %d rows, expected %d", char, len(glyph), glyphHeight)
		}

		for rowIndex, row := range glyph {
			for _, value := range row {
				if value == '\n' || value == '\r' || value == '\t' || value < 32 || value > 126 {
					return fmt.Errorf("invalid banner: character %q row %d contains unsupported character %q", char, rowIndex+1, value)
				}
			}
		}
	}

	for char := firstPrintable; char <= lastPrintable; char++ {
		if _, ok := banner[char]; !ok {
			return fmt.Errorf("invalid banner: missing character %q", char)
		}
	}

	return nil
}


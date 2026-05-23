package chargenerator

import "strings"

func GenerateBox(width int, height int) []string {
	if width < 2 || height < 2 {
		return []string{}
	}
	result := []string{}
	for row := 0; row < height; row++ {
		if row == 0 || row == height-1 {
			result = append(result, strings.Repeat("_", width))
		} else {
			middle := "|" + strings.Repeat(" ", width-2) + "|"
			result = append(result, middle)
		}
	}
	return result

}

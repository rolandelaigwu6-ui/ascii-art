package main

import (
	"strings"
	_ "strings"
)

func TrimArtRows(rows []string) []string {
	trimed := make([]string, len(rows))
	for char, row := range rows {
		trimed[char] = strings.TrimRight(row, " ")
	}
	return trimed
	// result := make([]string, 0, len(rows))

	// for _, row := range rows {
	// 	if row != "" {
	// 		result = append(result, row)
	// 	}
	// }
	// return result
}

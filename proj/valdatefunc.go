package main

// import (
// 	"fmt"
// 	"strings"
// )

// func validatebabber(banner map[rune][]string) error {
// 	const (
// 		FirstPrintchar = 32
// 		LastPrintchar = 126
// 		glyphHeight = 8
// 	)

// 	if banner == nil {
// 		return fmt.Errorf("invalid banner: banner not nil")
// 	}
// 	if len(banner) != int(LastPrintchar-FirstPrintchar+1) {
// 		return fmt.Errorf("invalid banner: length of banner must be 95 unit got %d", len(banner))
// 	}
// 	for char := FirstPrintchar; char <= LastPrintchar; char++ {
// 		gly,ok := banner[rune(char)]
// 		if !ok {
// 			return fmt.Errorf("invalid banner: character %q missing " ,char)
// 		}
// 		if len(gly) != glyphHeight {
// 			return fmt.Errorf("invalid banner: character %q has %d row expected %d", char, len(gly), glyphHeight)
// 		}
// 	}
// 	return nil
// }

// func merge(base,priority map[rune][]string)map[rune][]string{
// 	result := make(map[rune][]string, len(base)+len(priority))
// 	for ind, val := range base {
// 		result[ind] = append([]string(nil), val...)
// 	}
// 	return result
// }

// func trimrow(rows []string)[]string {
// 	result := make([]string, 0, len(rows))
// 	for i, row := range rows {
// 		result[i] = strings.TrimRight(row, " ")
// 	}
// 	return result
// }

// func padding()
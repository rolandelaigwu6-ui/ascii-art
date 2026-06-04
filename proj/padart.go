package main

import "strings"

func PadArtRows(rows []string, width int) []string {
	if width <= 0 {
		return rows
	}
	result := make([]string, len(rows))
	for i,row := range rows {
		if len(row) < width {
			spaceneeded := width - len(row)
			padding := strings.Repeat(" ", spaceneeded)
			result[i] = row+padding
		} else {
			result[i] = row
		}
	}
	return result
}
// func PadArtRows(rows []string, targetWidth int) []string {
// 	if targetWidth <= 0 {
// 		return rows
// 	}
// 	result := make([]string, len(rows))
// 	for i, new := range rows {
// 		padding := targetWidth - len(new)
// 		if padding > 0 {

// 			result[i] = new + strings.Repeat(" ", padding)
// 		} else {
// 			result[i] = new
// 		}
// 	}
// 	return result
// }

// func PadArtRows(rows []string, width int) []string {
// 	if width <= 0 {
// 		return rows
// 	}
// 	res := make([]string, len(rows))
// 	for i, v := range rows {
// 		padding := width - len(v)
// 		if padding > 0 {
// 			res[i] = v + strings.Repeat(" ", padding)
// 		} else {
// 			res[i] = v
// 		}
// 	}
// 	return res
// }

// 	if targetWidth <= 0 {
// 		return rows
// 	}
// 	result := make([]string, len(rows))

// 	for i, row := range rows {

// 		if len(row) < 1 {
// 			result[i] = row + strings.Repeat("", spaceneeded)
// 			//result[i] = padd + row
// 		} else {
// 			result[i] = row
// 		}
// 	}
// 	return result
// }

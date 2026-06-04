package main

// func MergeBanners(base map[rune][]string, priority map[rune][]string) map[rune][]string {
// 	banner := make(map[rune][]string)
// 	for char, glyph := range base {
// 		banner[char] = append([]string(nil), glyph...)
// 	}
// 	for char, glyph := range priority {
// 		banner[char] = append([]string(nil), glyph...)
// 	}
// 	return banner
// }

func MergeBanners(base map[rune][]string, priority map[rune][]string) map[rune][]string {
	banner := make(map[rune][]string)
	for r, t := range base {
		banner[r] = append([]string(nil), t...)
	}
	for i, s := range priority {
		banner[i] = append([]string(nil), s...)
	}
	return banner
}

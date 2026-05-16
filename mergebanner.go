package main

import (

)

func MergeBanners(base map[rune][]string, priority map[rune][]string)map[rune][]string {
	banner := make(map[rune][]string, len(base)+len(priority))
	 for char,glyph := range base {
		banner[char] = append([]string(nil),glyph... )
	 }
	 for char,glyph := range priority {
		banner[char] = append([]string(nil), glyph...)
	 }
	 return banner
}
package main

import (
	"fmt"
	//"golang.org/x/text/width"
)

// var ascii = map[rune][]string{
// 	'A': {
// 		"AzzzzA",
// 		"A    A",
// 		"AzzzzA",
// 	},
// }

var ascii = map[rune][]string{
	'A': {
		"AzzzzA",
		"A    A",
		"A    A",
	},
}

func printShap(char rune) {
	shape := ascii[char]

	height := len(shape)
	width := 0

	if len(shape) > 0 {
		width = len(shape[0])
	}

	for row := 0; row < height; row++ {
		fmt.Println(shape[row])
	}
	fmt.Println("height:", height)
	fmt.Println("width:", width)
}

// func printShap(char rune) {
// 	if char < 'A' || char > 'Z' {
// 		fmt.Printf("character must be capital, got %q\n", char)
// 		return
// 	}

// 	shape := ascii[char]
// 	if len(shape) == 0 {
// 		fmt.Printf("missing shape for %q\n", char)
// 		return
// 	}

// 	height := len(shape)
// 	width := 0

// 	if height > 0 {
// 		width = len(shape[0])
// 	}
// 	fmt.Println("height:", height)
// 	fmt.Println("width:", width)

// 	for row := 0; row < height; row++ {
// 		fmt.Println(shape[row])
// 	}
// }

func main() {
	printShap('A')
}

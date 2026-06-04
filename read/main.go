package main

import (
	"ascii-art/engin"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("error with the program [usage go run . sample.txt result.txt]")
		return
	}
	input := os.Args[1]
	output := os.Args[2]

	input1, err := os.ReadFile(input)
	if err != nil {
		fmt.Println("Error cannot read file")
		os.Exit(1)
	}
	s := string(input1)
	token := strings.Fields(s)

	for i, tok := range token {
		switch tok {
		// case "(up)":
		// 	if i > 0 {
		// 		token[i-1] = engin.ToUpper(token[i-1])
		// 	}
		// 	token[i] = ""
		// case "(low)":
		// 	if i > 0 {
		// 		token[i-1] = engin.ToLower(token[i-1])
		// 	}
		// 	token[i] = ""
		case "(~)":
			if i > 0 {
				token[i-1] = engin.TextToAscii(token[i-1])
			}
			token[i] = ""

		case "(?)":
			if i > 0 {
				token[i-1] = engin.Dicryption(token[i-1])
			}
			token[i] = ""
		}
	}
	joihn := strings.Join(token, " ")
	os.WriteFile(output, []byte(joihn), 0644)
}

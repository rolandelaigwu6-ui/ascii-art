package main

import (
	"fmt"
	"os"
	"bufio"
)

func LoadBanner(filename string)(map[rune][]string,error){
	if len(filename) == 0 {
		return nil, fmt.Errorf("Invalid banner: file cannot be empy %d", len(filename))
	}
	file,err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error while establishing connection %v", err)
	}
	defer file.Close()
	if err := Scanner.Err(); err != nil {
		return nil, fmt.Errorf("Something went wrong: ")
	}
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		lines = append(lines, scanner.Text())
	}

	if len(lines) != 855 {
		return nil, fmt.Errorf("Invalid file Expected 855 got %d", len(lines))
	}

	currentRune := rune(32)

}
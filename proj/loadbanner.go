// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func LoadBanner(filename string) (map[rune][]string, error) {
// 	if len(filename) == 0 {
// 		return nil, fmt.Errorf("Invalid banner: file cannot be empy %d", len(filename))
// 	}
// 	file, err := os.Open(filename)
// 	if err != nil {
// 		return nil, fmt.Errorf("error while establishing connection %v", err)
// 	}
// 	defer file.Close()

// 	var lines []string
// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		lines = append(lines, scanner.Text())
// 	}

// 	if len(lines) != 855 {
// 		return nil, fmt.Errorf("Invalid file Expected 855 got %d", len(lines))
// 	}
// 	banner := make(map[rune][]string)
// 	currentRune := rune(32)

// 	for i :=0;i < 8; i++{
// 		if i+8>len(lines){
// 			return nil, fmt.Errorf("inval %q:%d",currentRune, len(lines))
// 		}
// 		chablock :=lines[i+1 : i+9]
// 		banner[currentRune] = chablock
// 		currentRune++
// 	}
// 	if len(banner) != 95 {
// 		return nil, fmt.Errorf("Invalid file Expected 955 got %d", len(lines))
// 	}
// 	return banner,nil

// }
package main

// import (
// 	"fmt"
// 	"os"
// 	"bufio"
// )
// func LoadBanner(filename string) (map[rune][]string, error) {
// 	if len(filename) == 0 {
// 		return nil, fmt.Errorf("file cannot be empty")
// 	}
// 	file,err := os.Open(filename)
// 	if err != nil {
// 		return nil, fmt.Errorf("error reading file %v", err)
// 	}
// 	defer file.Close()
// 	scanner := bufio.NewScanner(file)
// 	var lines []string
	
// 	for scanner.Scan(){
// 		lines = append(lines, scanner.Text())
// 	}
// 	if len(lines) != 855 {
// 		return nil, fmt.Errorf("error expected 855 got %v", len(lines))
// 	}
// 	// if len(lines)%9 == 0 {
// 	// 	return nil, fmt.Errorf("Invalid: size/hight %v", len(lines))
// 	// }
// 	banner := make(map[rune][]string)
// 	currentrune := rune(32)

// 	for i := 0; i < len(lines); i+=9 {
// 		if i+8 > len(lines) {
// 			return nil, fmt.Errorf("unsupported rune %v at index %d", currentrune, i)
// 		}
// 		charblock := lines[i+1 : i+9]
// 		banner[currentrune] = charblock
// 		currentrune++
// 	}
// 	if len(banner) != 95 {
// 		return nil, fmt.Errorf("corrup banner file %v", len(lines))
// 	}
// 	return banner, nil
// }
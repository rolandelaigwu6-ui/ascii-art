package main

import (
	"fmt"
	"strconv"
)

func HexToDec(input []string) ([]int64, error) {
	result := []int64{}

	for _, inp := range input {
		inp1, err := strconv.ParseInt(inp, 16, 64)
		if err != nil {
			return nil, err

		}
		result = append(result, inp1)
		

	}
	return []int64(result), nil
	// decimal, err := strconv.ParseInt(input, 16, 64)
	// if err != nil {
	// 	return 0, err
	// }

	// return decimal, nil
}

// func DecToHex(input int64) string {
// 	return strconv.FormatInt(input, 16)
// }

func main() {
	deci, err := HexToDec([]string{"f11FF"})
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	//hex := DecToHex(deci)

	fmt.Println(deci)
	//fmt.Println(hex)
}

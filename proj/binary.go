package main

import "strings"

func Binary(text string) string {
	var result []string

	for _, char := range []byte(text) {
		result = append(result, ByteToBinary(char))
	}
	return strings.Join(result, " ")
}
 
func ByteToBinary(value byte) string {
	bits := make([]byte, 8)

	for i := 7; i >= 0; i-- {
		if value%2 == 1 {
			bits[i] = '1'
		} else {
			bits[i] = '0'
		}

		value = value / 2
	}

	return string(bits)
}

package main

import (

)

// func TextFromBinary(text byte)string {
// 	var result []string

// 	for _, char := range text{
// 		result = append(result, char)
// 	}
// 	return result[]

// }

func BinaryByteToNumber(bits string)byte {
	var value byte

	for _, bit := range bits {
		value = value*2

		if bit == '1'{
			value = value+1
		}
	}
	return value
}
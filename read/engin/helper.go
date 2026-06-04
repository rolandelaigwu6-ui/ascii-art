package engin

import "strings"

func ToUpper(input string) string {
	result := []byte(input)
	output := []string{}

	for i := range result {
		if result[i] >= 'a' && result[i] <= 'z' {
			result[i] = result[i] - 32
		}
		output = append(output, string(result))

	}
	return string(result)
}

func ToLower(input string) string {
	result := []byte(input)
	//output := []string{}

	for i := range result {
		if result[i] >= 'A' && result[i] <= 'Z' {
			result[i] = result[i] + 32
		}

	}
	return string(result)
}


func TextToAscii(word string) string {
	result := []string{}
	//word1 := []byte(word)
	for _, words  := range []byte(word) {
		result = append(result, BitToBinary(words))
	}
	return strings.Join(result, " ")
}

func BitToBinary(input byte) string {
	bits := make([]byte, 8)

	for i := 7; i >= 0; i-- {
		if input%2 == 1 {
			bits[i] = '1'
		} else {
			bits[i] = '0'
		}
		input = input/2
	}
	return string(bits)
}


func Dicryption(input string) string {
	total := 0 
	for _, inp := range input {
		total = total * 2

		if inp == '1' {
			total = total + 1
		}
	}
	return string(byte(total))
}
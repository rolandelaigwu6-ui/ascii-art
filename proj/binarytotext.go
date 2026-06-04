package main

func TextFromBinary(value string) string {
	total := 0
	for _, bit := range value {
		total = total * 2

		if bit == '1' {
			total = total + 1
		}
	}
	return string(byte(total))
}

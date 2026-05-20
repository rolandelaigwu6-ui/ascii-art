package main

func StackTwo(top []string, bottom []string) []string {
	result := make([]string, 0, len(top)+len(bottom))
	copy(result, top)
	copy(result[len(top):], bottom)
	// result = append(result, top...)
	// result = append(result, bottom...)
	return result
}

func StackAll(blocks []string) []string {
	result := []string{}
	for _, char := range blocks {
		result = StackTwo(result, char)
	}
	return result
}

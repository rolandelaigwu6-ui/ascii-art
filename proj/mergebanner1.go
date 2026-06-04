package main

func MergeBanner(base map[rune][]string, priority map[rune][]string) map[rune][]string {
	banner := make(map[rune][]string)
	for r := range base {
		banner[r] = append([]string(nil), string(r))
	}
	for i := range priority {
		banner[i] = append([]string(nil), string(i))
	}
	return banner
}

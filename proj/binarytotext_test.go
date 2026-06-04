package main

import (
	"testing"
)

func TestTextFromBinary_BinaryToText(t *testing.T) {
	test := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "biary A",
			in:   "01000001",
			want: "A",
		},
		{
			name: "bainary B",
			in:   "01000010",
			want: "B",
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			got := TextFromBinary(tt.in)
			if got != tt.want {
				t.Fatalf("TextFromBinary(%q) = %q  want %q",tt.in, got, tt.want)
			}
		})
	}
}

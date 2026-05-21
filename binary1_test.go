package main

import "testing"

func TestBinary_TextToBinary(t *testing.T) {
	tests := []struct {
		name  string
		in    string
		wants string
	}{
		{
			name:  "binary A",
			in:    "A",
			wants: "01000001",
		},
		{
			name:  "binary B",
			in:    "AB",
			wants: "01000001 01000010",
		},
		{
			name: "binary C",
			in: "",
			wants: "",

		},
		{
			name: "binary d",
			in: "A B",
			wants: "01000001 00100000 01000010",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Binary(tt.in)
			if got != tt.wants {
				t.Fatalf("Binary(%q) = %q, want %q", tt.in, got, tt.wants)
			}
		})
	}

}

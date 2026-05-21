package main

import "testing"

func TestByteToBinary_ASCIICharacters(t *testing.T) {
	tests := []struct {
		name string
		in   byte
		want string
	}{
		{name: "capital A", in: 'A', want: "01000001"},
		{name: "capital B", in: 'B', want: "01000010"},
		{name: "space", in: ' ', want: "00100000"},
		{name: "zero byte", in: 0, want: "00000000"},
		{name: "max byte", in: 255, want: "11111111"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ByteToBinary(tt.in)
			if got != tt.want {
				t.Fatalf("ByteToBinary(%d) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}

func TestByteToBinary_AlwaysEightBits(t *testing.T) {
	for i := 0; i <= 255; i++ {
		got := ByteToBinary(byte(i))
		if len(got) != 8 {
			t.Fatalf("ByteToBinary(%d) returned %q with length %d, want length 8", i, got, len(got))
		}
	}
}

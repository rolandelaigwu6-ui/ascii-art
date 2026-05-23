package inputvalidator

import "testing"

var testFont = map[byte][5]string{
	'1': {"  |  ", "  |  ", "  |  ", "  |  ", "  |  "},
	'2': {" ___ ", "    |", " ___|", "|    ", "|___ "},
}

func TestIsValidInput(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{name: "single line", input: "12", want: true},
		{name: "multiple lines", input: "1\n2", want: true},
		{name: "invalid character", input: "1a", want: false},
		{name: "empty input", input: "", want: false},
		{name: "empty second line", input: "1\n", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsValidInput(tt.input, testFont)
			if got != tt.want {
				t.Errorf("IsValidInput(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}


package fontbuilder

import "testing"

func TestBuildFontContainsDigits(t *testing.T) {
	font := BuildFont()

	for _, digit := range []byte{'0', '1', '2'} {
		if _, ok := font[digit]; !ok {
			t.Errorf("BuildFont missing digit %q", digit)
		}
	}
}

func TestBuildFontRowsHaveCorrectWidth(t *testing.T) {
	font := BuildFont()

	for digit, art := range font {
		for rowIndex, row := range art {
			if len(row) != 5 {
				t.Errorf("digit %q row %d has width %d, want 5", digit, rowIndex, len(row))
			}
		}
	}
}

func TestBuildFontDigitTwo(t *testing.T) {
	font := BuildFont()
	want := [5]string{
		" ___ ",
		"    |",
		" ___|",
		"|    ",
		"|___ ",
	}

	if font['2'] != want {
		t.Errorf("font['2'] = %q, want %q", font['2'], want)
	}
}


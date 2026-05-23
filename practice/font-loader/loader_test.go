package fontloader

import "testing"

func TestLoadFont(t *testing.T) {
	lines := []string{
		"  |  ",
		"  |  ",
		"  |  ",
		"  |  ",
		"  |  ",
		" ___ ",
		"    |",
		" ___|",
		"|    ",
		"|___ ",
	}
	chars := []byte{'1', '2'}

	font := LoadFont(lines, chars)
	if font == nil {
		t.Fatal("LoadFont returned nil")
	}

	if font['1'][0] != "  |  " {
		t.Errorf("font['1'][0] = %q, want %q", font['1'][0], "  |  ")
	}

	if font['2'][4] != "|___ " {
		t.Errorf("font['2'][4] = %q, want %q", font['2'][4], "|___ ")
	}
}

func TestLoadFontRejectsWrongLineCount(t *testing.T) {
	lines := []string{"one", "two", "three"}
	chars := []byte{'1'}

	if font := LoadFont(lines, chars); font != nil {
		t.Errorf("LoadFont returned %v, want nil", font)
	}
}


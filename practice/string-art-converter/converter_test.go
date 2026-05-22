package stringart

import "testing"

func TestStringToArtSingleDigitOne(t *testing.T) {
	got := StringToArt("1")
	want := "  |  \n  |  \n  |  \n  |  \n  |  \n"

	if got != want {
		t.Errorf("StringToArt(%q) = %q, want %q", "1", got, want)
	}
}

func TestStringToArtSingleDigitThree(t *testing.T) {
	got := StringToArt("3")
	want := " ___ \n    |\n ___|\n    |\n ___|\n"

	if got != want {
		t.Errorf("StringToArt(%q) = %q, want %q", "3", got, want)
	}
}

func TestStringToArtSideBySide(t *testing.T) {
	got := StringToArt("13")
	want := "  |   ___ \n  |      |\n  |   ___|\n  |      |\n  |   ___|\n"

	if got != want {
		t.Errorf("StringToArt(%q) = %q, want %q", "13", got, want)
	}
}

func TestStringToArtInvalidInput(t *testing.T) {
	got := StringToArt("1a")

	if got != "" {
		t.Errorf("StringToArt(%q) = %q, want empty string", "1a", got)
	}
}

func TestStringToArtEmptyInput(t *testing.T) {
	got := StringToArt("")

	if got != "" {
		t.Errorf("StringToArt(%q) = %q, want empty string", "", got)
	}
}

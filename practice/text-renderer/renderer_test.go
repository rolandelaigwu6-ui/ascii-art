package textrenderer

import "testing"

var font = map[byte][5]string{
	'1': {"  |  ", "  |  ", "  |  ", "  |  ", "  |  "},
	'2': {" ___ ", "    |", " ___|", "|    ", "|___ "},
}

func TestRenderSideBySide(t *testing.T) {
	got := Render("12", font)
	want := "  |   ___ \n  |      |\n  |   ___|\n  |  |    \n  |  |___ \n"

	if got != want {
		t.Errorf("Render(%q) = %q, want %q", "12", got, want)
	}
}

func TestRenderMultipleLines(t *testing.T) {
	got := Render("1\n2", font)
	want := "  |  \n  |  \n  |  \n  |  \n  |  \n" +
		" ___ \n    |\n ___|\n|    \n|___ \n"

	if got != want {
		t.Errorf("Render(%q) = %q, want %q", "1\\n2", got, want)
	}
}

func TestRenderRejectsUnknownCharacters(t *testing.T) {
	if got := Render("1a", font); got != "" {
		t.Errorf("Render(%q) = %q, want empty string", "1a", got)
	}
}


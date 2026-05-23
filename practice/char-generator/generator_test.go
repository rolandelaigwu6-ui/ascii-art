package chargenerator

import "testing"

func TestGenerateBox(t *testing.T) {
	got := GenerateBox(5, 4)
	want := []string{
		"_____",
		"|   |",
		"|   |",
		"_____",
	}

	if len(got) != len(want) {
		t.Fatalf("GenerateBox returned %d rows, want %d", len(got), len(want))
	}

	for i := range want {
		if got[i] != want[i] {
			t.Errorf("row %d = %q, want %q", i, got[i], want[i])
		}
	}
}

func TestGenerateBoxRejectsSmallSizes(t *testing.T) {
	if got := GenerateBox(1, 5); len(got) != 0 {
		t.Errorf("GenerateBox(1, 5) returned %v, want empty slice", got)
	}

	if got := GenerateBox(5, 1); len(got) != 0 {
		t.Errorf("GenerateBox(5, 1) returned %v, want empty slice", got)
	}
}


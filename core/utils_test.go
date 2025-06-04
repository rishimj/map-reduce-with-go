package core

import "testing"

func TestSplitWordsPunctuation(t *testing.T) {
	input := "Hello, Go!"
	expected := []string{"hello", "go"}
	got := SplitWords(input)
	if len(got) != len(expected) {
		t.Fatalf("expected %d words, got %d", len(expected), len(got))
	}
	for i, w := range expected {
		if got[i] != w {
			t.Errorf("at index %d: expected %q, got %q", i, w, got[i])
		}
	}
}

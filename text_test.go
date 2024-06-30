package ebitenpkg

import (
	"testing"
)

func TestTextSize(t *testing.T) {
	size := 50.0
	text := NewText("Helloooo", size)

	t.Log(text.Bound())

	if text.Size() != size {
		t.Fatalf("test size should be %f, but get %f", size, text.Size())
	}
}

package util

import "testing"

func TestVectorAdd(t *testing.T) {
	want := Vector2{3, 4}
	a := Vector2{1, 1}
	b := Vector2{2, 3}

	if got := a.Add(b); got != want {
		t.Errorf("%q.Add(%q) = %q, want %q", a, b, got, want)
	}
}

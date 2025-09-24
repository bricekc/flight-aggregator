package utils

import (
	"testing"
)

func TestTernary(t *testing.T) {
	// given
	cond := true
	a := "It's true !"
	b := "It's false !"
	// when
	result := Ternary(cond, a, b)

	// then
	expected := a
	if result != expected {
		t.Errorf("expected %q but got %q", expected, result)
	}
}

package semigroup

import (
	"testing"
)

func TestFromConcat(t *testing.T) {
	got := FromConcat(func(c1 string, c2 string) string {
		return c1 + c2
	})

	if got.Concat("a", "b") != "ab" {
		t.Errorf("Expected ab, got %s", got.Concat("a", "b"))
	}
}

func TestReverse(t *testing.T) {
	strSg := FromConcat(func(c1 string, c2 string) string {
		return c1 + c2
	})

	got := Reverse(strSg).Concat("aa", "bbb")

	if got != "bbbaa" {
		t.Errorf("Expected bbbaa, got %s", got)
	}
}

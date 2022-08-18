package monoid

import (
	"testing"

	"github.com/medmouine/gomad/semigroup"
)

func TestFromConcat(t *testing.T) {
	concat := func(a, b string) string {
		return a + b
	}
	m := FromConcat(concat, "a")
	if m.Concat("a", "b") != "b" {
		t.Errorf("Expected %s, got %s", m.Concat("a", "b"), "b")
	}

	if m.Concat("b", "a") != "b" {
		t.Errorf("Expected %s, got %s", m.Concat("b", "a"), "b")
	}

	if m.Concat("b", "c") != "bc" {
		t.Errorf("Expected %s, got %s", m.Concat("b", "c"), "bc")
	}
}

func TestFrom(t *testing.T) {
	concat := func(a, b string) string {
		return a + b
	}
	sg := semigroup.FromConcat(concat)

	m := From(sg, "a")

	if m.Concat("a", "b") != "b" {
		t.Errorf("Expected %s, got %s", m.Concat("a", "b"), "b")
	}

	if m.Concat("b", "a") != "b" {
		t.Errorf("Expected %s, got %s", m.Concat("b", "a"), "b")
	}

	if m.Concat("b", "c") != "bc" {
		t.Errorf("Expected %s, got %s", m.Concat("b", "c"), "bc")
	}
}

func TestMonoid_Fold(t *testing.T) {
	concat := func(a, b string) string {
		return a + b
	}
	m := FromConcat(concat, "a").Fold()

	if m([]string{"b", "c"}) != "abc" {
		t.Errorf("Expected %s, got %s", m([]string{"b", "c"}), "abc")
	}
}

func TestMonoid_FoldLeft(t *testing.T) {
	concat := func(a, b string) string {
		return a + b
	}
	m := FromConcat(concat, "a").FoldLeft()

	if m([]string{"b", "c"}) != "cba" {
		t.Errorf("Expected %s, got %s", m([]string{"b", "c"}), "cba")
	}
}

func TestMonoid_FoldMap(t *testing.T) {
	concat := func(a, b string) string {
		return a + b
	}
	m := FromConcat(concat, "a").FoldMap(func(a string) string {
		return a + "|"
	})

	if m([]string{"b", "c"}) != "a|b|c|" {
		t.Errorf("Expected %s, got %s", m([]string{"b", "c"}), "a|b|c|")
	}
}

func TestMonoid_FoldMapLeft(t *testing.T) {
	concat := func(a, b string) string {
		return a + b
	}
	m := FromConcat(concat, "a").FoldMapLeft(func(a string) string {
		return a + "|"
	})

	if m([]string{"b", "c"}) != "c|b|a|" {
		t.Errorf("Expected %s, got %s", m([]string{"b", "c"}), "c|b|a|")
	}
}

func TestMonoid_Concat(t *testing.T) {
	concat := func(a, b string) string {
		return a + b
	}
	m := FromConcat(concat, "a")

	if m.Concat("a", "b") != "b" {
		t.Errorf("Expected %s, got %s", m.Concat("a", "b"), "b")
	}

	if m.Concat("b", "a") != "b" {
		t.Errorf("Expected %s, got %s", m.Concat("b", "a"), "b")
	}

	if m.Concat("b", "c") != "bc" {
		t.Errorf("Expected %s, got %s", m.Concat("b", "c"), "bc")
	}
}

func TestMonoid_Empty(t *testing.T) {
	concat := func(a, b string) string {
		return a + b
	}
	m := FromConcat(concat, "a")

	if m.Empty() != "a" {
		t.Errorf("Expected %s, got %s", m.Empty(), "a")
	}
}

func TestFold(t *testing.T) {
	concat := func(a, b string) string {
		return a + b
	}
	m := Fold(FromConcat(concat, "a"))

	if m([]string{"b", "c"}) != "abc" {
		t.Errorf("Expected %s, got %s", m([]string{"b", "c"}), "abc")
	}

	sum := func(a, b int) int {
		return a + b
	}

	m2 := Fold(FromConcat(sum, 0))

	if m2([]int{1, 2, 3}) != 6 {
		t.Errorf("Expected %d, got %d", m2([]int{1, 2, 3}), 6)
	}
}

func TestFoldLeft(t *testing.T) {
	concat := func(a, b string) string {
		return a + b
	}
	m := FoldLeft(FromConcat(concat, "a"))

	if m([]string{"b", "c"}) != "cba" {
		t.Errorf("Expected %s, got %s", m([]string{"b", "c"}), "cba")
	}
}

func TestFoldMap(t *testing.T) {
	concat := func(a, b string) string {
		return a + b
	}
	m := FoldMap(FromConcat(concat, "a"), func(a string) string {
		return a + "|"
	})

	if m([]string{"b", "c"}) != "a|b|c|" {
		t.Errorf("Expected %s, got %s", m([]string{"b", "c"}), "a|b|c|")
	}
}

func TestFoldMapLeft(t *testing.T) {
	concat := func(a, b string) string {
		return a + b
	}
	m := FoldMapLeft(FromConcat(concat, "a"), func(a string) string {
		return a + "|"
	})

	if m([]string{"b", "c"}) != "c|b|a|" {
		t.Errorf("Expected %s, got %s", m([]string{"b", "c"}), "c|b|a|")
	}
}

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

func TestConcat(t *testing.T) {
	strSg := FromConcat(func(c1 string, c2 string) string {
		return c1 + c2
	})

	got := strSg.Concat("aa", "bbb")

	if got != "aabbb" {
		t.Errorf("Expected aabbb, got %s", got)
	}
}

func TestFold(t *testing.T) {
	strSg := FromConcat(func(c1 string, c2 string) string {
		return c1 + c2
	})

	got := Fold(strSg, "a")

	if got([]string{"b", "c", "d"}) != "abcd" {
		t.Errorf("Expected abcd, got %s", got([]string{"b", "c", "d"}))
	}
}

func TestFoldMap(t *testing.T) {
	strSg := FromConcat(func(c1 string, c2 string) string {
		return c1 + c2
	})

	got := FoldMap(strSg, func(s string) string {
		return s + "|"
	}, "a")

	if got([]string{"b", "c", "d"}) != "a|b|c|d|" {
		t.Errorf("Expected bacaad, got %s", got([]string{"b", "c", "d"}))
	}
}

func TestFoldF(t *testing.T) {
	strSg := FromConcat(func(c1 string, c2 string) string {
		return c1 + c2
	})

	got := FoldF(strSg)

	if got([]string{"b", "c", "d"}) != "bcd" {
		t.Errorf("Expected abcd, got %s", got([]string{"b", "c", "d"}))
	}
}

func TestFoldMapF(t *testing.T) {
	strSg := FromConcat(func(c1 string, c2 string) string {
		return c1 + c2
	})

	got := FoldMapF(strSg, func(s string) string {
		return s + "|"
	})

	if got([]string{"b", "c", "d"}) != "b|c|d|" {
		t.Errorf("Expected bacaad, got %s", got([]string{"b", "c", "d"}))
	}
}

func TestFoldLeft(t *testing.T) {
	strSg := FromConcat(func(c1 string, c2 string) string {
		return c1 + c2
	})

	got := FoldLeft(strSg, "a")

	if got([]string{"b", "c", "d"}) != "dcba" {
		t.Errorf("Expected dcba, got %s", got([]string{"b", "c", "d"}))
	}
}

func TestFoldMapLeft(t *testing.T) {
	strSg := FromConcat(func(c1 string, c2 string) string {
		return c1 + c2
	})

	got := FoldMapLeft(strSg, func(s string) string {
		return s + "|"
	}, "a")

	if got([]string{"b", "c", "d"}) != "d|c|b|a|" {
		t.Errorf("Expected d|c|b|a|, got %s", got([]string{"b", "c", "d"}))
	}
}

func TestFoldLeftF(t *testing.T) {
	strSg := FromConcat(func(c1 string, c2 string) string {
		return c1 + c2
	})

	got := FoldLeftF(strSg)

	if got([]string{"b", "c", "d"}) != "dcb" {
		t.Errorf("Expected dcb, got %s", got([]string{"b", "c", "d"}))
	}
}

func TestFoldMapLeftF(t *testing.T) {
	strSg := FromConcat(func(c1 string, c2 string) string {
		return c1 + c2
	})

	got := FoldMapLeftF(strSg, func(s string) string {
		return s + "|"
	})

	if got([]string{"b", "c", "d"}) != "d|c|b|" {
		t.Errorf("Expected d|c|b|, got %s", got([]string{"b", "c", "d"}))
	}
}

func TestSemigroup_Fold(t *testing.T) {
	strSg := FromConcat(func(c1 string, c2 string) string {
		return c1 + c2
	})

	got := strSg.Fold("a")

	if got([]string{"b", "c", "d"}) != "abcd" {
		t.Errorf("Expected abcd, got %s", got([]string{"b", "c", "d"}))
	}
}

func TestSemigroup_FoldMap(t *testing.T) {
	strSg := FromConcat(func(c1 string, c2 string) string {
		return c1 + c2
	})

	got := strSg.FoldMap(func(s string) string {
		return s + "|"
	}, "a")

	if got([]string{"b", "c", "d"}) != "a|b|c|d|" {
		t.Errorf("Expected a|b|c|d|, got %s", got([]string{"b", "c", "d"}))
	}
}

func TestSemigroup_FoldF(t *testing.T) {
	strSg := FromConcat(func(c1 string, c2 string) string {
		return c1 + c2
	})

	got := strSg.FoldF()

	if got([]string{"b", "c", "d"}) != "bcd" {
		t.Errorf("Expected abcd, got %s", got([]string{"b", "c", "d"}))
	}
}

func TestSemigroup_FoldLeft(t *testing.T) {
	strSg := FromConcat(func(c1 string, c2 string) string {
		return c1 + c2
	})

	got := strSg.FoldLeft("a")

	if got([]string{"b", "c", "d"}) != "dcba" {
		t.Errorf("Expected dcba, got %s", got([]string{"b", "c", "d"}))
	}
}

func TestSemigroup_FoldLeftF(t *testing.T) {
	strSg := FromConcat(func(c1 string, c2 string) string {
		return c1 + c2
	})

	got := strSg.FoldLeftF()

	if got([]string{"b", "c", "d"}) != "dcb" {
		t.Errorf("Expected dcb, got %s", got([]string{"b", "c", "d"}))
	}
}

func TestSemigroup_FoldMapF(t *testing.T) {
	strSg := FromConcat(func(c1 string, c2 string) string {
		return c1 + c2
	})

	got := strSg.FoldMapF(func(s string) string {
		return s + "|"
	})

	if got([]string{"b", "c", "d"}) != "b|c|d|" {
		t.Errorf("Expected bacaad, got %s", got([]string{"b", "c", "d"}))
	}
}

func TestSemigroup_FoldMapLeft(t *testing.T) {
	strSg := FromConcat(func(c1 string, c2 string) string {
		return c1 + c2
	})

	got := strSg.FoldMapLeft(func(s string) string {
		return s + "|"
	}, "a")

	if got([]string{"b", "c", "d"}) != "d|c|b|a|" {
		t.Errorf("Expected d|c|b|a|, got %s", got([]string{"b", "c", "d"}))
	}
}

func TestSemigroup_FoldMapLeftF(t *testing.T) {
	strSg := FromConcat(func(c1 string, c2 string) string {
		return c1 + c2
	})

	got := strSg.FoldMapLeftF(func(s string) string {
		return s + "|"
	})

	if got([]string{"b", "c", "d"}) != "d|c|b|" {
		t.Errorf("Expected d|c|b|, got %s", got([]string{"b", "c", "d"}))
	}
}

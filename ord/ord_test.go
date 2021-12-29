package ord

import (
	"testing"
)

func TestFromCompare(t *testing.T) {
	intCompare := func(a int, b int) int {
		return a - b
	}

	got := FromCompare(intCompare)

	if got.Compare(1, 2) != -1 {
		t.Errorf("got %v, want %v", got.Compare(1, 2), -1)
	}

	if got.Compare(2, 1) != 1 {
		t.Errorf("got %v, want %v", got.Compare(2, 1), 1)
	}

	if got.Compare(1, 1) != 0 {
		t.Errorf("got %v, want %v", got.Compare(1, 1), 0)
	}
}

func TestMax(t *testing.T) {
	intCompare := func(a int, b int) int {
		return a - b
	}

	got := Max(FromCompare(intCompare))

	if got(1, 2) != 2 {
		t.Errorf("got %v, want %v", got(1, 2), 2)
	}

	if got(2, 1) != 2 {
		t.Errorf("got %v, want %v", got(2, 1), 2)
	}
}

func TestBetween(t *testing.T) {
	intCompare := func(a int, b int) int {
		return a - b
	}

	got := Between(FromCompare(intCompare))

	if got(1, 2)(3) {
		t.Errorf("got %v, want %v", got(1, 2)(3), false)
	}
}

func TestEquals(t *testing.T) {
	intCompare := func(a int, b int) int {
		return a - b
	}

	got := Equals(FromCompare(intCompare))

	if got.Equals(1, 2) {
		t.Errorf("got %v, want %v", got.Equals(1, 2), false)
	}

	if !got.Equals(1, 1) {
		t.Errorf("got %v, want %v", got.Equals(1, 1), true)
	}
}

func TestGeq(t *testing.T) {
	intCompare := func(a int, b int) int {
		return a - b
	}

	got := Geq(FromCompare(intCompare))

	if got(1, 2) {
		t.Errorf("got %v, want %v", got(1, 2), false)
	}

	if !got(2, 1) {
		t.Errorf("got %v, want %v", got(2, 1), true)
	}

	if !got(1, 1) {
		t.Errorf("got %v, want %v", got(1, 1), true)
	}
}

func TestGt(t *testing.T) {
	intCompare := func(a int, b int) int {
		return a - b
	}

	got := Gt(FromCompare(intCompare))

	if got(1, 2) {
		t.Errorf("got %v, want %v", got(1, 2), false)
	}

	if !got(2, 1) {
		t.Errorf("got %v, want %v", got(2, 1), true)
	}

	if got(1, 1) {
		t.Errorf("got %v, want %v", got(1, 1), false)
	}
}

func TestLeq(t *testing.T) {
	intCompare := func(a int, b int) int {
		return a - b
	}

	got := Leq(FromCompare(intCompare))

	if !got(1, 2) {
		t.Errorf("got %v, want %v", got(1, 2), true)
	}

	if got(2, 1) {
		t.Errorf("got %v, want %v", got(2, 1), false)
	}

	if !got(1, 1) {
		t.Errorf("got %v, want %v", got(1, 1), true)
	}
}

func TestLt(t *testing.T) {
	intCompare := func(a int, b int) int {
		return a - b
	}

	got := Lt(FromCompare(intCompare))

	if !got(1, 2) {
		t.Errorf("got %v, want %v", got(1, 2), true)
	}

	if got(2, 1) {
		t.Errorf("got %v, want %v", got(2, 1), false)
	}

	if got(1, 1) {
		t.Errorf("got %v, want %v", got(1, 1), false)
	}
}

func TestMin(t *testing.T) {
	intCompare := func(a int, b int) int {
		return a - b
	}

	got := Min(FromCompare(intCompare))

	if got(1, 2) != 1 {
		t.Errorf("got %v, want %v", got(1, 2), 1)
	}

	if got(2, 1) != 1 {
		t.Errorf("got %v, want %v", got(2, 1), 1)
	}
}

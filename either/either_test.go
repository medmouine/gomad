package either

import (
	"go/types"
	"reflect"
	"testing"
)

func TestFromPredicate(t *testing.T) {
	gotRight := FromPredicate(func(i int) bool { return i%2 == 0 }, 2, "not even")

	if gotRight.IsLeft() {
		t.Errorf("gotRight isLeft %v, want %v", gotRight.IsLeft(), false)
	}
	if !gotRight.IsRight() {
		t.Errorf("gotRight isRight %v, want %v", gotRight.IsRight(), true)
	}
	if !reflect.DeepEqual(gotRight.Right(), 2) {
		t.Errorf("gotRight %v, want %v", gotRight.Right(), 2)
	}

	gotLeft := FromPredicate(func(i int) bool { return i%2 == 0 }, 3, "not even")

	if !gotLeft.IsLeft() {
		t.Errorf("gotLeft isLeft %v, want %v", gotLeft.IsLeft(), true)
	}
	if gotLeft.IsRight() {
		t.Errorf("gotLeft isRight %v, want %v", gotLeft.IsRight(), false)
	}
	if !reflect.DeepEqual(gotLeft.Left(), "not even") {
		t.Errorf("gotLeft %v, want %v", gotLeft.Left(), "not even")
	}
}

func TestEither_MapLeft(t *testing.T) {
	left := Left("foo")

	got := left.MapLeft(func(t string) string {
		return t + "bar"
	})

	if !reflect.DeepEqual(got.Left(), "foobar") {
		t.Errorf("MapLeft() = %v, want %v", got.Left(), "foobar")
	}

	right := Right("foo")

	var called = false
	right.MapLeft(func(t types.Nil) types.Nil {
		called = true
		return t
	})

	if !reflect.DeepEqual(right.Right(), "foo") {
		t.Errorf("MapLeft() = %v, want %v", right.Right(), "foo")
	}

	if !reflect.DeepEqual(called, false) {
		t.Errorf("MapLeft() = %v, want %v", called, false)
	}
}

func TestEither_MapRight(t *testing.T) {
	left := Left("foo")

	var called = false
	left.MapRight(func(t types.Nil) types.Nil {
		called = true
		return t
	})

	if !reflect.DeepEqual(left.Left(), "foo") {
		t.Errorf("MapRight() = %v, want %v", left.Left(), "foo")
	}
	if !reflect.DeepEqual(called, false) {
		t.Errorf("MapRight() = %v, want %v", called, false)
	}

	right := Right("foo")

	got := right.MapRight(func(t string) string {
		return t + "bar"
	})

	if !reflect.DeepEqual(got.Right(), "foobar") {
		t.Errorf("MapRight() = %v, want %v", got.Right(), "foobar")
	}
}

func TestEither_IfLeft(t *testing.T) {
	left := Left("foo")

	var called = false
	left.IfLeft(func(t string) {
		called = true
	})

	if !called {
		t.Errorf("IfLeft() called = %v, want %v", called, true)
	}

	var called2 = false
	newR[string, string]("foo").IfLeft(func(t string) {
		called2 = true
	})

	if called2 {
		t.Errorf("IfLeft() called2 = %v, want %v", called2, false)
	}
}

func TestEither_IfRight(t *testing.T) {
	right := Right("foo")

	var called = false
	right.IfRight(func(t string) {
		called = true
	})

	if !called {
		t.Errorf("IfRight() called = %v, want %v", called, true)
	}

	var called2 = false
	newL[string, string]("foo").IfRight(func(t string) {
		called2 = true
	})

	if called2 {
		t.Errorf("IfRight() called2 = %v, want %v", called2, false)
	}
}

func TestEither_LeftOr(t *testing.T) {
	got := Left("foo").LeftOr("bar")

	if !reflect.DeepEqual(got, "foo") {
		t.Errorf("LeftOr() = %v, want %v", got, "foo")
	}

	got2 := newR[string, string]("foo").LeftOr("bar")

	if !reflect.DeepEqual(got2, "bar") {
		t.Errorf("LeftOr() = %v, want %v", got2, "bar")
	}
}

func TestEither_RightOr(t *testing.T) {
	got := Right("foo").RightOr("bar")

	if !reflect.DeepEqual(got, "foo") {
		t.Errorf("RightOr() = %v, want %v", got, "foo")
	}

	got2 := newL[string, string]("foo").RightOr("bar")

	if !reflect.DeepEqual(got2, "bar") {
		t.Errorf("RightOr() = %v, want %v", got2, "bar")
	}
}

func TestEither_LeftOrElse(t *testing.T) {
	left := Left("foo")

	got := left.LeftOrElse(func() string {
		return "bar"
	})

	if !reflect.DeepEqual(got, "foo") {
		t.Errorf("LeftOrElse() = %v, want %v", got, "foo")
	}

	got2 := newR[string, string]("foo").LeftOrElse(func() string {
		return "bar"
	})

	if !reflect.DeepEqual(got2, "bar") {
		t.Errorf("LeftOrElse() = %v, want %v", got2, "bar")
	}
}

func TestEither_RightOrElse(t *testing.T) {
	right := Right("foo")

	got := right.RightOrElse(func() string {
		return "bar"
	})

	if !reflect.DeepEqual(got, "foo") {
		t.Errorf("RightOrElse() = %v, want %v", got, "foo")
	}

	got2 := newL[string, string]("foo").RightOrElse(func() string {
		return "bar"
	})

	if !reflect.DeepEqual(got2, "bar") {
		t.Errorf("RightOrElse() = %v, want %v", got2, "bar")
	}
}

func TestEither_IfNotLeft(t *testing.T) {
	left := Left("foo")

	var called = false
	left.IfNotLeft(func() {
		called = true
	})

	if called {
		t.Errorf("IfNotLeft() called = %v, want %v", called, false)
	}

	var called2 = false
	newR[string, string]("foo").IfNotLeft(func() {
		called2 = true
	})

	if !called2 {
		t.Errorf("IfNotLeft() called2 = %v, want %v", called2, true)
	}
}

func TestEither_IfNotRight(t *testing.T) {
	right := Right("foo")

	var called = false
	right.IfNotRight(func() {
		called = true
	})

	if called {
		t.Errorf("IfNotRight() called = %v, want %v", called, false)
	}

	var called2 = false
	newL[string, string]("foo").IfNotRight(func() {
		called2 = true
	})

	if !called2 {
		t.Errorf("IfNotRight() called2 = %v, want %v", called2, true)
	}
}

func TestEither_MaybeLeft(t *testing.T) {
	got := Left("foo").MaybeLeft()

	if !got.IsSome() {
		t.Errorf("MaybeLeft() = %v, want %v", got.IsSome(), true)
		if !reflect.DeepEqual(got.Unwrap(), "foo") {
			t.Errorf("MaybeLeft() = %v, want %v", got.Unwrap(), "foo")
		}
	}

	got2 := Right("foo").MaybeLeft()

	if got2.IsSome() {
		t.Errorf("MaybeLeft() = %v, want %v", got.IsSome(), false)
	}
}

func TestEither_MaybeRight(t *testing.T) {
	got := Right("foo").MaybeRight()

	if !got.IsSome() {
		t.Errorf("MaybeRight() = %v, want %v", got.IsSome(), true)
		if !reflect.DeepEqual(got.Unwrap(), "foo") {
			t.Errorf("MaybeRight() = %v, want %v", got.Unwrap(), "foo")
		}
	}

	got2 := Left("foo").MaybeRight()

	if got2.IsSome() {
		t.Errorf("MaybeRight() = %v, want %v", got2.IsSome(), false)
	}
}

func TestEither_Swap(t *testing.T) {
	got := Right("foo").Swap()

	if !got.IsLeft() {
		t.Errorf("Swap() = %v, want %v", got.IsLeft(), true)
		if !reflect.DeepEqual(got.Left(), "foo") {
			t.Errorf("Swap() = %v, want %v", got.Left(), "foo")
		}
	}

	got2 := Left("foo").Swap()

	if !got2.IsRight() {
		t.Errorf("Swap() = %v, want %v", got2.IsRight(), true)
		if !reflect.DeepEqual(got2.Right(), "foo") {
			t.Errorf("Swap() = %v, want %v", got2.Right(), "foo")
		}
	}
}

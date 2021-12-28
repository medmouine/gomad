package maybe

import (
	"reflect"
	"testing"
)

func TestMaybe_Just(t *testing.T) {
	integer := 123
	gotInt := Just(integer)

	if !reflect.DeepEqual(gotInt.Unwrap(), integer) {
		t.Errorf("Just() = %v, want %v", gotInt.Unwrap(), integer)
	}
	if !reflect.DeepEqual(gotInt.IsNil(), false) {
		t.Errorf("Just() = %v, want %v", gotInt.IsNil(), false)
	}

	_string := "abc"
	gotStr := Just(_string)

	if !reflect.DeepEqual(gotStr.Unwrap(), _string) {
		t.Errorf("Just() = %v, want %v", gotStr.Unwrap(), _string)
	}

	_slice := []int{1, 2, 3}
	gotSlice := Just(_slice)

	if !reflect.DeepEqual(gotSlice.Unwrap(), _slice) {
		t.Errorf("Just() = %v, want %v", gotSlice.Unwrap(), _slice)
	}
}

func TestMaybe_None(t *testing.T) {
	got := None[int]()

	if !got.IsNil() {
		t.Errorf("None() = %v, want %v", got.IsNil(), false)
	}
}

func TestMaybe_Nillable(t *testing.T) {
	integer := 123
	gotInteger := Nillable(&integer)

	if !reflect.DeepEqual(gotInteger.Unwrap(), integer) {
		t.Errorf("Nillable() = %v, want %v", gotInteger.Unwrap(), integer)
	}
	if gotInteger.IsNil() {
		t.Errorf("Nillable() = %v, want %v", gotInteger.IsNil(), false)
	}

	gotNil := Nillable[int](nil)

	if !gotNil.IsNil() {
		t.Errorf("Nillable() = %v, want %v", gotNil.IsNil(), true)
	}
}

func TestMaybe_Apply(t *testing.T) {
	integer := 1
	m := Just(integer)

	var called = false
	m.Apply(func(v int) {
		called = true
	})

	if !called {
		t.Errorf("Apply() called = %v, want %v", called, true)
	}
	if !reflect.DeepEqual(m.Unwrap(), integer) {
		t.Errorf("Apply() = %v, want %v", m.Unwrap(), integer)
	}

	var called2 = false
	None[int]().Apply(func(v int) {
		called2 = true
	})

	if called2 {
		t.Errorf("Apply() called = %v, want %v", called2, false)
	}
}

func TestMaybe_Map(t *testing.T) {
	got := Just(1).Map(func(v int) int {
		return v + 3
	})

	if !reflect.DeepEqual(got.Unwrap(), 4) {
		t.Errorf("Map() = %v, want %v", got.Unwrap(), 4)
	}
	if !got.IsSome() {
		t.Errorf("Map() = %v, want %v", got.IsSome(), true)
	}

	got2 := None[int]().Map(func(v int) int {
		return v + 3
	})

	if !got2.IsNil() {
		t.Errorf("Map() = %v, want %v", got2.IsNil(), true)
	}
}

func TestMaybe_Unwrap(t *testing.T) {

	got := Just(1).Unwrap()

	if !reflect.DeepEqual(got, 1) {
		t.Errorf("Unwrap() = %v, want %v", got, 1)
	}

	defer func() { recover() }()
	None[int]().Unwrap()
	t.Errorf("Unwrap() on Nil did not panic")
}

func TestMaybe_OrElse(t *testing.T) {

	got := None[int]().OrElse(func() int {
		return 3
	})

	if !reflect.DeepEqual(got, 3) {
		t.Errorf("OrElse() = %v, want %v", got, 3)
	}

	got2 := Just(1).OrElse(func() int {
		return 3
	})

	if !reflect.DeepEqual(got2, 1) {
		t.Errorf("OrElse() = %v, want %v", got2, 1)
	}
}

func TestMaybe_Or(t *testing.T) {

	got := None[int]().Or(3)

	if !reflect.DeepEqual(got, 3) {
		t.Errorf("Or() = %v, want %v", got, 3)
	}

	got2 := Just(1).Or(3)

	if !reflect.DeepEqual(got2, 1) {
		t.Errorf("Or() = %v, want %v", got2, 1)
	}
}

func TestMaybe_OrNil(t *testing.T) {
	got := None[int]().OrNil()

	if got != nil {
		t.Errorf("OrNil() = %v, want %v", got, nil)
	}

	got2 := Just(4).OrNil()

	if !reflect.DeepEqual(*got2, 4) {
		t.Errorf("OrNil() = %v, want %v", *got2, 4)
	}
}

func TestMaybe_IsNone(t *testing.T) {
	got := None[int]().IsNil()

	if !got {
		t.Errorf("IsNil() = %v, want %v", got, true)
	}

	got2 := Just(4).IsNil()

	if got2 {
		t.Errorf("IsNil() = %v, want %v", got2, false)
	}
}

func TestMaybe_IsSome(t *testing.T) {
	got := Just(1).IsSome()

	if !got {
		t.Errorf("IsSome() = %v, want %v", got, true)
	}

	got2 := None[int]().IsSome()

	if got2 {
		t.Errorf("IsSome() = %v, want %v", got2, false)
	}
}

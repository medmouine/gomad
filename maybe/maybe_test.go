package maybe

import (
	"reflect"
	"testing"
)

func TestMaybe_Just(t *testing.T) {
	integer := 123
	gotInt := Just(integer)

	if !reflect.DeepEqual(gotInt.value, &integer) {
		t.Errorf("Just() = %v, want %v", gotInt.value, &integer)
	}
	if !reflect.DeepEqual(gotInt.isNil, false) {
		t.Errorf("Just() = %v, want %v", gotInt.isNil, false)
	}

	_string := "abc"
	gotStr := Just(_string)

	if !reflect.DeepEqual(gotStr.value, &_string) {
		t.Errorf("Just() = %v, want %v", gotStr.value, &_string)
	}
	if !reflect.DeepEqual(gotInt.isNil, false) {
		t.Errorf("Just() = %v, want %v", gotStr.isNil, false)
	}

	_slice := []int{1, 2, 3}
	gotSlice := Just(_slice)

	if !reflect.DeepEqual(gotSlice.value, &_slice) {
		t.Errorf("Just() = %v, want %v", gotSlice.value, &_slice)
	}
	if !reflect.DeepEqual(gotInt.isNil, false) {
		t.Errorf("Just() = %v, want %v", gotSlice.isNil, false)
	}
}

func TestMaybe_None(t *testing.T) {
	got := None[int]()

	if got.value != nil {
		t.Errorf("None() = %v, want %v", got.value, nil)
	}
	if !reflect.DeepEqual(got.isNil, true) {
		t.Errorf("None() = %v, want %v", got.isNil, true)
	}
}

func TestMaybe_Nullable(t *testing.T) {
	integer := 123
	gotInteger := Nullable(&integer)

	if !reflect.DeepEqual(gotInteger.value, &integer) {
		t.Errorf("Nullable() = %v, want %v", gotInteger.value, &integer)
	}
	if !reflect.DeepEqual(gotInteger.isNil, false) {
		t.Errorf("Nullable() = %v, want %v", gotInteger.isNil, false)
	}

	gotNil := Nullable[int](nil)

	if gotNil.value != nil {
		t.Errorf("Nullable() = %v, want %v", gotNil.value, nil)
	}
	if !reflect.DeepEqual(gotNil.isNil, true) {
		t.Errorf("Nullable() = %v, want %v", gotNil.isNil, true)
	}
}

func TestMaybe_Apply(t *testing.T) {
	integer := 1
	m := Just(integer)

	var got int
	m.Apply(func(v int) {
		got = v
	})

	if !reflect.DeepEqual(got, integer) {
		t.Errorf("Apply() = %v, want %v", got, integer)
	}
	if !reflect.DeepEqual(m.value, &integer) {
		t.Errorf("Apply() = %v, want %v", got, &integer)
	}
	if !reflect.DeepEqual(m.isNil, false) {
		t.Errorf("Apply() = %v, want %v", m.isNil, false)
	}
}

func TestMaybe_Map(t *testing.T) {
	integer := 1
	m := Just(integer)

	m.Map(func(v int) int {
		return v + 3
	})

	if !reflect.DeepEqual(*m.value, 4) {
		t.Errorf("Map() = %v, want %v", m.value, 4)
	}
	if !reflect.DeepEqual(m.isNil, false) {
		t.Errorf("Map() = %v, want %v", m.isNil, false)
	}
}

func TestMaybe_Unwrap(t *testing.T) {
	integer := 1
	m := Just(integer)

	got := *m.Unwrap()

	if !reflect.DeepEqual(got, 1) {
		t.Errorf("Unwrap() = %v, want %v", got, 1)
	}
}

func TestMaybe_OrElse(t *testing.T) {
	m := Nullable[int](nil)

	got := m.OrElse(3)

	if !reflect.DeepEqual(got, 3) {
		t.Errorf("OrElse() = %v, want %v", got, 3)
	}

	val := 4
	m2 := Nullable[int](&val)

	got2 := m2.OrElse(3)

	if !reflect.DeepEqual(got2, 4) {
		t.Errorf("OrElse() = %v, want %v", got2, 4)
	}
}

func TestMaybe_OrNil(t *testing.T) {
	m := Nullable[int](nil)

	got := m.OrNil()

	if got != nil {
		t.Errorf("OrNil() = %v, want %v", got, nil)
	}

	val := 4
	m2 := Nullable[int](&val)

	got2 := *m2.OrNil()

	if !reflect.DeepEqual(got2, 4) {
		t.Errorf("OrNil() = %v, want %v", got2, 4)
	}
}

func TestMaybe_IsNone(t *testing.T) {
	m := Nullable[int](nil)

	if !m.IsNone() {
		t.Errorf("IsNone() = %v, want %v", m.IsNone(), true)
	}

	val := 4
	m2 := Nullable[int](&val)

	if m2.IsNone() {
		t.Errorf("IsNone() = %v, want %v", m.IsNone(), false)
	}
}

func TestMaybe_IsSome(t *testing.T) {
	m := Nullable[int](nil)

	if m.IsSome() {
		t.Errorf("IsSome() = %v, want %v", m.IsSome(), false)
	}

	val := 4
	m2 := Nullable[int](&val)

	if !m2.IsSome() {
		t.Errorf("IsSome() = %v, want %v", m.IsSome(), true)
	}
}
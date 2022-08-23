package maybe_test

import (
	"reflect"
	"strconv"
	"testing"

	. "github.com/medmouine/gomad/maybe"
)

func TestJust(t *testing.T) {
	t.Parallel()

	integer := 123
	gotInt := Just(integer)

	if !reflect.DeepEqual(*gotInt.Unwrap(), integer) {
		t.Errorf("Just() = %v, want %v", *gotInt.Unwrap(), integer)
	}

	if !reflect.DeepEqual(gotInt.IsNil(), false) {
		t.Errorf("Just() = %v, want %v", gotInt.IsNil(), false)
	}

	_string := "abc"
	gotStr := Just(_string)

	if !reflect.DeepEqual(*gotStr.Unwrap(), _string) {
		t.Errorf("Just() = %v, want %v", *gotStr.Unwrap(), _string)
	}

	_slice := []int{1, 2, 3}
	gotSlice := Just(_slice)

	if !reflect.DeepEqual(*gotSlice.Unwrap(), _slice) {
		t.Errorf("Just() = %v, want %v", *gotSlice.Unwrap(), _slice)
	}
}

func TestNone(t *testing.T) {
	t.Parallel()

	got := None[int]()
	if !got.IsNil() {
		t.Errorf("None() = %v, want %v", got.IsNil(), false)
	}
}

func TestMap(t *testing.T) {
	t.Parallel()

	got := Map(Just(1), func(v int) string {
		return strconv.Itoa(v) + "!"
	})

	if !reflect.DeepEqual(*got.Unwrap(), "1!") {
		t.Errorf("Map() = %v, want %v", *got.Unwrap(), "1!")
	}

	got = Map(None[int](), func(v int) string {
		return strconv.Itoa(v) + "!"
	})

	if !got.IsNil() {
		t.Errorf("Map() = %v, want %v", got.IsNil(), true)
	}
}

func TestOf(t *testing.T) {
	t.Parallel()

	integer := 123
	gotInt := Of(&integer)

	if !reflect.DeepEqual(*gotInt.Unwrap(), integer) {
		t.Errorf("Of() = %v, want %v", *gotInt.Unwrap(), integer)
	}

	if gotInt.IsNil() {
		t.Errorf("Of() = %v, want %v", gotInt.IsNil(), false)
	}

	gotNil := Of[int](nil)

	if !gotNil.IsNil() {
		t.Errorf("Of() = %v, want %v", gotNil.IsNil(), true)
	}
}

func TestMaybe_Apply(t *testing.T) {
	t.Parallel()

	integer := 1
	m := Just(integer)

	called := false

	m.Apply(func(v int) {
		called = true
	})

	if !called {
		t.Errorf("Apply() called = %v, want %v", called, true)
	}

	if !reflect.DeepEqual(*m.Unwrap(), integer) {
		t.Errorf("Apply() = %v, want %v", *m.Unwrap(), integer)
	}

	called2 := false

	None[int]().Apply(func(v int) {
		called2 = true
	})

	if called2 {
		t.Errorf("Apply() called = %v, want %v", called2, false)
	}
}

func TestMaybe_Map(t *testing.T) {
	t.Parallel()

	got := Just(1).Map(func(v int) int {
		return v + 3
	})
	if !reflect.DeepEqual(*got.Unwrap(), 4) {
		t.Errorf("Map() = %v, want %v", *got.Unwrap(), 4)
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
	t.Parallel()

	if got := Just(1).Unwrap(); !reflect.DeepEqual(*got, 1) {
		t.Errorf("Unwrap() = %v, want %v", *got, 1)
	}

	defer func() {
		err := recover()
		if err == nil {
			t.Errorf("Unwrap() on Nil did not panic")
		}
	}()
	None[int]().Unwrap()
}

func TestMaybe_OrElse(t *testing.T) {
	t.Parallel()

	got := None[int]().OrElse(func() int {
		return 3
	})
	if !reflect.DeepEqual(*got, 3) {
		t.Errorf("OrElse() = %v, want %v", *got, 3)
	}

	got2 := Just(1).OrElse(func() int {
		return 3
	})
	if !reflect.DeepEqual(*got2, 1) {
		t.Errorf("OrElse() = %v, want %v", *got2, 1)
	}
}

func TestMaybe_Or(t *testing.T) {
	t.Parallel()

	if got := None[int]().Or(3); !reflect.DeepEqual(*got, 3) {
		t.Errorf("Or() = %v, want %v", *got, 3)
	}

	if got2 := Just(1).Or(3); !reflect.DeepEqual(*got2, 1) {
		t.Errorf("Or() = %v, want %v", *got2, 1)
	}
}

func TestMaybe_OrNil(t *testing.T) {
	t.Parallel()

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
	t.Parallel()

	if got := None[int]().IsNil(); !got {
		t.Errorf("IsNil() = %v, want %v", got, true)
	}

	if got2 := Just(4).IsNil(); got2 {
		t.Errorf("IsNil() = %v, want %v", got2, false)
	}
}

func TestMaybe_IsSome(t *testing.T) {
	t.Parallel()

	if got := Just(1).IsSome(); !got {
		t.Errorf("IsSome() = %v, want %v", got, true)
	}

	if got2 := None[int]().IsSome(); got2 {
		t.Errorf("IsSome() = %v, want %v", got2, false)
	}
}

func TestMaybe_Bind(t *testing.T) {
	t.Parallel()

	got := Just(2).Bind(func(t int) Maybe[int] {
		return Just(t * t)
	})
	if !reflect.DeepEqual(got, Just(4)) {
		t.Errorf("Bind() = %v, want %v", got, Just(4))
	}

	got2 := None[int]().Bind(func(t int) Maybe[int] {
		return Just(t * t)
	})
	if !reflect.DeepEqual(got2, None[int]()) {
		t.Errorf("Bind() = %v, want %v", got2, None[int]())
	}
}

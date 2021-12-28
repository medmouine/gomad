package result

import (
	"errors"
	"github.com/medmouine/gomad/maybe"
	"reflect"
	"testing"
)

func TestOk(t *testing.T) {
	got := Ok(1)

	if !got.IsOk() {
		t.Errorf("Ok() = %v, want %v", got.IsOk(), true)
	}
	if got.IsErr() {
		t.Errorf("Ok() = %v, want %v", got.IsErr(), false)
	}
	if !reflect.DeepEqual(got.Ok(), 1) {
		t.Errorf("Ok() = %v, want %v", got.Ok(), 1)
	}
}

func TestErr(t *testing.T) {
	got := Err[int](errors.New("error"))

	if got.IsOk() {
		t.Errorf("Err() = %v, want %v", got.IsOk(), false)
	}
	if !got.IsErr() {
		t.Errorf("Err() = %v, want %v", got.IsErr(), true)
	}
	if !reflect.DeepEqual(got.Err(), errors.New("error")) {
		t.Errorf("Err() = %v, want %v", got.Err(), errors.New("error"))
	}
}

func TestOf(t *testing.T) {
	got := Of(1, nil)

	if !got.IsOk() {
		t.Errorf("Of() = %v, want %v", got.IsOk(), true)
	}
	if got.IsErr() {
		t.Errorf("Of() = %v, want %v", got.IsErr(), false)
	}
	if !reflect.DeepEqual(got.Ok(), 1) {
		t.Errorf("Of() = %v, want %v", got.Ok(), 1)
	}

	got2 := Of(1, errors.New("error"))

	if got2.IsOk() {
		t.Errorf("Of() = %v, want %v", got2.IsOk(), false)
	}
	if !got2.IsErr() {
		t.Errorf("Of() = %v, want %v", got2.IsErr(), true)
	}
	if !reflect.DeepEqual(got2.Err(), errors.New("error")) {
		t.Errorf("Of() = %v, want %v", got2.Err(), errors.New("error"))
	}
}

func TestFromMaybe(t *testing.T) {
	got := FromMaybe(maybe.Just(1), errors.New("maybe is nil"))

	if !got.IsOk() {
		t.Errorf("FromMaybe() = %v, want %v", got.IsOk(), true)
	}
	if got.IsErr() {
		t.Errorf("FromMaybe() = %v, want %v", got.IsErr(), false)
	}
	if !reflect.DeepEqual(got.Ok(), 1) {
		t.Errorf("FromMaybe() = %v, want %v", got.Ok(), 1)
	}

	got2 := FromMaybe(maybe.None[int](), errors.New("maybe is nil"))

	if got2.IsOk() {
		t.Errorf("FromMaybe() = %v, want %v", got2.IsOk(), false)
	}
	if !got2.IsErr() {
		t.Errorf("FromMaybe() = %v, want %v", got2.IsErr(), true)
	}
	if !reflect.DeepEqual(got2.Err(), errors.New("maybe is nil")) {
		t.Errorf("FromMaybe() = %v, want %v", got2.Err(), errors.New("maybe is nil"))
	}
}

func TestResult_Err(t *testing.T) {
	got := Err[int](errors.New("error")).Err()

	if !reflect.DeepEqual(got, errors.New("error")) {
		t.Errorf("Err() = %v, want %v", got, errors.New("error"))
	}

	defer func() { recover() }()
	Ok(1).Err()
	t.Errorf("Err() on Ok did not panic")
}

func TestResult_IfErr(t *testing.T) {
	var called = false
	Err[int](errors.New("error")).IfErr(func(err error) {
		called = true
	})

	if !called {
		t.Errorf("IfErr() = %v, want %v", called, true)
	}

	var called2 = false
	Ok(1).IfErr(func(err error) {
		called2 = true
	})

	if called2 {
		t.Errorf("IfErr() = %v, want %v", called2, false)
	}
}

func TestResult_IfOk(t *testing.T) {
	var called = false
	Err[int](errors.New("error")).IfOk(func(i int) {
		called = true
	})

	if called {
		t.Errorf("IfOk() = %v, want %v", called, false)
	}

	var called2 = false
	Ok(1).IfOk(func(i int) {
		called2 = true
	})

	if !called2 {
		t.Errorf("IfOk() = %v, want %v", called2, true)
	}
}

func TestResult_IsErr(t *testing.T) {
	got := Err[int](errors.New("error")).IsErr()

	if !got {
		t.Errorf("IsErr() = %v, want %v", got, true)
	}

	got2 := Ok(1).IsErr()

	if got2 {
		t.Errorf("IsErr() = %v, want %v", got2, false)
	}
}

func TestResult_IsOk(t *testing.T) {
	got := Err[int](errors.New("error")).IsOk()

	if got {
		t.Errorf("IsOk() = %v, want %v", got, false)
	}

	got2 := Ok(1).IsOk()

	if !got2 {
		t.Errorf("IsOk() = %v, want %v", got2, true)
	}
}

func TestResult_Map(t *testing.T) {
	got := Err[int](errors.New("error")).Map(func(t int) int {
		return 5
	})

	if got.IsOk() {
		t.Errorf("Map() = %v, want %v", got.IsOk(), false)
	}

	got2 := Ok(1).Map(func(t int) int {
		return 5
	})

	if !reflect.DeepEqual(got2.Ok(), 5) {
		t.Errorf("Map() = %v, want %v", got2.Ok(), 5)
	}
}

func TestResult_MapErr(t *testing.T) {
	got := Err[int](errors.New("error")).MapErr(func(e error) error {
		return errors.New("new error")
	})

	if !reflect.DeepEqual(got.Err(), errors.New("new error")) {
		t.Errorf("MapErr() = %v, want %v", got.Err(), errors.New("new error"))
	}

	got2 := Ok(1).MapErr(func(err error) error {
		return errors.New("new error")
	})

	if !reflect.DeepEqual(got2.Ok(), 1) {
		t.Errorf("MapErr() = %v, want %v", got2.Ok(), 1)
	}
}

func TestResult_Maybe(t *testing.T) {
	got := Err[int](errors.New("error")).Maybe()

	if !reflect.DeepEqual(got, maybe.None[int]()) {
		t.Errorf("Maybe() = %v, want %v", got, maybe.None[int]())
	}

	got2 := Ok(1).Maybe()

	if !reflect.DeepEqual(got2, maybe.Just(1)) {
		t.Errorf("Maybe() = %v, want %v", got2, maybe.Just(1))
	}
}

func TestResult_Ok(t *testing.T) {
	got := Ok(1).Ok()

	if !reflect.DeepEqual(got, 1) {
		t.Errorf("Ok() = %v, want %v", got, 1)
	}

	defer func() { recover() }()
	Err[int](errors.New("error")).Ok()
	t.Errorf("Ok() on Err did not panic")
}

func TestResult_WithDefault(t *testing.T) {
	got := Ok(1).WithDefault(10)

	if !reflect.DeepEqual(got.Ok(), 1) {
		t.Errorf("WithDefault() = %v, want %v", got, 1)
	}

	got2 := Err[int](errors.New("error")).WithDefault(10)

	if !reflect.DeepEqual(got2.Ok(), 10) {
		t.Errorf("WithDefault() = %v, want %v", got2, 10)
	}
}

func TestResult_Or(t *testing.T) {
	got := Ok(1).Or(10)

	if !reflect.DeepEqual(got, 1) {
		t.Errorf("Or() = %v, want %v", got, 1)
	}
	got2 := Err[int](errors.New("error")).Or(10)

	if !reflect.DeepEqual(got2, 10) {
		t.Errorf("Or() = %v, want %v", got2, 10)
	}
}

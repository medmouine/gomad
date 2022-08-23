package result_test

import (
	"errors"
	"reflect"
	"strconv"
	"testing"

	"github.com/medmouine/gomad/maybe"
	. "github.com/medmouine/gomad/result"
)

var err = errors.New("error")

func TestOk(t *testing.T) {
	t.Parallel()

	got := Ok(1)
	if !got.IsOk() {
		t.Errorf("Ok() = %v, want %v", got.IsOk(), true)
	}

	if got.IsErr() {
		t.Errorf("Ok() = %v, want %v", got.IsErr(), false)
	}

	if !reflect.DeepEqual(*got.Ok(), 1) {
		t.Errorf("Ok() = %v, want %v", *got.Ok(), 1)
	}
}

func TestErr(t *testing.T) {
	t.Parallel()

	got := Err[int](err)
	if got.IsOk() {
		t.Errorf("Err() = %v, want %v", got.IsOk(), false)
	}

	if !got.IsErr() {
		t.Errorf("Err() = %v, want %v", got.IsErr(), true)
	}

	if !reflect.DeepEqual(got.Err(), err) {
		t.Errorf("Err() = %v, want %v", got.Err(), err)
	}
}

func TestOf(t *testing.T) {
	t.Parallel()

	got := Of(1, nil)
	if !got.IsOk() {
		t.Errorf("Of() = %v, want %v", got.IsOk(), true)
	}

	if got.IsErr() {
		t.Errorf("Of() = %v, want %v", got.IsErr(), false)
	}

	if !reflect.DeepEqual(*got.Ok(), 1) {
		t.Errorf("Of() = %v, want %v", *got.Ok(), 1)
	}

	got2 := Of(1, err)
	if got2.IsOk() {
		t.Errorf("Of() = %v, want %v", got2.IsOk(), false)
	}

	if !got2.IsErr() {
		t.Errorf("Of() = %v, want %v", got2.IsErr(), true)
	}

	if !reflect.DeepEqual(got2.Err(), err) {
		t.Errorf("Of() = %v, want %v", got2.Err(), err)
	}
}

func TestMap(t *testing.T) {
	t.Parallel()

	got := Map(Ok(1), func(i int) string {
		return strconv.Itoa(i) + "!"
	})

	if !reflect.DeepEqual(got, Ok("1!")) {
		t.Errorf("Map() = %v, want %v", got, Ok("1!"))
	}

	got2 := Map(Err[int](err), func(i int) string {
		return strconv.Itoa(i) + "!"
	})

	if !reflect.DeepEqual(got2, Err[string](err)) {
		t.Errorf("Map() = %v, want %v", got2, Err[string](err))
	}
}

func TestFromMaybe(t *testing.T) {
	t.Parallel()

	got := FromMaybe(maybe.Just(1), err)
	if !got.IsOk() {
		t.Errorf("FromMaybe() = %v, want %v", got.IsOk(), true)
	}

	if got.IsErr() {
		t.Errorf("FromMaybe() = %v, want %v", got.IsErr(), false)
	}

	if !reflect.DeepEqual(*got.Ok(), 1) {
		t.Errorf("FromMaybe() = %v, want %v", got.Ok(), 1)
	}

	got2 := FromMaybe(maybe.None[int](), err)
	if got2.IsOk() {
		t.Errorf("FromMaybe() = %v, want %v", got2.IsOk(), false)
	}

	if !got2.IsErr() {
		t.Errorf("FromMaybe() = %v, want %v", got2.IsErr(), true)
	}

	if !reflect.DeepEqual(got2.Err(), err) {
		t.Errorf("FromMaybe() = %v, want %v", got2.Err(), err)
	}
}

func TestResult_Err(t *testing.T) {
	t.Parallel()

	if got := Err[int](err).Err(); !reflect.DeepEqual(got, err) {
		t.Errorf("Err() = %v, want %v", got, err)
	}

	defer func() {
		if err := recover(); err == nil {
			t.Errorf("Err() on Ok did not panic")
		}
	}()
	Ok(1).Err()
}

func TestResult_IfErr(t *testing.T) {
	t.Parallel()

	called := false

	Err[int](err).IfErr(func(err error) {
		called = true
	})

	if !called {
		t.Errorf("IfErr() = %v, want %v", called, true)
	}

	called2 := false

	Ok(1).IfErr(func(err error) {
		called2 = true
	})

	if called2 {
		t.Errorf("IfErr() = %v, want %v", called2, false)
	}
}

func TestResult_IfOk(t *testing.T) {
	t.Parallel()

	called := false

	Err[int](err).IfOk(func(i int) {
		called = true
	})

	if called {
		t.Errorf("IfOk() = %v, want %v", called, false)
	}

	called2 := false

	Ok(1).IfOk(func(i int) {
		called2 = true
	})

	if !called2 {
		t.Errorf("IfOk() = %v, want %v", called2, true)
	}
}

func TestResult_IsErr(t *testing.T) {
	t.Parallel()

	if got := Err[int](err).IsErr(); !got {
		t.Errorf("IsErr() = %v, want %v", got, true)
	}

	if got2 := Ok(1).IsErr(); got2 {
		t.Errorf("IsErr() = %v, want %v", got2, false)
	}
}

func TestResult_IsOk(t *testing.T) {
	t.Parallel()

	got := Err[int](err).IsOk()
	if got {
		t.Errorf("IsOk() = %v, want %v", got, false)
	}

	got2 := Ok(1).IsOk()

	if !got2 {
		t.Errorf("IsOk() = %v, want %v", got2, true)
	}
}

func TestResult_Map(t *testing.T) {
	t.Parallel()

	got := Err[int](err).Map(func(t int) int {
		return 5
	})

	if got.IsOk() {
		t.Errorf("Map() = %v, want %v", got.IsOk(), false)
	}

	got2 := Ok(1).Map(func(t int) int {
		return 5
	})

	if !reflect.DeepEqual(*got2.Ok(), 5) {
		t.Errorf("Map() = %v, want %v", *got2.Ok(), 5)
	}
}

func TestResult_MapErr(t *testing.T) {
	t.Parallel()

	got := Err[int](err).MapErr(func(e error) error {
		return err
	})

	if !reflect.DeepEqual(got.Err(), err) {
		t.Errorf("MapErr() = %v, want %v", got.Err(), err)
	}

	got2 := Ok(1).MapErr(func(err error) error {
		return err
	})

	if !reflect.DeepEqual(*got2.Ok(), 1) {
		t.Errorf("MapErr() = %v, want %v", *got2.Ok(), 1)
	}
}

func TestResult_Maybe(t *testing.T) {
	t.Parallel()

	if got := Err[int](err).Maybe(); !reflect.DeepEqual(got, maybe.None[int]()) {
		t.Errorf("Maybe() = %v, want %v", got, maybe.None[int]())
	}

	if got2 := Ok(1).Maybe(); !reflect.DeepEqual(got2, maybe.Just(1)) {
		t.Errorf("Maybe() = %v, want %v", got2, maybe.Just(1))
	}
}

func TestResult_Ok(t *testing.T) {
	t.Parallel()

	if got := *Ok(1).Ok(); !reflect.DeepEqual(got, 1) {
		t.Errorf("Ok() = %v, want %v", got, 1)
	}

	defer func() { recover() }()
	Err[int](err).Ok()
	t.Errorf("Ok() on Err did not panic")
}

func TestResult_WithDefault(t *testing.T) {
	t.Parallel()

	if got := Ok(1).WithDefault(10); !reflect.DeepEqual(*got.Ok(), 1) {
		t.Errorf("WithDefault() = %v, want %v", *got.Ok(), 1)
	}

	got2 := Err[int](err).WithDefault(10)

	if !reflect.DeepEqual(*got2.Ok(), 10) {
		t.Errorf("WithDefault() = %v, want %v", got2, 10)
	}
}

func TestResult_Or(t *testing.T) {
	t.Parallel()

	if got := *Ok(1).Or(10); !reflect.DeepEqual(got, 1) {
		t.Errorf("Or() = %v, want %v", got, 1)
	}

	if got2 := *Err[int](err).Or(10); !reflect.DeepEqual(got2, 10) {
		t.Errorf("Or() = %v, want %v", got2, 10)
	}
}

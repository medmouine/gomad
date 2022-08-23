package result

import "github.com/medmouine/gomad/maybe"

/*
Result aims at abstracting all logic related to operations susceptible to failures, such as external API calls, etc.
It offers constructors and methods to safely manipulate the result in case of success and handle errors gracefully
in case of failure.
*/
type Result[T any] interface {
	WithDefault(T) Result[T]

	Ok() *T
	IsOk() bool
	IfOk(f func(T)) Result[T]
	Map(func(T) T) Result[T]
	Or(T) *T

	Err() error
	IsErr() bool
	IfErr(f func(error)) Result[T]
	MapErr(func(error) error) Result[T]

	Maybe() maybe.Maybe[T]
}

/*
Map applies a function to the value of the Result and returns a new Result of a new type.
*/
func Map[T2, T any](r Result[T], f func(T) T2) Result[T2] {
	if r.IsOk() {
		v := f(*r.Ok())

		return Ok(v)
	}

	return Err[T2](r.Err())
}

/*
Ok creates a new Result from a valid value.
*/
func Ok[T any](val T) Result[T] {
	return result[T]{val: &val}
}

/*
Err creates a new Result from an invalid value (error).
*/
func Err[T any](err error) Result[T] {
	return result[T]{err: err}
}

/*
FromMaybe creates a new Result from a Maybe instance.
*/
func FromMaybe[T any](m maybe.Maybe[T], err error) Result[T] {
	if m.IsNil() {
		return Err[T](err)
	}

	v := m.Unwrap()

	return Ok[T](*v)
}

/*
Of creates a new Result from a possibly valid value and an error.
*/
func Of[T any](val T, err error) Result[T] {
	if err != nil {
		return Err[T](err)
	}

	return Ok(val)
}

func (r result[T]) Ok() *T {
	if r.IsOk() {
		return r.val
	}

	panic(any("result.Ok() called on Err() result"))
}

func (r result[T]) Or(val T) *T {
	if r.IsOk() {
		return r.Ok()
	}

	return &val
}

func (r result[T]) Err() error {
	if r.IsErr() {
		return r.err
	}

	panic(any("result.Err() called on Ok() result"))
}

func (r result[T]) WithDefault(val T) Result[T] {
	if r.IsOk() {
		return r
	}

	return Ok(val)
}

func (r result[T]) Maybe() maybe.Maybe[T] {
	if r.IsErr() {
		return maybe.None[T]()
	}

	return maybe.Just[T](*r.Ok())
}

func (r result[T]) MapErr(f func(error) error) Result[T] {
	if r.IsErr() {
		return Err[T](f(r.Err()))
	}

	return r
}

func (r result[T]) Map(f func(T) T) Result[T] {
	if r.IsOk() {
		return Ok(f(*r.Ok()))
	}

	return r
}

func (r result[T]) IfErr(f func(error)) Result[T] {
	if r.IsErr() {
		f(r.err)
	}

	return r
}

func (r result[T]) IfOk(f func(T)) Result[T] {
	if r.IsOk() {
		f(*r.Ok())
	}

	return r
}

func (r result[T]) IsOk() bool {
	return !r.IsErr()
}

func (r result[T]) IsErr() bool {
	return r.err != nil
}

type result[T any] struct {
	Result[T]
	val *T
	err error
}

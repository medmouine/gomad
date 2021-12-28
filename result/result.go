package result

import "github.com/medmouine/gomad/maybe"

type Result[T any] interface {
	WithDefault(T) Result[T]

	Ok() T
	IsOk() bool
	IfOk(f func(T)) Result[T]
	Map(func(T) T) Result[T]
	Or(T) T

	Err() error
	IsErr() bool
	IfErr(f func(error)) Result[T]
	MapErr(func(error) error) Result[T]

	Maybe() maybe.Maybe[T]
}

type result[T any] struct {
	Result[T]
	val *T
	err error
}

func Ok[T any](val T) Result[T] {
	return result[T]{val: &val}
}

func Err[T any](err error) Result[T] {
	return result[T]{err: err}
}

func FromMaybe[T any](m maybe.Maybe[T], err error) Result[T] {
	if m.IsNil() {
		return Err[T](err)
	}
	return Ok(m.Unwrap())
}

func Of[T any](val T, err error) Result[T] {
	if err != nil {
		return Err[T](err)
	}
	return Ok(val)
}

func (r result[T]) Ok() T {
	if r.IsOk() {
		return *r.val
	}
	panic(any("result.Ok() called on Err() result"))
}

func (r result[T]) Or(val T) T {
	if r.IsOk() {
		return r.Ok()
	}
	return val
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
	return maybe.Just[T](r.Ok())
}

func (r result[T]) MapErr(f func(error) error) Result[T] {
	if r.IsErr() {
		return Err[T](f(r.Err()))
	}
	return r
}

func (r result[T]) Map(f func(T) T) Result[T] {
	if r.IsOk() {
		return Ok(f(r.Ok()))
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
		f(r.Ok())
	}
	return r
}

func (r result[T]) IsOk() bool {
	return !r.IsErr()
}

func (r result[T]) IsErr() bool {
	return r.err != nil
}

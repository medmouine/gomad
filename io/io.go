package io

type IO[T any] interface {
	Call() T
	Map(f func(T) T) IO[T]
}

type io[T any] struct {
	IO[T]
	f func() T
}

func Of[T any](t T) IO[T] {
	return io[T]{
		f: func() T {
			return t
		},
	}
}

func From[T any](f func() T) IO[T] {
	return io[T]{
		f: f,
	}
}

func Map[T any, U any](io IO[T], fn func(T) U) IO[U] {
	f := func() U {
		return fn(io.Call())
	}
	return From(f)
}

func (i io[T]) Call() T {
	return i.f()
}

func (i io[T]) Map(fn func(T) T) IO[T] {
	f := func() T {
		return fn(i.Call())
	}
	return From(f)
}

package io

type IO[T any] interface {
	Run() T
	Bind(func(T) IO[T]) IO[T]
}

type io[T any] struct {
	v T
	f func()
}

func Of[T any](v T, f func()) IO[T] {
	return &io[T]{v: v, f: f}
}

func (io io[T]) Run() T {
	io.f()
	return io.v
}

func (io io[T]) Bind(f func(T) IO[T]) IO[T] {
	return f(io.v)
}

package maybe

import "go/types"

type nillable interface {
	any | types.Nil
}

type Maybe[T any] interface {
	Unwrap() T
	Map(func(T) T) Maybe[T]
	Apply(func(T)) Maybe[T]
	IsSome() bool
	IsNil() bool
	OrElse(func() T) T
	Or(T) T
	OrNil() *T
}

type maybe[T any] struct {
	Maybe[T]
	val *T
}

func Nillable[T nillable](val *T) Maybe[T] {
	if val == nil {
		return maybe[T]{val: nil}
	}
	return maybe[T]{val: val}
}

func Just[T nillable](val T) Maybe[T] {
	return maybe[T]{val: &val}
}

func None[T nillable]() Maybe[T] {
	return maybe[T]{val: nil}
}

func (m maybe[T]) Unwrap() T {
	if m.IsSome() {
		return *m.val
	}
	panic(any("unwrap of empty Maybe"))
}

func (m maybe[T]) Map(f func(T) T) Maybe[T] {
	if m.IsSome() {
		return Just(f(*m.val))
	}
	return m
}

func (m maybe[T]) Apply(f func(x T)) Maybe[T] {
	if m.IsSome() {
		f(m.Unwrap())
	}
	return m
}

func (m maybe[T]) IsSome() bool {
	return m.val != nil
}

func (m maybe[T]) IsNil() bool {
	return m.val == nil
}

func (m maybe[T]) OrElse(f func() T) T {
	if m.IsNil() {
		return f()
	}
	return m.Unwrap()
}

func (m maybe[T]) OrNil() *T {
	if m.IsNil() {
		return nil
	}
	return m.val
}

func (m maybe[T]) Or(val T) T {
	if m.IsNil() {
		return val
	}
	return m.Unwrap()
}

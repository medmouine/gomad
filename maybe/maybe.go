package maybe

import "go/types"

type nillable interface {
	any | types.Nil
}

/*
Maybe is a monadic pattern allowing for data manipulation while abstracting whether the value actually exists or is nil.
For example, if we fetch data from an external API that could be nil, we can still perform manipulation on it while disregarding its actual state.
The Maybe struct will take care of managing the value itself. This is similar to the Maybe interface in Elm or Haskell or Optional in Java.
This is helpful for CRUD operations by simplifying the code and allowing for seamless manipulation of nullable data.
*/
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

/*
Of returns a new Maybe based on a value that may or may not be nil.
*/
func Of[T nillable](val *T) Maybe[T] {
	if val == nil {
		return maybe[T]{val: nil}
	}
	return maybe[T]{val: val}
}

/*
Just returns a new Maybe based on a value that we know is not nil.
*/
func Just[T nillable](val T) Maybe[T] {
	return maybe[T]{val: &val}
}

/*
None returns a new Maybe with an empty value we know is nil.
*/
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

package maybe

import "go/types"

type nullable interface {
	any | types.Nil
}

type Maybe[T any] struct {
	value *T
	isNil bool
}

func Nullable[U nullable](val *U) Maybe[U] {
	if val == nil {
		return Maybe[U]{nil, true}
	}
	return Maybe[U]{val, false}
}

func Just[U nullable](val U) Maybe[U] {
	if nullable(val) == nil {
		return Maybe[U]{nil, true}
	}
	return Maybe[U]{&val, false}
}

func None[V nullable]() Maybe[V] {
	return Maybe[V]{nil, true}
}

func (m Maybe[T]) Unwrap() *T {
	if !m.isNil {
		return m.value
	}
	panic(any("unwrap of empty Maybe"))
}

func (m *Maybe[T]) Map(f func(T) T) *Maybe[T] {
	if !m.isNil {
		v := f(*m.value)
		m.value = &v
	}
	return m
}

func (m Maybe[T]) Apply(f func(x T)) *Maybe[T] {
	if !m.isNil {
		v := *m.value
		f(v)
	}
	return &m
}

func (m *Maybe[T]) IsSome() bool {
	return !m.isNil
}

func (m *Maybe[T]) IsNone() bool {
	return m.isNil
}

func (m Maybe[T]) OrElse(val T) T {
	if m.isNil {
		return val
	}
	return *m.value
}

func (m Maybe[T]) OrNil() *T {
	if m.isNil {
		return nil
	}
	return m.value
}

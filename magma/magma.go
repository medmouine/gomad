package magma

import "github.com/medmouine/gomad/fold"

type Magma[T any] interface {
	Concat(T) T
}

type magma[T any] struct {
	values []T
	op     func(T, T) T
}

func Of[T any](values []T, op func(T, T) T) Magma[T] {
	return &magma[T]{values: values, op: op}
}

func (m magma[T]) Concat(x T) T {
	return fold.Left(m.values, m.op, x)
}

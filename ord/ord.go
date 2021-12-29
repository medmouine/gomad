package ord

import "github.com/medmouine/gomad/eq"

type Ord[T comparable] interface {
	Compare(T, T) int
}

type ord[T comparable] struct {
	compare func(T, T) int
}

func FromCompare[T comparable](compare func(T, T) int) Ord[T] {
	return &ord[T]{compare}
}

func Min[T comparable](o Ord[T]) func(T, T) T {
	return func(a, b T) T {
		if o.Compare(a, b) < 0 {
			return a
		}
		return b
	}
}

func Max[T comparable](o Ord[T]) func(T, T) T {
	return func(a, b T) T {
		if o.Compare(a, b) > 0 {
			return a
		}
		return b
	}
}

func Leq[T comparable](o Ord[T]) func(T, T) bool {
	return func(a, b T) bool {
		return o.Compare(a, b) <= 0
	}
}

func Geq[T comparable](o Ord[T]) func(T, T) bool {
	return func(a, b T) bool {
		return o.Compare(a, b) >= 0
	}
}

func Equals[T comparable](o Ord[T]) eq.Eq[T] {
	return eq.FromEquals(func(a, b T) bool {
		return o.Compare(a, b) == 0
	})
}

func Between[T comparable](o Ord[T]) func(T, T) func(T) bool {
	return func(a, b T) func(T) bool {
		return func(x T) bool {
			return o.Compare(a, x) <= 0 && o.Compare(x, b) <= 0
		}
	}
}

func Lt[T comparable](o Ord[T]) func(T, T) bool {
	return func(a, b T) bool {
		return o.Compare(a, b) < 0
	}
}

func Gt[T comparable](o Ord[T]) func(T, T) bool {
	return func(a, b T) bool {
		return o.Compare(a, b) > 0
	}
}

func (o ord[T]) Compare(a, b T) int {
	if o.compare(a, b) < 0 {
		return -1
	}
	if o.compare(a, b) > 0 {
		return 1
	}
	return 0
}

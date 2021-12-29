package monoid

import (
	"github.com/medmouine/gomad/semigroup"
)

type IMonoid[T any] interface {
	semigroup.ISemigroup[T]
	Empty() T
}

type Monoid[T comparable] struct {
	sg    semigroup.Semigroup[T]
	empty T
}

func (m Monoid[T]) Empty() T {
	return m.empty
}

func FromConcat[T comparable](concat func(T, T) T, empty T) Monoid[T] {
	return From(semigroup.FromConcat(concat), empty)
}

func From[T comparable](sg semigroup.Semigroup[T], empty T) Monoid[T] {
	return Monoid[T]{
		sg,
		empty,
	}
}

func (m Monoid[T]) Concat(a T, b T) T {
	if a == m.empty {
		return b
	}
	if b == m.empty {
		return a
	}
	return m.sg.Concat(a, b)
}

func Fold[T comparable](m Monoid[T]) func([]T) T {
	return semigroup.Fold[T](m.sg, m.Empty())
}

func FoldMap[T comparable](m Monoid[T], f func(T) T) func([]T) T {
	return semigroup.FoldMap[T](m.sg, f, m.Empty())
}

func FoldLeft[T comparable](m Monoid[T]) func([]T) T {
	return semigroup.FoldLeft[T](m.sg, m.Empty())
}

func FoldMapLeft[T comparable](m Monoid[T], f func(T) T) func([]T) T {
	return semigroup.FoldMapLeft[T](m.sg, f, m.Empty())
}

func (m Monoid[T]) Fold() func([]T) T {
	return Fold[T](m)
}

func (m Monoid[T]) FoldMap(f func(T) T) func([]T) T {
	return FoldMap[T](m, f)
}

func (m Monoid[T]) FoldLeft() func([]T) T {
	return FoldLeft[T](m)
}

func (m Monoid[T]) FoldMapLeft(f func(T) T) func([]T) T {
	return FoldMapLeft[T](m, f)
}

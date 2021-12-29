package semigroup

import "github.com/medmouine/gomad/identity"

type Semigroup[T any] interface {
	Concat(T, T) T
	Fold(init T) func(slice []T) T
	FoldF() func(slice []T) T

	FoldMap(fn func(T) T, init T) func(slice []T) T
	FoldMapF(fn func(T) T) func(slice []T) T

	FoldLeft(init T) func(slice []T) T
	FoldLeftF() func(slice []T) T

	FoldMapLeft(fn func(T) T, init T) func(slice []T) T
	FoldMapLeftF(fn func(T) T) func(slice []T) T
}

type semigroup[T any] struct {
	concat func(T, T) T
}

func FromConcat[T any](concat func(T, T) T) Semigroup[T] {
	return &semigroup[T]{concat}
}

func Reverse[T any](sg Semigroup[T]) Semigroup[T] {
	return semigroup[T]{
		func(a T, b T) T {
			return sg.Concat(b, a)
		},
	}
}

func Fold[T any](sg Semigroup[T], init T) func([]T) T {
	return FoldMap(sg, identity.Identity[T], init)
}

func FoldF[T any](sg Semigroup[T]) func([]T) T {
	return FoldMapF(sg, identity.Identity[T])
}

func FoldMap[T any](sg Semigroup[T], f func(T) T, init T) func([]T) T {
	return func(xs []T) T {
		var acc = f(init)
		for _, x := range xs {
			acc = sg.Concat(acc, f(x))
		}
		return acc
	}
}

func FoldMapF[T any](sg Semigroup[T], f func(T) T) func([]T) T {
	return func(xs []T) T {
		var acc = f(xs[0])
		for _, x := range xs[1:] {
			acc = sg.Concat(acc, f(x))
		}
		return acc
	}
}

func FoldLeft[T any](sg Semigroup[T], init T) func([]T) T {
	return FoldMapLeft(sg, identity.Identity[T], init)
}

func FoldLeftF[T any](sg Semigroup[T]) func([]T) T {
	return FoldMapLeftF(sg, identity.Identity[T])
}

func FoldMapLeft[T any](sg Semigroup[T], f func(T) T, init T) func([]T) T {
	return FoldMap(Reverse(sg), f, init)
}

func FoldMapLeftF[T any](sg Semigroup[T], f func(T) T) func([]T) T {
	return FoldMapF(Reverse(sg), f)
}

func (s semigroup[T]) Concat(a T, b T) T {
	return s.concat(a, b)
}

func (s semigroup[T]) Fold(init T) func([]T) T {
	return Fold[T](s, init)
}

func (s semigroup[T]) FoldF() func([]T) T {
	return FoldF[T](s)
}

func (s semigroup[T]) FoldMap(f func(T) T, init T) func([]T) T {
	return FoldMap[T](s, f, init)
}

func (s semigroup[T]) FoldMapF(f func(T) T) func([]T) T {
	return FoldMapF[T](s, f)
}

func (s semigroup[T]) FoldLeft(init T) func([]T) T {
	return FoldLeft[T](s, init)
}

func (s semigroup[T]) FoldLeftF() func([]T) T {
	return FoldLeftF[T](s)
}

func (s semigroup[T]) FoldMapLeft(f func(T) T, init T) func([]T) T {
	return FoldMapLeft[T](s, f, init)
}

func (s semigroup[T]) FoldMapLeftF(f func(T) T) func([]T) T {
	return FoldMapLeftF[T](s, f)
}

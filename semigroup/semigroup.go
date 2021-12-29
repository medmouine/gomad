package semigroup

import "github.com/medmouine/gomad/identity"

type ISemigroup[T any] interface {
	Concat(T, T) T
}

func FromConcat[T any](concat func(T, T) T) Semigroup[T] {
	return Semigroup[T]{concat: concat}
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

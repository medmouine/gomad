package semigroup

type Semigroup[T any] interface {
	Concat(T, T) T
}

type semigroup[T any] struct {
	concat func(T, T) T
}

func FromConcat[T any](concat func(T, T) T) Semigroup[T] {
	return &semigroup[T]{concat}
}

func (s semigroup[T]) Concat(a T, b T) T {
	return s.concat(a, b)
}

func Reverse[T any](sg Semigroup[T]) Semigroup[T] {
	return semigroup[T]{
		func(a T, b T) T {
			return sg.Concat(b, a)
		},
	}
}

package semigroup

func Reverse[T any](sg Semigroup[T]) Semigroup[T] {
	return Semigroup[T]{
		concat: func(a T, b T) T {
			return sg.Concat(b, a)
		},
	}
}

package semigroup

type Semigroup[T any] struct {
	concat func(T, T) T
}

func (s Semigroup[T]) Concat(a T, b T) T {
	return s.concat(a, b)
}

func (s Semigroup[T]) Fold(init T) func([]T) T {
	return Fold[T](s, init)
}

func (s Semigroup[T]) FoldF() func([]T) T {
	return FoldF[T](s)
}

func (s Semigroup[T]) FoldMap(f func(T) T, init T) func([]T) T {
	return FoldMap[T](s, f, init)
}

func (s Semigroup[T]) FoldMapF(f func(T) T) func([]T) T {
	return FoldMapF[T](s, f)
}

func (s Semigroup[T]) FoldLeft(init T) func([]T) T {
	return FoldLeft[T](s, init)
}

func (s Semigroup[T]) FoldLeftF() func([]T) T {
	return FoldLeftF[T](s)
}

func (s Semigroup[T]) FoldMapLeft(f func(T) T, init T) func([]T) T {
	return FoldMapLeft[T](s, f, init)
}

func (s Semigroup[T]) FoldMapLeftF(f func(T) T) func([]T) T {
	return FoldMapLeftF[T](s, f)
}

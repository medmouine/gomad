package functor

type IFunctor[T any] interface {
	Map(T) T
}

type Functor[T any] struct {
	Val T
}

func Lift[T any, U any](f func(T) U) func(Functor[T]) Functor[U] {
	return func(fa Functor[T]) Functor[U] {
		return Functor[U]{
			Val: f(fa.Val),
		}
	}
}

func Map[T any, U any](fa Functor[T], f func(T) U) Functor[U] {
	return Functor[U]{
		Val: f(fa.Val),
	}
}

func (fa Functor[T]) Map(f func(T) T) Functor[T] {
	return Functor[T]{
		Val: f(fa.Val),
	}
}

package functor

type IFunctor[T any] interface {
	Map(T) T
}

type Functor[T any] struct {
	val T
}

func Lift[T any, U any](f func(T) U) func(Functor[T]) Functor[U] {
	return func(fa Functor[T]) Functor[U] {
		return Functor[U]{
			val: f(fa.val),
		}
	}
}

func Map[T any, U any](fa Functor[T], f func(T) U) Functor[U] {
	return Functor[U]{
		val: f(fa.val),
	}
}

func (fa Functor[T]) Map(f func(T) T) Functor[T] {
	return Functor[T]{
		val: f(fa.val),
	}
}

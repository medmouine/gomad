package identity

type Identity[T any] interface {
	Bind(func(T) Identity[T]) Identity[T]
	Unwrap() T
}

type identity[T any] struct {
	v T
}

func Of[T any](v T) Identity[T] {
	return &identity[T]{v}
}

func (i *identity[T]) Bind(f func(T) Identity[T]) Identity[T] {
	return f(i.v)
}

func (i *identity[T]) Unwrap() T {
	return i.v
}

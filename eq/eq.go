package eq

type Eq[T comparable] interface {
	Equals(T, T) bool
}

type eq[T comparable] struct {
	equals func(T, T) bool
}

func FromEquals[T comparable](e func(T, T) bool) Eq[T] {
	return &eq[T]{e}
}

func (e eq[T]) Equals(x T, y T) bool {
	return e.equals(x, y)
}

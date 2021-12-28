package either

import "go/types"

type left[L any, R any] struct {
	*either[L, R]
	val L
}

func newL[L any, R any](val L) *left[L, R] {
	l := &left[L, R]{
		val: val,
	}

	l.either = &either[L, R]{
		Either: l,
		left:   l,
	}
	return l
}

func Left[L any](value L) Either[L, types.Nil] {
	return newL[L, types.Nil](value)
}

func (l left[L, R]) IsRight() bool {
	return false
}

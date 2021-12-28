package either

import "go/types"

type right[L any, R any] struct {
	*either[L, R]
	val R
}

func newR[L any, R any](val R) *right[L, R] {
	r := &right[L, R]{
		val: val,
	}

	r.either = &either[L, R]{
		Either: r,
		right:  r,
	}
	return r
}

/*
Right returns a new Either value with Right as the passed argument.
*/
func Right[R any](value R) Either[types.Nil, R] {
	return newR[types.Nil, R](value)
}

func (r right[L, R]) IsRight() bool {
	return true
}

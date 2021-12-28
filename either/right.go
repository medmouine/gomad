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

func Right[R any](value R) Either[types.Nil, R] {
	return newR[types.Nil, R](value)
}

func (r right[L, R]) IsRight() bool {
	return true
}

func (r *right[L, R]) Map(f func(R) R) *right[L, R] {
	r.val = f(r.val)
	return r
}

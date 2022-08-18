package either

type right[L, R any] struct {
	Either[L, R]
	val R
}

func newR[L, R any](val R) *right[L, R] {
	r := new(right[L, R])
	r.val = val
	r.Either = &either[L, R]{
		Either: r,
	}

	return r
}

/*
Right returns a new Either value with Right as the passed argument.
By default, the Left Type is the same as the Right Type.

*/
func Right[R any](value R) Either[R, R] {
	return newR[R, R](value)
}

func (r right[L, R]) Right() *R {
	return &r.val
}

func (r right[L, R]) Left() *L {
	panic(any("called Left on Right"))
}

func (r right[L, R]) IsLeft() bool {
	return false
}

func (r right[L, R]) IsRight() bool {
	return true
}

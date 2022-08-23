package either

type left[L, R any] struct {
	Either[L, R]
	val L
}

func newL[L, R any](val L) Either[L, R] {
	l := new(left[L, R])
	l.val = val
	l.Either = &either[L, R]{
		Either: l,
	}

	return l
}

/*
Left returns a new Either value with Left as the passed argument.
By default, the Right Type is the same as the Left Type.
*/
func Left[L any](value L) Either[L, L] {
	return newL[L, L](value)
}

func (l left[L, R]) Left() *L {
	return &l.val
}

func (left[L, R]) Right() *R {
	panic(any("called Right on Left"))
}

func (left[L, R]) IsLeft() bool {
	return true
}

func (left[L, R]) IsRight() bool {
	return false
}

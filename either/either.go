package either

import "github.com/medmouine/gomad/maybe"

type Either[L any, R any] interface {
	IsRight() bool
	MapRight(func(R) R) Either[L, R]
	Right() R
	RightOr(R) R
	RightOrElse(func() R) R
	IfRight(func(R)) Either[L, R]
	IfNotRight(func()) Either[L, R]
	MaybeRight() maybe.Maybe[R]

	IsLeft() bool
	MapLeft(func(L) L) Either[L, R]
	Left() L
	LeftOr(L) L
	LeftOrElse(func() L) L
	IfLeft(func(L)) Either[L, R]
	IfNotLeft(func()) Either[L, R]
	MaybeLeft() maybe.Maybe[L]

	Swap() Either[R, L]
}

type either[L any, R any] struct {
	Either[L, R]
	left  *left[L, R]
	right *right[L, R]
}

func FromPredicateC[L any, R any](predicate func(R) bool, left L) func(R) Either[L, R] {
	return func(candidate R) Either[L, R] {
		if predicate(candidate) {
			return newR[L, R](candidate)
		}
		return newL[L, R](left)
	}
}

func FromPredicate[L any, R any](predicate func(R) bool, right R, left L) Either[L, R] {
	return FromPredicateC[L, R](predicate, left)(right)
}

func (e either[L, R]) Swap() Either[R, L] {
	if e.IsRight() {
		return newL[R, L](e.Right())
	}
	return newR[R, L](e.Left())
}

func (e *either[L, R]) IsLeft() bool {
	return !e.IsRight()
}

func (e *either[L, R]) MapRight(f func(R) R) Either[L, R] {
	if e.IsRight() {
		return newR[L, R](f(e.Right()))
	}
	return e
}

func (e *either[L, R]) MapLeft(f func(L) L) Either[L, R] {
	if e.IsLeft() {
		return newL[L, R](f(e.Left()))
	}
	return e
}

func (e either[L, R]) Left() L {
	if e.IsLeft() {
		return e.left.val
	}
	panic(any("either is not Left"))
}

func (e either[L, R]) Right() R {
	if e.IsRight() {
		return e.right.val
	}
	panic(any("either is not Right"))
}

// ################
// #### Left #####
// ################

func (e either[L, R]) LeftOr(or L) L {
	if e.IsLeft() {
		return e.Left()
	}
	return or
}

func (e either[L, R]) LeftOrElse(orElse func() L) L {
	if e.IsLeft() {
		return e.Left()
	}
	return orElse()
}

func (e *either[L, R]) IfLeft(f func(L)) Either[L, R] {
	if e.IsLeft() {
		f(e.Left())
		return e
	}
	return e
}

func (e *either[L, R]) IfNotLeft(f func()) Either[L, R] {
	if !e.IsLeft() {
		f()
		return e
	}
	return e
}

func (e either[L, R]) MaybeLeft() maybe.Maybe[L] {
	if e.IsLeft() {
		return maybe.Just(e.Left())
	}
	return maybe.None[L]()
}

// ################
// #### Right #####
// ################

func (e either[L, R]) RightOr(or R) R {
	if e.IsRight() {
		return e.Right()
	}
	return or
}

func (e either[L, R]) RightOrElse(orElse func() R) R {
	if e.IsRight() {
		return e.Right()
	}
	return orElse()
}

func (e *either[L, R]) IfRight(f func(R)) Either[L, R] {
	if e.IsRight() {
		f(e.Right())
		return e
	}
	return e
}

func (e *either[L, R]) IfNotRight(f func()) Either[L, R] {
	if !e.IsRight() {
		f()
		return e
	}
	return e
}

func (e either[L, R]) MaybeRight() maybe.Maybe[R] {
	if e.IsRight() {
		return maybe.Just(e.Right())
	}
	return maybe.None[R]()
}

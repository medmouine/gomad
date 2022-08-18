package either

import "github.com/medmouine/gomad/maybe"

/*
Either allows to manipulate pairs of mutually exclusive data.
For example, if we would want to fall back to a value B if A answers to a specific predicate.
This interface allows integrating this behavior seamlessly by abstracting all the underlying logic of
managing both values.
A common use case for this is form validation for front-end applications.
*/
type Either[L, R any] interface {
	IsRight() bool
	Right() *R
	RightOr(R) *R
	RightOrElse(func() R) *R
	IfRight(func(R)) Either[L, R]
	IfNotRight(func()) Either[L, R]
	MaybeRight() maybe.Maybe[R]

	IsLeft() bool
	Left() *L
	LeftOr(L) *L
	LeftOrElse(func() L) *L
	IfLeft(func(L)) Either[L, R]
	IfNotLeft(func()) Either[L, R]
	MaybeLeft() maybe.Maybe[L]

	Swap() Either[R, L]
}

type either[L any, R any] struct {
	Either[L, R]
}

/*
FromPredicateC returns a function that creates a new Either based on a given predicate and a Left value in case of
failure of the predicate.
*/
func FromPredicateC[L, R any](predicate func(R) bool, left L) func(R) Either[L, R] {
	return func(candidate R) Either[L, R] {
		if predicate(candidate) {
			return newR[L, R](candidate)
		}

		return newL[L, R](left)
	}
}

/*
FromPredicate returns a new Either based on a given predicate and a Left value in case of failure of the predicate
and a Right value which will be tested against the predicate.
*/
func FromPredicate[L, R any](predicate func(R) bool, right R, left L) Either[L, R] {
	return FromPredicateC[L, R](predicate, left)(right)
}

func MapRight[R2, L, R any](e Either[L, R], f func(R) R2) Either[L, R2] {
	if e.IsRight() {
		return newR[L, R2](f(*e.Right()))
	}

	return newL[L, R2](*e.Left())
}

func MapLeft[L2, L, R any](e Either[L, R], f func(L) L2) Either[L2, R] {
	if e.IsLeft() {
		return newL[L2, R](f(*e.Left()))
	}

	return newR[L2, R](*e.Right())
}

func (e either[L, R]) Swap() Either[R, L] {
	if e.IsRight() {
		return newL[R, L](*e.Right())
	}

	return newR[R, L](*e.Left())
}

// ################
// #### Left #####
// ################

func (e either[L, R]) LeftOr(or L) *L {
	if e.IsLeft() {
		return e.Left()
	}

	return &or
}

func (e either[L, R]) LeftOrElse(orElse func() L) *L {
	if e.IsLeft() {
		return e.Left()
	}

	v := orElse()

	return &v
}

func (e either[L, R]) IfLeft(f func(L)) Either[L, R] {
	if e.IsLeft() {
		f(*e.Left())
	}

	return &e
}

func (e either[L, R]) IfNotLeft(f func()) Either[L, R] {
	if !e.IsLeft() {
		f()
	}

	return &e
}

func (e either[L, R]) MaybeLeft() maybe.Maybe[L] {
	if e.IsLeft() {
		return maybe.Just(*e.Left())
	}

	return maybe.None[L]()
}

// ################
// #### Right #####
// ################

func (e either[L, R]) RightOr(or R) *R {
	if e.IsRight() {
		return e.Right()
	}

	return &or
}

func (e either[L, R]) RightOrElse(orElse func() R) *R {
	if e.IsRight() {
		return e.Right()
	}

	v := orElse()

	return &v
}

func (e either[L, R]) IfRight(f func(R)) Either[L, R] {
	if e.IsRight() {
		f(*e.Right())
	}

	return &e
}

func (e either[L, R]) IfNotRight(f func()) Either[L, R] {
	if !e.IsRight() {
		f()
	}

	return &e
}

func (e either[L, R]) MaybeRight() maybe.Maybe[R] {
	if e.IsRight() {
		return maybe.Just(*e.Right())
	}

	return maybe.None[R]()
}

package functor

import "github.com/medmouine/gomad/HTK"

type Functor[T, U any] interface {
	HTK.HKT[T, U]
}

type functor[T, U any] struct {
	Functor[T, U]
}

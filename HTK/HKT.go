package HTK

type HKT[T, U any] interface {
	Unwrap() T
}

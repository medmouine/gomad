package fold

func Left[T any](coll []T, op func(T, T) T, initial T) T {
	for _, x := range coll {
		initial = op(initial, x)
	}

	// return the final result
	return initial
}

func Right[T any](coll []T, op func(T, T) T, initial T) T {
	for i := len(coll) - 1; i >= 0; i-- {
		initial = op(coll[i], initial)
	}

	return initial
}

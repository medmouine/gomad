package identity

func Identity[T any](i T) *T {
	return &i
}

package sli

func Copy[T any](s []T) []T {
	c := make([]T, len(s))
	copy(c, s)
	return c
}

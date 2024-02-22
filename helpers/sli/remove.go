package sli

func PopHeap[T any](s []T) (T, []T) {
	return s[0], s[1:]
}

func PopQueue[T any](s []T) (T, []T) {
	return s[len(s)-1], s[:len(s)-1]
}

func PushHeap[T any](s []T, e T) []T {
	s = append(s, e)
	return s
}

func PushQueue[T any](s []T, e T) []T {
	return append([]T{e}, s...)
}

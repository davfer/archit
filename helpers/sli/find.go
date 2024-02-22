package sli

func Find[T any](s []T, f func(T) bool) (t T, ok bool) {
	for _, v := range s {
		if f(v) {
			ok = true
			t = v
			return
		}
	}
	return
}

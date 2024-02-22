package opts

import "github.com/davfer/archit/helpers/ref"

type Opt[K any] func(K) K

func New[K any](options ...Opt[K]) (k K) {
	for _, o := range options {
		k = o(k)
	}
	return
}

func WithField[K any](field string, val any) Opt[K] {
	return func(s K) K {
		s, _ = ref.SetField(s, field, val)
		return s
	}
}

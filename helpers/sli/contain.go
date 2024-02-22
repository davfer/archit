package sli

import "golang.org/x/exp/slices"

func AppendIfMissing[T comparable](to []T, es ...T) []T {
	for _, e := range es {
		if !slices.Contains(to, e) {
			to = append(to, e)
		}
	}

	return to
}

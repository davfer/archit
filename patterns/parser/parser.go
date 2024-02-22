package parser

type Parser[K, T any] interface {
	Parse(T) (K, bool)
	Build(K) (T, bool)
}

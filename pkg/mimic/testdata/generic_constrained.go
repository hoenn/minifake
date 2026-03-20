package test

//go:generate go run github.com/hoenn/mimic ./generic_constrained.go Sorter,KeyValue
type Sorter[T comparable] interface {
	Sort(items []T) []T
	Contains(items []T, target T) bool
	Min(items []T) T
}

type KeyValue[K comparable, V any] interface {
	Get(key K) (V, bool)
	Set(key K, value V)
	Keys() []K
}

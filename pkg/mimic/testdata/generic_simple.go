package test

//go:generate go run github.com/hoenn/mimic ./generic_simple.go Cache
type Cache[T any] interface {
	Get(key string) (T, bool)
	Set(key string, value T)
	Delete(key string)
}

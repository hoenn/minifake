package test

//go:generate go run github.com/hoenn/mimic ./test_generic_embed.go ExtendedCache

type Reader[T any] interface {
	Read(key string) (T, bool)
}

type ExtendedCache[T any] interface {
	Reader[T]
	Write(key string, value T)
}

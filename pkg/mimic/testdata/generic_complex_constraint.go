package test

type Number interface {
	~int | ~float64
}

//go:generate go run github.com/hoenn/mimic ./generic_complex_constraint.go Math
type Math[T Number] interface {
	Add(a, b T) T
	Zero() T
}

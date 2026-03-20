package test

//go:generate go run github.com/hoenn/mimic ./generic_multi.go Converter,Swapper
type Converter[From, To any] interface {
	Convert(input From) (To, error)
	ConvertAll(inputs []From) ([]To, error)
}

type Swapper[From, To any] interface {
	Swap(a From, b To) (To, From)
	Ptr(input From) *To
}

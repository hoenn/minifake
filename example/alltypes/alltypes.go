package alltypes

//go:generate go run github.com/hoenn/minifake ./alltypes.go MyInterface
type MyInterface interface {
	BoolMethod(b bool) bool
	StringMethod(s string) string
	IntMethod(i int) int
	Int8Method(i int8) int8
	Int16Method(i int16) int16
	Int32Method(i int32) int32
	Int64Method(i int64) int64
	UintMethod(u uint) uint
	Uint8Method(u uint8) uint8
	Uint16Method(u uint16) uint16
	Uint32Method(u uint32) uint32
	Uint64Method(u uint64) uint64
	UintptrMethod(u uintptr) uintptr
	ByteMethod(b byte) byte
	RuneMethod(r rune) rune
	Float32Method(f float32) float32
	Float64Method(f float64) float64
	Complex64Method(c complex64) complex64
	Complex128Method(c complex128) complex128
	ChanMethod(ch chan int) chan int
	SliceMethod(s []int) []int
	ArrayMethod(a [5]int) [5]int
	MapMethod(m map[string]int) map[string]int
	InterfaceMethod(i interface{}) interface{}
	FunctionMethod(f func(int) int) func(int) int
	PointerMethod(p *int) *int
	UnnamedParam(int) bool
	VariadicParam(s string, ss ...string) bool
}

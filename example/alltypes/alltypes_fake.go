// Code generated by ministub; DO NOT EDIT.
package alltypes

// FakeMyInterface implements MyInterface.
type FakeMyInterface struct {
	BoolMethodStub       func(b bool) bool
	StringMethodStub     func(s string) string
	IntMethodStub        func(i int) int
	Int8MethodStub       func(i int8) int8
	Int16MethodStub      func(i int16) int16
	Int32MethodStub      func(i int32) int32
	Int64MethodStub      func(i int64) int64
	UintMethodStub       func(u uint) uint
	Uint8MethodStub      func(u uint8) uint8
	Uint16MethodStub     func(u uint16) uint16
	Uint32MethodStub     func(u uint32) uint32
	Uint64MethodStub     func(u uint64) uint64
	UintptrMethodStub    func(u uintptr) uintptr
	ByteMethodStub       func(b byte) byte
	RuneMethodStub       func(r rune) rune
	Float32MethodStub    func(f float32) float32
	Float64MethodStub    func(f float64) float64
	Complex64MethodStub  func(c complex64) complex64
	Complex128MethodStub func(c complex128) complex128
	ChanMethodStub       func(ch chan int) chan int
	SliceMethodStub      func(s []int) []int
	ArrayMethodStub      func(a [5]int) [5]int
	MapMethodStub        func(m map[string]int) map[string]int
	InterfaceMethodStub  func(i interface{}) interface{}
	FunctionMethodStub   func(f func(int) int) func(int) int
	PointerMethodStub    func(p *int) *int
	UnnamedParamStub     func(arg0 int) bool
	VariadicParamStub    func(s string, ss ...string) bool
}

func (fakeImpl *FakeMyInterface) BoolMethod(b bool) bool {
	return fakeImpl.BoolMethodStub(b)
}
func (fakeImpl *FakeMyInterface) StringMethod(s string) string {
	return fakeImpl.StringMethodStub(s)
}
func (fakeImpl *FakeMyInterface) IntMethod(i int) int {
	return fakeImpl.IntMethodStub(i)
}
func (fakeImpl *FakeMyInterface) Int8Method(i int8) int8 {
	return fakeImpl.Int8MethodStub(i)
}
func (fakeImpl *FakeMyInterface) Int16Method(i int16) int16 {
	return fakeImpl.Int16MethodStub(i)
}
func (fakeImpl *FakeMyInterface) Int32Method(i int32) int32 {
	return fakeImpl.Int32MethodStub(i)
}
func (fakeImpl *FakeMyInterface) Int64Method(i int64) int64 {
	return fakeImpl.Int64MethodStub(i)
}
func (fakeImpl *FakeMyInterface) UintMethod(u uint) uint {
	return fakeImpl.UintMethodStub(u)
}
func (fakeImpl *FakeMyInterface) Uint8Method(u uint8) uint8 {
	return fakeImpl.Uint8MethodStub(u)
}
func (fakeImpl *FakeMyInterface) Uint16Method(u uint16) uint16 {
	return fakeImpl.Uint16MethodStub(u)
}
func (fakeImpl *FakeMyInterface) Uint32Method(u uint32) uint32 {
	return fakeImpl.Uint32MethodStub(u)
}
func (fakeImpl *FakeMyInterface) Uint64Method(u uint64) uint64 {
	return fakeImpl.Uint64MethodStub(u)
}
func (fakeImpl *FakeMyInterface) UintptrMethod(u uintptr) uintptr {
	return fakeImpl.UintptrMethodStub(u)
}
func (fakeImpl *FakeMyInterface) ByteMethod(b byte) byte {
	return fakeImpl.ByteMethodStub(b)
}
func (fakeImpl *FakeMyInterface) RuneMethod(r rune) rune {
	return fakeImpl.RuneMethodStub(r)
}
func (fakeImpl *FakeMyInterface) Float32Method(f float32) float32 {
	return fakeImpl.Float32MethodStub(f)
}
func (fakeImpl *FakeMyInterface) Float64Method(f float64) float64 {
	return fakeImpl.Float64MethodStub(f)
}
func (fakeImpl *FakeMyInterface) Complex64Method(c complex64) complex64 {
	return fakeImpl.Complex64MethodStub(c)
}
func (fakeImpl *FakeMyInterface) Complex128Method(c complex128) complex128 {
	return fakeImpl.Complex128MethodStub(c)
}
func (fakeImpl *FakeMyInterface) ChanMethod(ch chan int) chan int {
	return fakeImpl.ChanMethodStub(ch)
}
func (fakeImpl *FakeMyInterface) SliceMethod(s []int) []int {
	return fakeImpl.SliceMethodStub(s)
}
func (fakeImpl *FakeMyInterface) ArrayMethod(a [5]int) [5]int {
	return fakeImpl.ArrayMethodStub(a)
}
func (fakeImpl *FakeMyInterface) MapMethod(m map[string]int) map[string]int {
	return fakeImpl.MapMethodStub(m)
}
func (fakeImpl *FakeMyInterface) InterfaceMethod(i interface{}) interface{} {
	return fakeImpl.InterfaceMethodStub(i)
}
func (fakeImpl *FakeMyInterface) FunctionMethod(f func(int) int) func(int) int {
	return fakeImpl.FunctionMethodStub(f)
}
func (fakeImpl *FakeMyInterface) PointerMethod(p *int) *int {
	return fakeImpl.PointerMethodStub(p)
}
func (fakeImpl *FakeMyInterface) UnnamedParam(arg0 int) bool {
	return fakeImpl.UnnamedParamStub(arg0)
}
func (fakeImpl *FakeMyInterface) VariadicParam(s string, ss ...string) bool {
	return fakeImpl.VariadicParamStub(s, ss...)
}

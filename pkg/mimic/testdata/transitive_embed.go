package test

// Transitive embedded interfaces, A -> B -> C
//   A = A.Methods U {B.Methods}
//   B = B.Methods U {C.Methods}
//   C = {C.Methods}
//   A = A.Methods U B.Methods U C.Methods

//go:generate go run github.com/hoenn/mimic ./transitive_embed.go InterfaceA,InterfaceB,InterfaceC
type InterfaceA interface {
	Foo() string
	InterfaceB
}

type InterfaceB interface {
	Bar() string
	InterfaceC
}

type InterfaceC interface {
	Qux() string
}

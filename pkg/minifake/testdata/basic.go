package test

import "testing"

//go:generate go run github.com/hoenn/minifake ./basic.go Basic
type Basic interface {
	Foo() error
	Bar(m map[string]string, qux bool)
	Abc(t *testing.T) error
}

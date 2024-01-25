package ministub

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseAndStub(t *testing.T) {
	src := `package test

	import "testing"

	type Tester interface {
		Foo() error
		Bar(m map[string]string, qux bool)
		Abc(t *testing.T) error
	}
`
	bs, err := ParseAndStub([]string{"Tester"}, "", src, true)
	require.NoError(t, err)

	expected := `// Code generated by ministub; DO NOT EDIT.
package test

import "testing"

// FakeTester implements Tester.
type FakeTester struct {
	FooStub func() error
	BarStub func(m map[string]string, qux bool)
	AbcStub func(t *testing.T) error
}

func (fakeImpl *FakeTester) Foo() error {
	return fakeImpl.FooStub()
}
func (fakeImpl *FakeTester) Bar(m map[string]string, qux bool) {
	return fakeImpl.BarStub(m, qux)
}
func (fakeImpl *FakeTester) Abc(t *testing.T) error {
	return fakeImpl.AbcStub(t)
}
`
	require.Equal(t, expected, string(bs))
}

func TestParseAndStubEmptyInterface(t *testing.T) {
	src := `package test

	type Tester interface {
	}
`
	bs, err := ParseAndStub([]string{"Tester"}, "", src, true)
	require.NoError(t, err)

	expected := `// Code generated by ministub; DO NOT EDIT.
package test

// FakeTester implements Tester.
type FakeTester struct {
}
`
	require.Equal(t, expected, string(bs))
}

func TestParseAndStubUnnamedParameters(t *testing.T) {
	src := `package test

type Tester interface {
	UnnamedParam(int) bool
	UnnamedParams(int, bool) int
}
`
	bs, err := ParseAndStub([]string{"Tester"}, "", src, true)
	require.NoError(t, err)

	expected := `// Code generated by ministub; DO NOT EDIT.
package test

// FakeTester implements Tester.
type FakeTester struct {
	UnnamedParamStub  func(arg0 int) bool
	UnnamedParamsStub func(arg0 int, arg1 bool) int
}

func (fakeImpl *FakeTester) UnnamedParam(arg0 int) bool {
	return fakeImpl.UnnamedParamStub(arg0)
}
func (fakeImpl *FakeTester) UnnamedParams(arg0 int, arg1 bool) int {
	return fakeImpl.UnnamedParamsStub(arg0, arg1)
}
`
	require.Equal(t, expected, string(bs))
}

func TestParseAndStubVariadicParam(t *testing.T) {
	src := `package test

type Tester interface {
	VariadicParam(s string, ss ...string) bool
	UnnamedVariadicParam(arg0 ...string) bool
}
`
	bs, err := ParseAndStub([]string{"Tester"}, "", src, true)
	require.NoError(t, err)

	expected := `// Code generated by ministub; DO NOT EDIT.
package test

// FakeTester implements Tester.
type FakeTester struct {
	VariadicParamStub        func(s string, ss ...string) bool
	UnnamedVariadicParamStub func(arg0 ...string) bool
}

func (fakeImpl *FakeTester) VariadicParam(s string, ss ...string) bool {
	return fakeImpl.VariadicParamStub(s, ss...)
}
func (fakeImpl *FakeTester) UnnamedVariadicParam(arg0 ...string) bool {
	return fakeImpl.UnnamedVariadicParamStub(arg0...)
}
`
	require.Equal(t, expected, string(bs))
}

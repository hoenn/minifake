package ministub

import (
	"context"
	"testing"
)

type Tester interface {
	Foo(ctx context.Context, s string) error
	Bar(m map[string]string, qux bool)
	Buz(b bool, s ...string) (bool, error)
	Abc(t *testing.T) error
}

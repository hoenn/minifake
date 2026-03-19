package test

import "io"

//go:generate go run github.com/hoenn/mimic ./crosspkgembed.go CrossPackageEmbed
type CrossPackageEmbed interface {
	io.Reader
}

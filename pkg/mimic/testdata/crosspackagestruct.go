package test

import "image"

//go:generate go run github.com/hoenn/mimic ./crosspackagestruct.go CrossPackageStruct
type CrossPackageStruct interface {
	image.Image
}

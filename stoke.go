package gg

import (
	"github.com/golang/freetype/raster"
	"golang.org/x/image/math/fixed"
)

// MiterJoiner adds bevel joins to a stroked path.
var MiterJoiner raster.Joiner = raster.JoinerFunc(miterJoiner)

func miterJoiner(lhs, rhs raster.Adder, haflWidth fixed.Int26_6, pivot, n0, n1 fixed.Point26_6) {
	lhs.Add1(pivot.Add(n1).Add(n0))
	rhs.Add1(pivot.Sub(n1))
}

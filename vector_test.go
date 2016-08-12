package geo

import (
  "math"
  "testing"
)

var v1 = NewVector(5, 12)
var v2 = NewVector(-4, 8)

func TestVectorMath(t *testing.T) {
  if !v1.Clone().Add(v2).Compare(NewVector(1, 20)) {
    t.Error("vector addition should work")
  }
  if !v1.Clone().Sub(v2).Compare(NewVector(9, 4)) {
    t.Error("vector subtraction should work")
  }
}

func TestVectorRotationConversion(t *testing.T) {
  rot := 1.36
  res := NewVector(0, 0).FromRotation(rot, 1).ToRotation()
  if !(math.Floor(rot*1000) == math.Floor(res*1000)) {
    t.Error("vector rotation in should equal rotation out")
  }
}

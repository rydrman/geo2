package geo

import "testing"

func TestTriangleContains(t *testing.T) {
  tri := NewTriangle([]*Vector{
    NewVector(0, 0),
    NewVector(1, 0),
    NewVector(0.5, 1),
  })
  if false == tri.Contains(NewVector(0.5, 0.5)) {
    t.Error("triangle should contain point")
  }

  if true == tri.Contains(NewVector(1, 0.5)) {
    t.Error("triangle should not contain point")
  }

  if true == tri.Contains(NewVector(0, 1)) {
    t.Error("triangle should not contain point")
  }
}

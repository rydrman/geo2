package geo2

import "math"

// Rectangle represents a 2D rectangle
type Rectangle struct {
  X      float64
  Y      float64
  Width  float64
  Height float64
}

// NewRectangle creates a new rectangle from the given dimensions
func NewRectangle(x, y, w, h float64) *Rectangle {
  return &Rectangle{x, y, w, h}
}

// Set sets the dimensions of this rectangle
func (rect *Rectangle) Set(x, y, w, h float64) {
  rect.X = x
  rect.Y = y
  rect.Width = w
  rect.Height = h
}

// SetPosition sets the position of this rectangle
func (rect *Rectangle) SetPosition(x, y float64) {
  rect.X = x
  rect.Y = y
}

// SetPositionVector sets the position of this rectangle
// from the given Vector point
func (rect *Rectangle) SetPositionVector(pos *Vector) {
  rect.X = pos.X
  rect.Y = pos.Y
}

// SetSize sets the size of this rectangle
func (rect *Rectangle) SetSize(w, h float64) {
  rect.Width = w
  rect.Height = h
}

// SetSizeVector sets the size of this rectangle
// from the given Vector point
func (rect *Rectangle) SetSizeVector(size *Vector) {
  rect.Width = size.X
  rect.Height = size.Y
}

// Center returns the center point of this rectangle
func (rect *Rectangle) Center() *Vector {
  return NewVector(
    rect.X+rect.Width*0.5,
    rect.Y+rect.Height*0.5,
  )
}

// Edges returns the edges of this rectangle
// (starting at the top and going clockwise)
func (rect *Rectangle) Edges() []*Line {
  bottom := rect.Y + rect.Height
  right := rect.X + rect.Width
  return []*Line{
    NewLine(NewVector(rect.X, rect.Y), NewVector(right, rect.Y)),
    NewLine(NewVector(right, rect.Y), NewVector(right, bottom)),
    NewLine(NewVector(right, bottom), NewVector(rect.X, bottom)),
    NewLine(NewVector(rect.X, bottom), NewVector(rect.X, rect.Y)),
  }
}

// Contains returns true if this rectangle contains the given point
func (rect *Rectangle) Contains(vec *Vector) bool {
  return (vec.X >= rect.X &&
    vec.X <= rect.X+rect.Width &&
    vec.Y >= rect.Y &&
    vec.Y <= rect.Y+rect.Height)
}

// DistanceTo calculates the closest distance from the edges
// of this rectangle to the given point
func (rect *Rectangle) DistanceTo(vec *Vector) float64 {
  var min float64
  for _, edge := range rect.Edges() {
    min = math.Min(
      min,
      edge.DistanceToPoint(vec, true),
    )
  }

  return min
}

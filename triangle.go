package geo2

//Triangle represents a 2d triangle
type Triangle struct {
  Points []*Vector
}

//NewTriangle creates a new triangle
//object from the given points
func NewTriangle(points []*Vector) *Triangle {
  return &Triangle{
    points[0:3],
  }
}

//Contains returns true if the given point is within this triangle
func (t *Triangle) Contains(point *Vector) bool {
  cp1 := NewLine(t.Points[0], t.Points[1]).CrossWithPoint(point)
  cp2 := NewLine(t.Points[1], t.Points[2]).CrossWithPoint(point)
  cp3 := NewLine(t.Points[2], t.Points[0]).CrossWithPoint(point)

  return ((cp1 < 0 && cp2 < 0 && cp3 < 0) ||
    (cp1 >= 0 && cp2 >= 0 && cp3 >= 0))
}

//ToFloat32Array returns a float32
//array of the points in this triangle
func (t *Triangle) ToFloat32Array() []float32 {
  return []float32{
    float32(t.Points[0].X), float32(t.Points[0].Y),
    float32(t.Points[1].X), float32(t.Points[1].Y),
    float32(t.Points[2].X), float32(t.Points[2].Y),
  }
}

// TriangleList represents a slice of triangles
type TriangleList []*Triangle

// ToFloat32Array converts this list of triangles into
// a flat list of points [x1, y1, x2, y2, ... xn, yn]
func (tris *TriangleList) ToFloat32Array() []float32 {
  var floats []float32
  for _, tri := range *tris {
    floats = append(floats, tri.ToFloat32Array()...)
  }
  return floats
}

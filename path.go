package geo

import "math"

// Path2 represents a string of connected lines
// which may or may not represent a loop
type Path2 struct {
	Points []*Vector
}

// NewPath2FromFloat32Array creates a new path from the given array of
// float32 values (flat vector set [x1, y1, x2, y2, ... xn, yn])
func NewPath2FromFloat32Array(values []float32) *Path2 {
	path := &Path2{make([]*Vector, len(values)/2)}
	for i := range path.Points {
		path.Points[i] = NewVector(float64(values[i*2]), float64(values[i*2+1]))
	}
	return path
}

// ToFloat32Array converts this path into a flat array of
// float32 values [x1, y1, x2, y2, ... xn, yn]
func (path *Path2) ToFloat32Array() []float32 {
	values := make([]float32, len(path.Points)*2)
	for i, p := range path.Points {
		values[i*2] = float32(p.X)
		values[i*2+1] = float32(p.Y)
	}
	return values
}

// Clone creates a copy of this path by value
func (path *Path2) Clone() *Path2 {
	clone := &Path2{make([]*Vector, len(path.Points))}
	for i, p := range path.Points {
		clone.Points[i] = p.Clone()
	}
	return clone
}

// Append appends the given point to the end of this path
func (path *Path2) Append(vec *Vector) {
	path.Points = append(path.Points, vec)
}

// Triangulate triangulates this path into individual
// triangles representing the area covered by this path
//
// This process assumes that this path is a closed loop
// and triangulates the area within in. This function
// does not guarentee results for self-intersecting paths
func (path *Path2) Triangulate() *TriangleList {
	path = path.Clone()
	//triangulate the given polygon
	triangles := make(TriangleList, 0)
	for true {
		created := false
		for i := len(path.Points) - 1; i >= 0; i-- {
			vPrev := path.Points[(len(path.Points)+i-1)%len(path.Points)]
			vCurr := path.Points[i]
			vNext := path.Points[(i+1)%len(path.Points)]

			//make sure the triangle is concave
			//removed because it assumes a CC or CCW path direction
			//(now uses "center of third edge check", see below)
			/*if vPrev.Clone().Sub(vCurr).Cross(vNext.Clone().Sub(vCurr)) <= 0 {
			  continue
			}*/

			tri := NewTriangle([]*Vector{vPrev, vCurr, vNext})
			inside := false

			//if any of the other points are in this triangle, move on
			//also make sure that the center of the new edge is inside the
			//shape as an entirety
			edgeCenter := NewLine(vNext, vPrev).GetPosition(0.5)
			fromEdge := NewLine(
				NewVector(math.SmallestNonzeroFloat64, edgeCenter.Y),
				edgeCenter,
			)
			intersectionCount := 0
			for j, vert := range path.Points {
				//check if the center of this new edge is inside the shape
				vEnd := path.Points[(j+1)%len(path.Points)]
				if nil != NewLine(vert, vEnd).Intersection(fromEdge, true) {
					intersectionCount++
				}
				//dont check if the triangles verts are inside it :P
				if vert == vPrev || vert == vCurr || vert == vNext {
					continue
				}
				if tri.Contains(vert) {
					inside = true
				}
			}
			if inside || intersectionCount%2 == 0 {
				continue
			}

			//add this triangle to the list
			created = true
			triangles = append(triangles, tri)
			path.Points = append(path.Points[:i], path.Points[i+1:]...)
			break
		}
		if false == created || 0 == len(path.Points) {
			break
		}
	}
	return &triangles
}

package geo2

import "math"

// Path represents a string of connected lines
// which may or may not represent a loop
type Path []*Vector

// NewPathFromFloat32Array creates a new path from the given array of
// float32 values (flat vector set [x1, y1, x2, y2, ... xn, yn])
func NewPathFromFloat32Array(values []float32) *Path {
	path := (Path)(make([]*Vector, len(values)/2))
	for i := range path {
		path[i] = NewVector(float64(values[i*2]), float64(values[i*2+1]))
	}
	return &path
}

// ToFloat32Array converts this path into a flat array of
// float32 values [x1, y1, x2, y2, ... xn, yn]
func (path *Path) ToFloat32Array() []float32 {
	values := make([]float32, len(*path)*2)
	for i, p := range *path {
		values[i*2] = float32(p.X)
		values[i*2+1] = float32(p.Y)
	}
	return values
}

// Clone creates a copy of this path by value
func (path *Path) Clone() *Path {
	clone := (Path)(make([]*Vector, len(*path)))
	for i, point := range *path {
		clone[i] = point.Clone()
	}
	return &clone
}

// Append appends the given point to the end of this path
func (path *Path) Append(vec *Vector) {
	(*path) = append(*path, vec)
}

// Triangulate triangulates this path into individual
// triangles representing the area covered by this path
//
// This process assumes that this path is a closed loop
// and triangulates the area within in. This function
// does not guarentee results for self-intersecting paths
func (path *Path) Triangulate() *TriangleList {
	path = path.Clone()
	//triangulate the given polygon
	triangles := make(TriangleList, 0)
	for true {
		created := false
		for i := len(*path) - 1; i >= 0; i-- {
			vPrev := (*path)[(len(*path)+i-1)%len(*path)]
			vCurr := (*path)[i]
			vNext := (*path)[(i+1)%len(*path)]

			//make sure the triangle is concave
			//removed because it assumes a CC or CCW path direction
			//(now uses "center of third edge check", see below)
			/*if vPrev.Clone().Sub(vCurr).Cross(vNext.Clone().Sub(vCurr)) <= 0 {
			  continue
			}*/

			tri := NewTriangle([]*Vector{vPrev, vCurr, vNext})

			//if any of the other points are in this triangle, move on
			//also make sure that the center of the new edge is inside the
			//shape as an entirety
			edgeCenter := NewLine(vNext, vPrev).GetPosition(0.5)
			fromEdge := NewLine(
				NewVector(math.SmallestNonzeroFloat64, edgeCenter.Y),
				edgeCenter,
			)
			inside := false
			intersectionCount := 0
			for j, vert := range *path {
				//check if the center of this new edge is inside the shape
				vEnd := (*path)[(j+1)%len(*path)]
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
			(*path) = append((*path)[:i], (*path)[i+1:]...)
			break
		}
		if false == created || 0 == len(*path) {
			break
		}
	}
	return &triangles
}

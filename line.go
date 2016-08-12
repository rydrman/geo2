package geo

import "math"

//Line represents a 2d line
type Line struct {
	A *Vector
	B *Vector
}

//NewLine creates a new Line
//object from the given points
func NewLine(pointA, pointB *Vector) *Line {
	return &Line{pointA, pointB}
}

//ToVector returns the vector of line.A -> line.B
func (line *Line) ToVector() *Vector {
	return line.B.Clone().Sub(line.A)
}

// LengthSqd returns the squared length of line line
func (line *Line) LengthSqd() float64 {
	return line.ToVector().LengthSqd()
}

// Length returns the lenght of line line
func (line *Line) Length() float64 {
	return line.ToVector().Length()
}

// GetPosition gets the position along this line
//
// perc = 0 would be point A, perc of 1 would be point B
// but the passed perc can be any number (not just 0-1)
func (line *Line) GetPosition(perc float64) *Vector {
	return NewVector(
		line.A.X+(line.B.X-line.A.X)*perc,
		line.A.Y+(line.B.Y-line.A.Y)*perc,
	)
}

// GetPerc gets the percentage along this line of the given point
//
// this point is assumed to be on the line, otherwise
// the result is undeterministic
func (line *Line) GetPerc(point *Vector) float64 {
	return (point.X - line.A.X) / (line.B.X - line.A.X)
}

//CrossWithPoint performs the cross product of line line
//with the line going from line line to a given point
//to figure out which side of the line the point is on
func (line *Line) CrossWithPoint(v *Vector) float64 {
	return line.ToVector().Cross(v.Clone().Sub(line.A))
}

// DistanceToPoint returns the distance from this line to the
// given point. Use clamp to ensure that the distance is measured
// from a point between A and B rather than anywhere on the line
// defined by A and B
func (line *Line) DistanceToPoint(point *Vector, clamp bool) float64 {
	result := line.ClosestPoint(point, clamp)
	return result.Sub(point).Length()
}

// ClosestPoint returns the closest point on this line to the
// given point. Use clamp to ensure that the point returned is
// a point between A and B rather than anywhere on the line
// defined by A and B
func (line *Line) ClosestPoint(point *Vector, clamp bool) *Vector {
	ap := point.Clone().Sub(line.A)
	ab := point.Clone().Sub(line.B)

	var perc = (ap.X*ab.X + ap.Y*ab.Y) / ab.LengthSqd()

	if clamp {
		perc = math.Min(1, math.Max(0, perc))
	}

	return line.GetPosition(perc)
}

// Intersection returns the intersection between this line and
// another. Returns nil if there is no intersection. Use clamp
// to specify that the point should exist between A and B on each line
func (line *Line) Intersection(Line *Line, clamp bool) *Vector {
	det := (line.A.X-line.B.X)*(Line.A.Y-Line.B.Y) - (line.A.Y-line.B.Y)*(Line.A.X-Line.B.X)
	if 0 == det {
		return nil
	}
	point := &Vector{
		((line.A.X*line.B.Y-line.A.Y*line.B.X)*(Line.A.X-Line.B.X) -
			(line.A.X-line.B.X)*(Line.A.X*Line.B.Y-Line.A.Y*Line.B.X)) / det,
		((line.A.X*line.B.Y-line.A.Y*line.B.X)*(Line.A.Y-Line.B.Y) -
			(line.A.Y-line.B.Y)*(Line.A.X*Line.B.Y-Line.A.Y*Line.B.X)) / det,
	}
	if clamp {
		perc1 := line.GetPerc(point)
		perc2 := Line.GetPerc(point)
		if perc1 < 0 || perc1 > 1 ||
			perc2 < 0 || perc2 > 1 {
			return nil
		}
	}
	return point
}

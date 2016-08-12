package geo2

import (
  "fmt"
  "math"
)

//Vector represents a 2-dimentional vector
type Vector struct {
  X float64
  Y float64
}

//NewVector creates a new vector object with the
//given x and y components
func NewVector(x float64, y float64) *Vector {
  return new(Vector).Set(x, y)
}

//Set sets the values of this vector
func (v *Vector) Set(x float64, y float64) *Vector {
  v.X = x
  v.Y = y
  return v
}

//SetX sets the x value of this vector
func (v *Vector) SetX(x float64) *Vector {
  v.X = x
  return v
}

//SetY sets the y value of this vector
func (v *Vector) SetY(y float64) *Vector {
  v.Y = y
  return v
}

//Copy copy the value of another vector into this one
func (v *Vector) Copy(v2 *Vector) *Vector {
  v.X = v2.X
  v.Y = v2.Y
  return v
}

//Clone creates a copy of this vector
func (v *Vector) Clone() *Vector {
  return NewVector(v.X, v.Y)
}

//Add add a vector to this one
func (v *Vector) Add(v2 *Vector) *Vector {
  v.X += v2.X
  v.Y += v2.Y
  return v
}

//AddVectors set this vector to the addition of the given two vector
func (v *Vector) AddVectors(v1 *Vector, v2 *Vector) *Vector {
  v.X = v1.X + v2.X
  v.Y = v1.Y + v2.Y
  return v
}

//Sub subtract a vector from this vector
func (v *Vector) Sub(v2 *Vector) *Vector {
  v.X -= v2.X
  v.Y -= v2.Y
  return v
}

//SubVectors set the value of this vector to the subtraction of the given two vectors
func (v *Vector) SubVectors(v1 *Vector, v2 *Vector) *Vector {
  v.X = v1.X - v2.X
  v.Y = v1.Y - v2.Y
  return v
}

//Multiply multiplies this vector by another
func (v *Vector) Multiply(v2 *Vector) *Vector {
  v.X *= v2.X
  v.Y *= v2.Y
  return v
}

//MultiplyMatrix multiplies this vector by the given matrix
func (v *Vector) MultiplyMatrix(m *Matrix3) *Vector {
  x := v.X
  v.X = x*m[0][0] + v.Y*m[0][1] + m[0][2]
  v.Y = x*m[1][0] + v.Y*m[1][1] + m[1][2]
  return v
}

//MultiplyScalar multiplies this vector by the given scalar value
func (v *Vector) MultiplyScalar(value float64) *Vector {
  v.X *= value
  v.Y *= value
  return v
}

//Divide divides this vector by another
func (v *Vector) Divide(v2 *Vector) *Vector {
  v.X /= v2.X
  v.Y /= v2.Y
  return v
}

//DivideScalar divides this vector by the given scalar value
func (v *Vector) DivideScalar(s float64) *Vector {
  v.X /= s
  v.Y /= s
  return v
}

//Normalize normalizes this vector to a length of 1
func (v *Vector) Normalize() *Vector {
  l := v.Length()
  v.X /= l
  v.Y /= l
  return v
}

//Cross find the cross product of this
//vector and the given vector
func (v *Vector) Cross(v2 *Vector) float64 {
  return v.X*v2.Y - v.Y*v2.X
}

//Dot find the dot product of this
//vector and the given vector
func (v *Vector) Dot(v2 *Vector) float64 {
  return v.X*v2.Y + v.Y*v2.X
}

//Negate negate this vector (make it it's exact opposite)
func (v *Vector) Negate() *Vector {
  return v.MultiplyScalar(-1)
}

//Length return the length of this vector, if simple
//comparisons are needed its more efficient to use
//the LengthSqd() value
func (v *Vector) Length() float64 {
  return math.Sqrt(v.LengthSqd())
}

//LengthSqd returns the squared value of the length of this vector
func (v *Vector) LengthSqd() float64 {
  return v.X*v.X + v.Y*v.Y
}

//Limit if the length of this vector is greater than the given maximum, shrink it
func (v *Vector) Limit(max float64) *Vector {
  if v.Length() > max {
    v.Normalize().MultiplyScalar(max)
  }
  return v
}

//Clamp similar to Limit() but works on a max and a min
func (v *Vector) Clamp(max float64, min float64) *Vector {
  l := v.Length()
  if l > max {
    v.Normalize().MultiplyScalar(max)
  } else if l < min {
    v.Normalize().MultiplyScalar(min)
  }
  return v
}

//FromRotation set this vector to the vector which points
//to the given rotation value at the given radius distance
func (v *Vector) FromRotation(rotation float64, radius float64) *Vector {
  v.X = math.Cos(rotation) * radius
  v.Y = math.Sin(rotation) * radius
  return v
}

//ToRotation converts this vector to a rotation value
func (v *Vector) ToRotation() float64 {
  return math.Atan2(v.Y, v.X)
}

//String outputs a string representation of the given vector
func (v *Vector) String() string {
  return fmt.Sprintf("{x: %.4f, y: %.4f}", v.X, v.Y)
}

//Compare compare two vectors to see if they are exactly equal
func (v *Vector) Compare(v2 *Vector) bool {
  return v.X == v2.X && v.Y == v2.Y
}

// CloseEnough compares two vectors to see if they are equal
// to a given number of significant digits
// (for example 0.01, would mean that  1.45478 =  1.45229)
// (for example 10.0, would mean that 22.45478 = 24.26783)
func (v *Vector) CloseEnough(v2 *Vector, sigDigits float64) bool {
  factor := 1.0 / sigDigits
  lhs := v.Clone().MultiplyScalar(factor).Add(NewVector(0.5, 0.5))
  lhs.X = math.Floor(lhs.X)
  lhs.Y = math.Floor(lhs.Y)
  rhs := v2.Clone().MultiplyScalar(factor).Add(NewVector(0.5, 0.5))
  rhs.X = math.Floor(rhs.X)
  rhs.Y = math.Floor(rhs.Y)
  return lhs.Compare(rhs)
}

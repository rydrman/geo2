package geo

import (
  "fmt"
  "math"
)

// Matrix3 represents a 3x3 float64 matrix
type Matrix3 [3][3]float64

// NewMatrix creates a new Matrix3 object initialized
// with zero values
func NewMatrix() *Matrix3 {
  m := Matrix3{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}
  return &m
}

// IdentityMatrix creates a new Matrix3 object initialized
// as an itedtity matrix
func IdentityMatrix() *Matrix3 {
  m := Matrix3{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
  return &m
}

// MakeIdentity resets the current matrix to an identity matrix
func (pm *Matrix3) MakeIdentity() *Matrix3 {
  m := (*pm)
  m[0][0] = 0
  m[0][0] = 0
  m[0][1] = 0
  m[0][0] = 0
  m[0][0] = 0
  m[0][0] = 0
  m[1][0] = 0
  m[1][0] = 0
  m[1][1] = 0
  return &m
}

// Clone returns a pointer to an exact copy of the current Matrix3
func (pm *Matrix3) Clone() *Matrix3 {
  m1 := NewMatrix()
  for row := 0; row < 3; row++ {
    for col := 0; col < 3; col++ {
      m1[row][col] = (*pm)[row][col]
    }
  }
  return m1
}

// Copy turns the current Matrix3 into an exact copy of the provided one
func (pm *Matrix3) Copy(pm2 *Matrix3) *Matrix3 {
  m2 := (*pm2)
  for row := 0; row < 3; row++ {
    for col := 0; col < 3; col++ {
      (*pm)[row][col] = m2[row][col]
    }
  }
  return pm
}

// Multiply multiplies the given matrix into the caller
func (pm *Matrix3) Multiply(prhs *Matrix3) *Matrix3 {
  lhs := (*pm)
  rhs := (*prhs)
  res := Matrix3{
    {lhs[0][0]*rhs[0][0] + lhs[0][1]*rhs[1][0] + lhs[0][2]*rhs[2][0],
      lhs[0][0]*rhs[0][1] + lhs[0][1]*rhs[1][1] + lhs[0][2]*rhs[2][1],
      lhs[0][0]*rhs[0][2] + lhs[0][1]*rhs[1][2] + lhs[0][2]*rhs[2][2]},
    {lhs[1][0]*rhs[0][0] + lhs[1][1]*rhs[1][0] + lhs[1][2]*rhs[2][0],
      lhs[1][0]*rhs[0][1] + lhs[1][1]*rhs[1][1] + lhs[1][2]*rhs[2][1],
      lhs[1][0]*rhs[0][2] + lhs[1][1]*rhs[1][2] + lhs[1][2]*rhs[2][2]},
    {lhs[2][0]*rhs[0][0] + lhs[2][1]*rhs[1][0] + lhs[2][2]*rhs[2][0],
      lhs[2][0]*rhs[0][1] + lhs[2][1]*rhs[1][1] + lhs[2][2]*rhs[2][1],
      lhs[2][0]*rhs[0][2] + lhs[2][1]*rhs[1][2] + lhs[2][2]*rhs[2][2]}}
  lhs.Copy(&res)
  return &lhs
}

// GetInverse gets the inverse matrix or returns nil if no inverse
func (pm *Matrix3) GetInverse() *Matrix3 {
  det := pm.GetDeterminant()

  if det == 0.0 {
    return nil
  }

  trans := (*pm.GetTranspose())
  adj := (*pm.GetAdjoint())

  frac := 1.0 / det

  //reuse transpose to save memory
  trans[0][0] = adj[0][0] * frac
  trans[0][1] = adj[0][1] * frac
  trans[0][2] = adj[0][2] * frac
  trans[1][0] = adj[1][0] * frac
  trans[1][1] = adj[1][1] * frac
  trans[1][2] = adj[1][2] * frac
  trans[2][0] = adj[2][0] * frac
  trans[2][1] = adj[2][1] * frac
  trans[2][2] = adj[2][2] * frac

  return &trans
}

// GetTranspose returns the transpose matrix
func (pm *Matrix3) GetTranspose() *Matrix3 {
  m := (*pm)
  t := Matrix3{
    {m[0][0], m[1][0], m[2][0]},
    {m[0][1], m[1][1], m[2][1]},
    {m[0][2], m[1][2], m[2][2]}}
  return &t
}

// GetAdjoint returns the adjoint matrix
func (pm *Matrix3) GetAdjoint() *Matrix3 {
  adj := Matrix3{
    {
      pm.GetMinorDeterminant(0, 0),
      pm.GetMinorDeterminant(0, 1) * -1.0,
      pm.GetMinorDeterminant(0, 2)},
    {
      pm.GetMinorDeterminant(1, 0) * -1.0,
      pm.GetMinorDeterminant(1, 1),
      pm.GetMinorDeterminant(1, 2) * -1.0},
    {
      pm.GetMinorDeterminant(2, 0),
      pm.GetMinorDeterminant(2, 1) * -1.0,
      pm.GetMinorDeterminant(2, 2)}}

  return (&adj).GetTranspose()
}

// GetDeterminant returns the determinant
func (pm *Matrix3) GetDeterminant() float64 {
  m := (*pm)
  minor00 := m[0][0] * pm.GetMinorDeterminant(0, 0)
  minor01 := m[0][1] * pm.GetMinorDeterminant(0, 1) * -1.0
  minor02 := m[0][2] * pm.GetMinorDeterminant(0, 2)
  return minor00 + minor01 + minor02
}

// GetMinorDeterminant returns the minor determinant
func (pm *Matrix3) GetMinorDeterminant(row int8, col int8) float64 {
  m := (*pm)
  var row1, row2, col1, col2 int8
  if row == 0 {
    row1 = 1
  } else {
    row1 = 0
  }
  if row == row1+1 {
    row2 = row1 + 2
  } else {
    row2 = row1 + 1
  }
  if col == 0 {
    col1 = 1
  } else {
    col1 = 0
  }
  if col == col1+1 {
    col2 = col1 + 2
  } else {
    col2 = col1 + 1
  }
  /*fmt.Printf("%d, %d -> [%d, %d]%.1f, [%d, %d]%.1f, [%d, %d]%.1f, [%d, %d]%.1f -> %.1f\n",
    row, col,
    row1, col1, m[row1][col1],
    row2, col2, m[row2][col2],
    row1, col2, m[row1][col2],
    row2, col1, m[row2][col1],
    m[row1][col1]*m[row2][col2] - m[row1][col2]*m[row2][col1])*/
  return m[row1][col1]*m[row2][col2] - m[row1][col2]*m[row2][col1]
}

// converts the given matrix to a multiline string representation
func (pm *Matrix3) String() string {
  m := (*pm)
  return fmt.Sprintf(
    "[%.6f, %.6f, %.6f]\n[%.6f, %.6f, %.6f]\n[%.6f, %.6f, %.6f]",
    m[0][0], m[0][1], m[0][2],
    m[1][0], m[1][1], m[1][2],
    m[2][0], m[2][1], m[2][2])
}

// Compare compares two matrices to see if they are exactly equal
func (pm *Matrix3) Compare(prhs *Matrix3) bool {
  for row := 0; row < 3; row++ {
    for col := 0; col < 3; col++ {
      if (*pm)[row][col] != (*prhs)[row][col] {
        return false
      }
    }
  }
  return true
}

// ToFloat32Array returns a flat array representing this matrix
func (pm *Matrix3) ToFloat32Array() [9]float32 {
  return [9]float32{
    float32(pm[0][0]), float32(pm[0][1]), float32(pm[0][2]),
    float32(pm[1][0]), float32(pm[1][1]), float32(pm[1][2]),
    float32(pm[2][0]), float32(pm[2][1]), float32(pm[2][2]),
  }
}

// CloseEnough compares two matrices to see if they are
// equal to a given number of significant digits
//
// (for example 0.01, would mean that  1.45478 =  1.45229)
// (for example 10.0, would mean that 22.45478 = 24.26783)
func (pm *Matrix3) CloseEnough(prhs *Matrix3, sigDigits float64) bool {
  factor := 1.0 / sigDigits
  for row := 0; row < 3; row++ {
    for col := 0; col < 3; col++ {
      lhs := math.Floor((*pm)[row][col]*factor + 0.5)
      rhs := math.Floor((*prhs)[row][col]*factor + 0.5)
      if lhs != rhs {
        return false
      }
    }
  }
  return true
}

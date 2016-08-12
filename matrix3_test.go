package geo

import "testing"

var input Matrix3 = Matrix3{
  {14, 9, 3},
  {46, 12, 99},
  {49, 38, 4}}
var pInput *Matrix3 = &input

func TestMatrixCompare(t *testing.T) {
  if !IdentityMatrix().Compare(IdentityMatrix()) {
    t.Error("identical matrices should compare to be equal")
  }
}

func TestMatrixCopy(t *testing.T) {
  if !NewMatrix().Copy(IdentityMatrix()).Compare(IdentityMatrix()) {
    t.Error("matrices should copy exactly")
  }
}

func TestMatrixMultiplication(t *testing.T) {
  result := IdentityMatrix().Multiply(IdentityMatrix())
  if !result.Compare(IdentityMatrix()) {
    t.Error("identity * identity should = identity")
  }
  rhs := Matrix3{
    {8, 56, -16},
    {84, -28, 81},
    {-71, 13, 9}}
  expected := Matrix3{
    {655, 571, 532},
    {-5653, 3527, 1127},
    {3300, 1732, 2330}}
  if !pInput.Multiply(&rhs).Compare(&expected) {
    t.Error("matrix multiplication should return correctly")
  }
}

func TestMatrixMinorDeterminants(t *testing.T) {
  if -3714 != pInput.GetMinorDeterminant(0, 0) ||
    -4667 != pInput.GetMinorDeterminant(0, 1) ||
    1160 != pInput.GetMinorDeterminant(0, 2) ||
    -78 != pInput.GetMinorDeterminant(1, 0) ||
    -91 != pInput.GetMinorDeterminant(1, 1) ||
    91 != pInput.GetMinorDeterminant(1, 2) ||
    855 != pInput.GetMinorDeterminant(2, 0) ||
    1248 != pInput.GetMinorDeterminant(2, 1) ||
    -246 != pInput.GetMinorDeterminant(2, 2) {
    t.Error("minor determinent should calculate properly")
  }
}

func TestMatrixDeterminant(t *testing.T) {
  if -6513 != pInput.GetDeterminant() {
    t.Error("determinant should calculate properly")
  }
  if 1 != IdentityMatrix().GetDeterminant() {
    t.Error("determinant of identity should be 1")
  }
}

func TestTranspose(t *testing.T) {
  expected := Matrix3{
    {14, 46, 49},
    {9, 12, 38},
    {3, 99, 4}}
  if !pInput.GetTranspose().Compare(&expected) {
    t.Error("matrix transpose should return correctly")
  }
  if !IdentityMatrix().GetTranspose().Compare(IdentityMatrix()) {
    t.Error("transpose of identity matrix should be identity matrix")
  }
}

func TestAdjoint(t *testing.T) {
  expected := Matrix3{
    {-3714, 78, 855},
    {4667, -91, -1248},
    {1160, -91, -246}}
  if !pInput.GetAdjoint().CloseEnough(&expected, 0.001) {
    t.Error("matrix adjoint should return correctly")
  }
  if !IdentityMatrix().GetAdjoint().Compare(IdentityMatrix()) {
    t.Error("adjoint of identity matrix should be identity matrix")
  }
}

func TestMatrixInversion(t *testing.T) {
  det := pInput.GetDeterminant()
  expected := Matrix3{
    {-3714 / det, 78 / det, 855 / det},
    {4667 / det, -91 / det, -1248 / det},
    {1160 / det, -91 / det, -246 / det}}
  result := pInput.GetInverse()
  if !result.CloseEnough(&expected, 0.001) {
    t.Error("matrix inverse should return correctly")
  }
  result = pInput.Multiply(result)
  if !result.CloseEnough(IdentityMatrix(), 0.01) {
    t.Error("inverse times original should equal identity")
  }
  result = IdentityMatrix().GetInverse()
  if !result.CloseEnough(IdentityMatrix(), 0.001) {
    t.Error("inverse of identity should be identity")
  }
}

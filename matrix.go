package glm

import (
	"fmt"
	"math"
)

type Matrix4 [16]float32

// Construct the identity 4x4 matrix
func Matrix4Identity() Matrix4 {
	return Matrix4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1}
}

func MatrixMultiply(m1, m2 Matrix4) Matrix4 {
	out := Matrix4Identity()
	out[0] = Dot4(m1.Row(0), m2.Column(0))
	out[1] = Dot4(m1.Row(1), m2.Column(0))
	out[2] = Dot4(m1.Row(2), m2.Column(0))
	out[3] = Dot4(m1.Row(3), m2.Column(0))
	out[4] = Dot4(m1.Row(0), m2.Column(1))
	out[5] = Dot4(m1.Row(1), m2.Column(1))
	out[6] = Dot4(m1.Row(2), m2.Column(1))
	out[7] = Dot4(m1.Row(3), m2.Column(1))
	out[8] = Dot4(m1.Row(0), m2.Column(2))
	out[9] = Dot4(m1.Row(1), m2.Column(2))
	out[10] = Dot4(m1.Row(2), m2.Column(2))
	out[11] = Dot4(m1.Row(3), m2.Column(2))
	out[12] = Dot4(m1.Row(0), m2.Column(3))
	out[13] = Dot4(m1.Row(1), m2.Column(3))
	out[14] = Dot4(m1.Row(2), m2.Column(3))
	out[15] = Dot4(m1.Row(3), m2.Column(3))
	return out
}

// Construct a translation matrix
func MatrixTranslate(x, y, z float32) Matrix4 {
	return Matrix4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		x, y, z, 1}
}

// Construct a scaling matrix
func MatrixScale(x, y, z float32) Matrix4 {
	return Matrix4{
		x, 0, 0, 0,
		0, y, 0, 0,
		0, 0, z, 0,
		0, 0, 0, 1}
}

// Construct a perspective matrix
func MatrixPerspective(fov, aspect, near, far float32) Matrix4 {

	f := 1.0 / float32(math.Tan(float64(fov/2)))
	m10 := (near + far) / (near - far)
	w := (2 * near * far) / (near - far)

	return Matrix4{
		f / aspect, 0, 0, 0,
		0, f, 0, 0,
		0, 0, m10, -1,
		0, 0, w, 1}
}

func (m Matrix4) Row(n int) Vector4 {
	return Vector4{m[0+n], m[4+n], m[8+n], m[12+n]}
}

func (m Matrix4) Column(n int) Vector4 {
	return Vector4{m[0+n*4], m[1+n*4], m[2+n*4], m[3+n*4]}
}

//calculate the determinant
func (m Matrix4) Determinant() float32 {
	m00 := m[0]
	m01 := m[1]
	m02 := m[2]
	m03 := m[3]
	m10 := m[4]
	m11 := m[5]
	m12 := m[6]
	m13 := m[7]
	m20 := m[8]
	m21 := m[9]
	m22 := m[10]
	m23 := m[11]
	m30 := m[12]
	m31 := m[13]
	m32 := m[14]
	m33 := m[15]

	det := m03*m12*m21*m30 - m02*m13*m21*m30 -
		m03*m11*m22*m30 + m01*m13*m22*m30 +
		m02*m11*m23*m30 - m01*m12*m23*m30 -
		m03*m12*m20*m31 + m02*m13*m20*m31 +
		m03*m10*m22*m31 - m00*m13*m22*m31 -
		m02*m10*m23*m31 + m00*m12*m23*m31 +
		m03*m11*m20*m32 - m01*m13*m20*m32 -
		m03*m10*m21*m32 + m00*m13*m21*m32 +
		m01*m10*m23*m32 - m00*m11*m23*m32 -
		m02*m11*m20*m33 + m01*m12*m20*m33 +
		m02*m10*m21*m33 - m00*m12*m21*m33 -
		m01*m10*m22*m33 + m00*m11*m22*m33

	return det
}

func (m Matrix4) String() string {
	return fmt.Sprintf("%v\n%v\n%v\n%v", m.Row(0), m.Row(1), m.Row(2), m.Row(3))
}

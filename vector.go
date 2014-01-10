package glm

import "math"

type Vector3 [3]float32
type Vector4 [4]float32

// vector dot product
func Dot(v1, v2 Vector3) float32 {
	return v1[0]*v2[0] + v1[1]*v2[1] + v1[2]*v2[2]
}

func Dot4(v1, v2 Vector4) float32 {
	return v1[0]*v2[0] + v1[1]*v2[1] + v1[2]*v2[2] + v1[3]*v2[3]
}

// vector cross product
func Cross(v1, v2 Vector3) Vector3 {
	xh := v1[1]*v2[2] - v2[1]*v1[2]
	yh := v1[2]*v2[0] - v2[2]*v1[0]
	zh := v1[0]*v2[1] - v2[0]*v1[1]
	return Vector3{xh, yh, zh}
}

// vector scalar product
func ScalarProduct(v Vector3, s float32) Vector3 {
	return Vector3{v[0] * s, v[1] * s, v[2] * s}
}

// normalize vector
func Normalize(v Vector3) Vector3 {
	t := math.Sqrt(float64(v[0]*v[0] + v[1]*v[1] + v[2]*v[2]))
	return Vector3{v[0] / float32(t), v[1] / float32(t), v[2] / float32(t)}
}

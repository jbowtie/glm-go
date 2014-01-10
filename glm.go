package glm

import (
	"container/list"
)

//linear interpolation
func Mix(x, y, a float32) float32 {
	return x*(1-a) + y*a
}

type MatrixStack struct {
	current Matrix4
	stack   *list.List
}

func NewMatrixStack() *MatrixStack {
	return &MatrixStack{Matrix4Identity(), list.New()}
}

func (s *MatrixStack) Top() Matrix4 {
	return s.current
}

func (s *MatrixStack) Push() {
	s.stack.PushBack(s.current)
}

func (s *MatrixStack) Pop() {
	e := s.stack.Back()
	s.stack.Remove(e)
	s.current = e.Value.(Matrix4)
}

func (s *MatrixStack) Perspective(fov, aspect, near, far float32) {
	s.current = MatrixMultiply(s.current, MatrixPerspective(fov, aspect, near, far))
}
func (s *MatrixStack) Translate(x, y, z float32) {
	s.current = MatrixMultiply(s.current, MatrixTranslate(x, y, z))
}
func (s *MatrixStack) Scale(x, y, z float32) {
	s.current = MatrixMultiply(s.current, MatrixScale(x, y, z))
}

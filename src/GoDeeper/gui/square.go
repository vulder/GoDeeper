package gui

import "github.com/go-gl/gl/v2.1/gl"

type Square struct {
	P1, P2, P3, P4 Point
	color rgb
}

// TODO: please merge me
// TODO: this should be (left, top int)
func NewSquareTS(top, left int, color rgb) *Square {
	s := new(Square)
	s.P1 = Point{left, top}
	s.P2 = Point{left + tile_size, top}
	s.P3 = Point{left + tile_size, top + tile_size}
	s.P4 = Point{left, top + tile_size}
	s.color = color
	return s
}

// TODO: this should be (left, top int)
func NewSquare(top, left, w, h int, color rgb) *Square {
	s := new(Square)
	s.P1 = Point{left, top}
	s.P2 = Point{left + w, top}
	s.P3 = Point{left + w, top + h}
	s.P4 = Point{left, top + h}
	s.color = color
	return s
}

var counter = 0

func (s *Square) Draw() {
	s.color.SetGLColor()

	gl.Begin(gl.POLYGON)
	drawVecPoint(s.P1)
	drawVecPoint(s.P2)
	drawVecPoint(s.P3)
	drawVecPoint(s.P4)
	gl.End()
}

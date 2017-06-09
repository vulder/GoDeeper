package gui

import "github.com/go-gl/gl/v2.1/gl"

type Square struct {
	P1, P2, P3, P4 Point
}

func (s *Square) Draw() {
	gl.Begin(gl.POLYGON)
	drawVecPoint(s.P1)
	drawVecPoint(s.P2)
	drawVecPoint(s.P3)
	drawVecPoint(s.P4)
	gl.End()
}

package gui

import "github.com/go-gl/gl/v2.1/gl"

type Square struct {
	P1, P2, P3, P4 Point
}

var counter = 0

func (s *Square) Draw() {
	gl.ClearColor(255,255,255,0)
	gl.Color3f(0,1,0)

	gl.Begin(gl.POLYGON)
	drawVecPoint(s.P1)
	drawVecPoint(s.P2)
	drawVecPoint(s.P3)
	drawVecPoint(s.P4)
	gl.End()
}

package gui

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
)

func DrawScene(window *glfw.Window) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	drawFoo()
	window.SwapBuffers()
}

func drawFoo() {
	s := Square{Point{1,1},Point{11,1}, Point{11,11}, Point{1, 11}}
	s.Draw()
}

/* utils */
func drawVecPoint(p Point) {
	gl.Vertex2i(int32(p.X), int32(p.Y))
}

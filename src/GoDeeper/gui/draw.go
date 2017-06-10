package gui

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
)

const (
	w_tiles = 42
	h_tiles = 42
)

func GetWidth() int {
	return w_tiles * 10
}

func GetHigh() int {
	return h_tiles * 10
}

type context struct {
	Window *glfw.Window
}

var currentContext = context{}

func (c context) GetWidthScale() int {
	width, _ := c.Window.GetSize()
	return width / w_tiles
}

func (c context) GetHeightScale() int {
	_, height := c.Window.GetSize()
	return height / h_tiles
}

func DrawScene(window *glfw.Window, w int32, h int32) {
	currentContext.Window = window

	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	drawFoo()

	gl.ClearColor(255, 255, 255, 0)
	gl.Color3f(1, 0, 0)

	gl.Begin(gl.POLYGON)
	gl.Vertex2i(0, 0)
	gl.Vertex2i(w, 0)
	gl.Vertex2i(w, h)
	gl.Vertex2i(0, h)
	gl.End()

	window.SwapBuffers()
}

func drawFoo() {
	s := Square{Point{0, 0}, Point{10, 0}, Point{10, 10}, Point{0, 10}}
	s.Draw()
}

/* utils */
func drawVecPoint(p Point) {
	gl.Vertex2i(int32(p.X), int32(p.Y))
}

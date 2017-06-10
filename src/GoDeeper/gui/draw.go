package gui

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.1/glfw"

	"GoDeeper/game"
)

const (
	w_tiles = 200
	h_tiles = 100
	tile_size = 42
)

func GetWidth() int {
	return w_tiles * tile_size
}

func GetHigh() int {
	return h_tiles * tile_size
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

	drawBoard()

	window.SwapBuffers()
}

func drawBoard() {
	board := game.GetBoard()
	for i := 0; i < 200; i++ {
		for j := 0; j < 100; j++ {
			switch board.GetCell(i, j) {
			case game.Earth:
				DrawEarth(coordsToScreen(i, j))
			case game.Pipe:
				DrawPipe(coordsToScreen(i, j))
			case game.Power:
				DrawPower(coordsToScreen(i, j))
			case game.Water:
				DrawWater(coordsToScreen(i, j))
			}
			DrawErr(i * tile_size,j * tile_size)
		}
	}
}

/* utils */
func coordsToScreen(x, y int) (sx, sy int){
	sx = x * tile_size
	sy = y * tile_size
	return
}

func drawVecPoint(p Point) {
	gl.Vertex2i(int32(p.X), int32(p.Y))
}

type rgb struct {
	r,g,b uint8
}

func (c rgb) SetGLColor() {
	gl.ClearColor(255,255,255,0)
	gl.Color3ub(c.r, c.g, c.b)
}

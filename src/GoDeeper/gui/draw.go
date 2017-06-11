package gui

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.1/glfw"

	"GoDeeper/game"
)

const (
	tile_size = 10
)

func GetWidth() int {
	return game.BOARD_WIDTH * tile_size
}

func GetHigh() int {
	return game.BOARD_HEIGHT * tile_size
}

type context struct {
	Window *glfw.Window
}

var currentContext = context{}

func (c context) GetWidthScale() int {
	width, _ := c.Window.GetSize()
	return width / game.BOARD_WIDTH
}

func (c context) GetHeightScale() int {
	_, height := c.Window.GetSize()
	return height / game.BOARD_HEIGHT
}

func DrawScene(window *glfw.Window) {
	currentContext.Window = window

	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	drawBG()

	drawBoard()

	DrawWater(0,0)

	window.SwapBuffers()
}

func drawBG () {
	for i := 0; i < game.BOARD_HEIGHT; i++ {
		for j := 0; j < game.BOARD_WIDTH; j++ {
			sx, sy := coordsToScreen(i, j)
			DrawErr(sx, sy)
		}
	}
}

func drawBoard() {
	board := game.GetBoard()
	for i := 0; i < game.BOARD_HEIGHT; i++ {
		for j := 0; j < game.BOARD_WIDTH; j++ {
			sx, sy := coordsToScreen(i, j)

			switch board.GetCell(i, j) {
			case game.Earth:
				DrawEarth(sx, sy)
				break
			case game.Tunnel:
				DrawTunnel(sx, sy)
				break
			case game.Gopher:
				DrawGopher(sx, sy)
				break
			case game.GopherSuperPower:
				DrawGopherSuperPower(sx,sy)
				break
			case game.Pipe:
				DrawPipe(sx, sy)
				break
			case game.Power:
				DrawPower(sx, sy)
				break
			case game.Water:
				DrawWater(sx, sy)
				break
			case game.Enemy:
				DrawEnemy(sx, sy)
				break
			case game.SuperPowerFood:
				DrawSuperPowerFood(sx, sy)
        break
      case game.Wormhole:
				DrawWormwhole(sx, sy)
				break
			}
			//DrawErr(sx, sy)
		}
	}
}

/* utils */
func coordsToScreen(x, y int) (sx, sy int) {
	sx = x * tile_size
	sy = y * tile_size
	return
}

func drawVecPoint(p Point) {
	gl.Vertex2i(int32(p.X), int32(p.Y))
}

type rgb struct {
	r, g, b uint8
}

func (c rgb) SetGLColor() {
	gl.ClearColor(255, 255, 255, 0)
	gl.Color3ub(c.r, c.g, c.b)
}

func num2PixelArray(n int) *[][]int {
	var nDigits = n%10 + 1
	var result [][]int = make([][]int, 5)
	for i := 0; i < 5; i++ {
		result[i] = make([]int, nDigits * 5)
	}

	var currDigitPos = nDigits

	for n > 0 {
		var currArr *[][]int = digit2PixelArray(n % 10)
		var xOffset int = (currDigitPos - 1) * 5
		for i:= 0; i < len(*currArr); i++ {
			for j:= 0; j < len((*currArr)[i]); j++ {
				result[i][j + xOffset] = (*currArr)[i][j]
			}
		}
		n /= 10
		currDigitPos -= 1
	}
	return &result
}

func digit2PixelArray(digit int) *[][]int {
	switch digit {
	case 0:
		return &[][]int{
			[]int{0, 1, 1, 1, 0},
			[]int{0, 1, 0, 1, 0},
			[]int{0, 1, 0, 1, 0},
			[]int{0, 1, 0, 1, 0},
			[]int{0, 1, 1, 1, 0},
		}
	case 1:
		return &[][]int{
			[]int{0, 0, 1, 0, 0},
			[]int{0, 0, 1, 0, 0},
			[]int{0, 0, 1, 0, 0},
			[]int{0, 0, 1, 0, 0},
			[]int{0, 0, 1, 0, 0},
		}
	case 2:
		return &[][]int{
			[]int{0, 1, 1, 1, 0},
			[]int{0, 0, 0, 1, 0},
			[]int{0, 1, 1, 1, 0},
			[]int{0, 1, 0, 0, 0},
			[]int{0, 1, 1, 1, 0},
		}
	case 3:
		return &[][]int{
			[]int{0, 1, 1, 1, 0},
			[]int{0, 0, 0, 1, 0},
			[]int{0, 1, 1, 1, 0},
			[]int{0, 0, 0, 1, 0},
			[]int{0, 1, 1, 1, 0},
		}
	case 4:
		return &[][]int{
			[]int{0, 1, 0, 1, 0},
			[]int{0, 1, 0, 1, 0},
			[]int{0, 1, 1, 1, 0},
			[]int{0, 0, 0, 1, 0},
			[]int{0, 0, 0, 1, 0},
		}
	case 5:
		return &[][]int{
			[]int{0, 1, 1, 1, 0},
			[]int{0, 1, 0, 0, 0},
			[]int{0, 1, 1, 1, 0},
			[]int{0, 0, 0, 1, 0},
			[]int{0, 1, 1, 1, 0},
		}
	case 6:
		return &[][]int{
			[]int{0, 1, 1, 1, 0},
			[]int{0, 1, 0, 0, 0},
			[]int{0, 1, 1, 1, 0},
			[]int{0, 1, 0, 1, 0},
			[]int{0, 1, 1, 1, 0},
		}
	case 7:
		return &[][]int{
			[]int{0, 1, 1, 1, 0},
			[]int{0, 0, 0, 1, 0},
			[]int{0, 0, 1, 0, 0},
			[]int{0, 1, 0, 0, 0},
			[]int{0, 1, 0, 0, 0},
		}
	case 8:
		return &[][]int{
			[]int{0, 1, 1, 1, 0},
			[]int{0, 1, 0, 1, 0},
			[]int{0, 1, 1, 1, 0},
			[]int{0, 1, 0, 1, 0},
			[]int{0, 1, 1, 1, 0},
		}
	case 9:
		return &[][]int{
			[]int{0, 1, 1, 1, 0},
			[]int{0, 1, 0, 1, 0},
			[]int{0, 1, 1, 1, 0},
			[]int{0, 0, 0, 1, 0},
			[]int{0, 1, 1, 1, 0},
		}
	}
	return nil
}

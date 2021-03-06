package gui

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.1/glfw"

	"GoDeeper/game"
	"time"
	"syscall"
)

const (
	tile_size = 20
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

	drawEffect()

	exit := drawScore()

	window.SwapBuffers()

	if exit {
		t := time.NewTimer(time.Second * 5)
		<-t.C
		syscall.Exit(0)
	}
}

func drawBG () {
	for i := 0; i < game.BOARD_HEIGHT; i++ {
		for j := 0; j < game.BOARD_WIDTH; j++ {
			sx, sy := coordsToScreen(i, j)
			DrawErr(sx, sy)
		}
	}
}

func drawScore() bool {
	return drawScoreboard()
}

func drawEffect() {
	effectChan := game.GetEffectChan()
	select {
	case effect := <-*effectChan:
		switch effect.GetMsg() {
		case game.MSG_PORTAL_USED:
			gl.Enable(gl.BLEND)
			defer gl.Disable(gl.BLEND)
			sx, sy := coordsToScreen(effect.GetCol(),effect.GetRow())
			DrawTextureTS(sx, sy, WormSpawn)
			break
		}
		break
	default:
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

package gui

import (
	"GoDeeper/game"

	"github.com/go-gl/gl/v2.1/gl"
	"time"
)

type Earth struct {
	x, y int
}

func InitDrawPieceStates() {
	go powerChange()
}

var colorMap = map[int]rgb{
	game.Earth:          rgb{139, 69, 19},
	game.Tunnel:         rgb{102, 51, 0},
	game.Pipe:           rgb{166, 163, 157},
	game.Power:          rgb{255, 247, 0},
	game.Water:          rgb{47, 172, 250},
	game.Enemy:          rgb{0, 0, 0},
	game.SuperPowerFood: rgb{255, 0, 0},
	game.Wormhole:       rgb{51, 153, 102},
}

func DrawGopher(x, y int) {
	DrawTextureTS(x, y, Tunnel)

	gl.Enable(gl.BLEND)
	defer gl.Disable(gl.BLEND)
	DrawTextureTS(x, y, Gopher)
}

func DrawGopherSuperPower(x, y int) {
	gl.Enable(gl.BLEND)
	defer gl.Disable(gl.BLEND)
	DrawTextureTS(x, y, Gopher)
}

func DrawTunnel(x, y int) {
	DrawTextureTS(x, y, Tunnel)
}

func DrawEarth(x, y int) {
	DrawTextureTS(x, y, Earth1)
}

func DrawPipe(x, y int) {
	DrawTextureTS(x, y, Wall1)
}

var powerState = false

func powerChange() {
	ticker := time.NewTicker(time.Millisecond * 500)
	for range ticker.C {
		powerState = !powerState
	}
}
func DrawPower(x, y int) {
	DrawTextureTS(x, y, Earth1)

	gl.Enable(gl.BLEND)
	defer gl.Disable(gl.BLEND)

	if powerState {
		DrawTextureTS(x, y, PowerP)
	} else {
		DrawTextureTS(x, y, PowerG)
	}
}

func DrawWater(x, y int) {
	DrawTextureTS(x, y, Water)
}

func DrawEnemy(x, y int) {
	DrawTextureTS(x, y, Tunnel)

	gl.Enable(gl.BLEND)
	defer gl.Disable(gl.BLEND)
	DrawTextureTS(x, y, Badger)
}

func DrawWormwhole(x, y int) {
	DrawTextureTS(x, y, Earth1)

	gl.Enable(gl.BLEND)
	defer gl.Disable(gl.BLEND)
	DrawTextureTS(x, y, WormHole)
}

func DrawSuperPowerFood(x, y int) {
	DrawTextureTS(x, y, Earth1)

	gl.Enable(gl.BLEND)
	defer gl.Disable(gl.BLEND)
	DrawTextureTS(x, y, Strawberry)
}

func DrawErr(x, y int) {
	color := rgb{30, 255, 0}
	if ((x/tile_size + y/tile_size) % 2) == 0 {
		color = rgb{255, 0, 191}
	}
	NewSquareTS(x, y, color).Draw()
}

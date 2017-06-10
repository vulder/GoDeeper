package gui

import (
	"GoDeeper/game"
)

type Earth struct {
	x, y int
}

var colorMap = map[int]rgb{
	game.Earth: rgb{139,69,19},
	game.Tunnel: rgb{102, 51, 0},
	game.Pipe: rgb{166,163,157},
	game.Power: rgb{255,247,0},
	game.Water: rgb{47,172,250},
	game.Enemy: rgb{0,0,0},
	game.SuperPowerFood: rgb{255,0,0},
	game.Wormwhole: rgb{51, 153, 102},
}

func DrawGopher(x, y int) {
	DrawTextureTS(x, y, Gopher)
}

func DrawTunnel(x, y int) {
	DrawTextureTS(x, y, Tunnel)
}

func DrawEarth(x, y int) {
	DrawTextureTS(x, y, Earth1)
}

func DrawPipe(x, y int) {
	NewSquareTS(x, y, colorMap[game.Pipe]).Draw()
}

func DrawPower(x, y int) {
	NewSquareTS(x, y, colorMap[game.Power]).Draw()
}

func DrawWater(x, y int) {
	DrawTextureTS(x,y,Water)
}

func DrawEnemy(x, y int) {
	DrawTextureTS(x, y, Badger)
}

func DrawWormwhole(x, y int) {
	NewSquareTS(x, y, colorMap[game.Wormwhole]).Draw()
}

func DrawSuperPowerFood(x, y int) {
	NewSquareTS(x, y, colorMap[game.SuperPowerFood]).Draw()
}

func DrawErr(x, y int) {
	color := rgb{30,255,0}
	if ((x/tile_size + y/tile_size) % 2) == 0 {
		color = rgb{255,0,191}
	}
	NewSquareTS(x, y, color).Draw()
}


package gui

import (
	"GoDeeper/game"
)

type Earth struct {
	x, y int
}

var colorMap = map[int]rgb{
	game.Earth: rgb{139,69,19},
	game.Pipe: rgb{166,163,157},
	game.Power: rgb{255,247,0},
	game.Water: rgb{47,172,250},
}

func DrawEarth(x, y int) {
	NewSquareTS(x, y, colorMap[game.Earth]).Draw()
}

func DrawPipe(x, y int) {
	NewSquareTS(x, y, colorMap[game.Pipe]).Draw()
}

func DrawPower(x, y int) {
	NewSquareTS(x, y, colorMap[game.Power]).Draw()
}

func DrawWater(x, y int) {
	NewSquareTS(x, y, colorMap[game.Water]).Draw()
}

func DrawErr(x, y int) {
	color := rgb{30,255,0}
	if ((x/tile_size + y/tile_size) % 2) == 0 {
		color = rgb{255,0,191}
	}
	NewSquareTS(x, y, color).Draw()
}


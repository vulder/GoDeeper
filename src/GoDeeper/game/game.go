package game

import (
	"math/rand"
)

const BOARD_HEIGHT int = 100
const BOARD_WIDTH int = 200
const MAX_RAND_ROWS_AT_ONCE = 10
const BARRIERS_MIN_ROWS_BETWEEN = 2
const P_BARRIER = 0.3

const(
	Earth int = iota
	Tunnel int = iota
	Gopher int = iota
	Pipe int = iota
	Power int = iota
	Water int = iota
	Enemy int = iota
)

type GameBoard struct {
	array             [][]int
	gopherX           int
	gopherY           int
	offsetLastBarrier int
}

func (board GameBoard) GetCell(row, col int) int {
	return board.array[row][col]
}

func (board GameBoard) addRow(row []int, hasBarrier bool) {
	for row := 0; row < BOARD_HEIGHT - 1; row++ {
		for col := 0; col < BOARD_WIDTH; col++ {
			board.array[row][col] = board.array[row + 1][col]
		}
	}

	for col := 0; col < BOARD_WIDTH; col++ {
		board.array[BOARD_HEIGHT - 1][col] = row[col]
	}

	if hasBarrier {
		board.offsetLastBarrier = 0
	} else {
		board.offsetLastBarrier += 1
	}
}

func (board GameBoard) moveGopher(x, y int) {
}

func newBoard() GameBoard {
	var array [][]int = make([][]int, BOARD_HEIGHT, BOARD_HEIGHT)
	for i := 0; i < BOARD_HEIGHT; i++ {
		array[i] = make([]int, BOARD_WIDTH, BOARD_WIDTH)
	}
	board := GameBoard{array, 0, 0, 0 }

	for row := 0; row < BOARD_HEIGHT; row++ {
		newRow, containsBarrier := genRandRow(board.offsetLastBarrier)
		board.addRow(newRow, containsBarrier)
	}
	return board
}

func genRandRow(offsetLastBarrier int) ([]int, bool) {
	var row []int = make([] int, BOARD_WIDTH)

	for j := 0; j < BOARD_WIDTH; j++ {
		row[j] = Earth
	}

	containsBarrier := offsetLastBarrier >= BARRIERS_MIN_ROWS_BETWEEN && rand.Float32() <= P_BARRIER

	if containsBarrier {
		barrierHoleLeft := rand.Intn(BOARD_WIDTH)
		barrierHoleRight := barrierHoleLeft + rand.Intn(BOARD_WIDTH - barrierHoleLeft)
		for j := 0; j < BOARD_WIDTH; j++ {
			var generateBarrier bool = false
			var currBarrierType int

			if rand.Float32() <= 0.3 {
				generateBarrier = !generateBarrier

				if generateBarrier {
					switch rand.Intn(3) {
					case 0:
						currBarrierType = Pipe
						break
					case 1:
						currBarrierType = Water
						break
					case 2:
						currBarrierType = Power
						break
					}
				}
			}
			if j >= barrierHoleLeft || j <= barrierHoleRight {
				row[j] = Earth
			} else if generateBarrier {
				row[j] = currBarrierType
			}
		}
	}

	return row, containsBarrier
}

var board GameBoard

func Init() {
	board = newBoard()

}

func GetBoard() *GameBoard {
	return &board
}

func Update() {

}

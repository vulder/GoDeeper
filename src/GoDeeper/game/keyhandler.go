package game

import (
	"fmt"
	"math"
)

func GoDown() {
	changeBoard(board.gopherRow+1, board.gopherCol)
}

func GoRight() {
	changeBoard(board.gopherRow, board.gopherCol+1)
}

func GoUp() {
	changeBoard(board.gopherRow-1, board.gopherCol)
}

func GoLeft() {
	changeBoard(board.gopherRow, board.gopherCol-1)
}

func changeBoard(row int, col int) {

	if checkForCellKind(row, col, SuperPowerFood) {
		effectChan <- &TriggeredEffect{MSG_ATE_FRUIT, row, col}
		board.gopherSuperPowerCycleCount += GOPHER_SUPER_POWER_DURATION
	}

	res := board.moveGopher(row, col)
	if res != nil {
		viewEventChan <- res
	}

	checkNarrowPassage()
	checkCloseToEnemy()
}

func checkForCellKind(row int, col int, _ int) bool {
	return row >= 0 && row < BOARD_HEIGHT && col >= 0 && col < BOARD_WIDTH && board.GetCell(row, col) == SuperPowerFood
}

func checkNarrowPassage() {
	row := board.gopherRow
	col := board.gopherCol

	getsBonus := false

	if col == 0 {
		getsBonus = isBorder(row, col+1) || isBorder(row, col+2)
	} else if col == BOARD_WIDTH-1 {
		getsBonus = isBorder(row, col-1) || isBorder(row, col-2)
	} else {
		getsBonus = (col > 1 && isBorder(row, col-2) && isBorder(row, col+1)) ||
				(isBorder(row, col-1) && isBorder(row, col+1)) ||
				(col < BOARD_WIDTH-3 && isBorder(row, col-1) && isBorder(row, col+2))
	}
	if getsBonus {
		fmt.Println("Narrow Passage!")
		score += NARROW_PASSAGE_REWARD
		scoreChan <- &ScoreUpdate{MSG_BEST_NAVIGATOR, PassedNarrowPassage, score, NARROW_PASSAGE_REWARD }
	}
}

func checkCloseToEnemy() {
	row := board.gopherRow
	col := board.gopherCol

	getsBonus := false

	for i := int(math.Max(0, float64(row-1))); i < int(math.Min(float64(BOARD_HEIGHT), float64(row+2))); i++ {
		for j := int(math.Max(0, float64(col-1))); j < int(math.Min(float64(BOARD_WIDTH), float64(col+2))); j++ {
			getsBonus = getsBonus || (i != j && board.GetCell(i, j) == Enemy)
		}
	}
	if getsBonus {
		fmt.Println("That was close!")
		score += CLOSE_TO_ENEMY_REWARD
		scoreChan <- &ScoreUpdate{MSG_THAT_WAS_CLOSE, BarelyEscaped, score, CLOSE_TO_ENEMY_REWARD}
	}
}

func isBorder(row, col int) bool {
	switch board.GetCell(row, col) {
	case Water, Pipe, Power:
		return true
	default:
		return false
	}
}

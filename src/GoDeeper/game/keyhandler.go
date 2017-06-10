package game

import "fmt"

func GoDown() {
	changeBoard(board.gopherRow + 1, board.gopherCol)
}

func GoRight() {
	changeBoard(board.gopherRow, board.gopherCol + 1)
}

func GoUp() {
	changeBoard(board.gopherRow - 1, board.gopherCol)
}

func GoLeft() {
	changeBoard(board.gopherRow, board.gopherCol - 1)
}

func changeBoard(row int, col int) {
	res := board.moveGopher(row, col)
	if res != nil {
		viewEventChan <- res
	}

	checkNarrowPassage()
}

func checkNarrowPassage() {
	row := board.gopherRow
	col := board.gopherCol

	getsBonus := false

	if col == 0 && (isBorder(row, col + 1) || isBorder(row, col + 2)) {
		getsBonus = true
	} else if col == BOARD_WIDTH - 1 && (isBorder(row, col - 1) || isBorder(row, col - 2)) {
		getsBonus = true
	} else {
		getsBonus = (col > 1 && isBorder(row, col - 2) && isBorder(row, col + 1)) ||
				        (isBorder(row, col - 1) && isBorder(row, col + 1)) ||
								(col < BOARD_WIDTH - 3 && isBorder(row, col - 1) && isBorder(row, col + 2))
	}
	if getsBonus {
		fmt.Println("Narrow Passage!")
		score += NARROW_PASSAGE_UPDATE
		scoreChan <- &ScoreUpdate{ MSG_BEST_NAVIGATOR, PassedNarrowPassage, score, NARROW_PASSAGE_UPDATE }
	}
}

func isBorder(row, col int) bool {
	switch board.GetCell(row, col) {
	case Water, Pipe, Power:
		return true
	default: return false
	}
}
package game

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
}

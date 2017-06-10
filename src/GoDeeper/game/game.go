package game

import (
	"math/rand"
)

const BOARD_HEIGHT int = 100
const BOARD_WIDTH int = 200
const BARRIERS_MIN_ROWS_BETWEEN int = 2
const P_BARRIER float32 = 0.7
const P_NEW_BADGER float32 = 0.05
const MAX_N_BADGERS int = 10
const BADGER_STEP_SIZE int = 10

const (
	Earth  int = iota
	Tunnel int = iota
	Gopher int = iota
	Pipe   int = iota
	Power  int = iota
	Water  int = iota
	Enemy  int = iota
)

const (
	keepLeft int = iota
	keepRight int = iota
	right int = iota
	left int = iota
	down int = iota
)

type Badger struct {
	currRow int
	currCol int
	remainingRowsDownward int
	direction int
}

var badgers []*Badger = make([]*Badger, MAX_N_BADGERS, MAX_N_BADGERS)

var currNBadgers int = 0

type GameBoard struct {
	array             [][]int
	gopherCol         int
	gopherRow         int
	offsetLastBarrier int
}

func (board *GameBoard) GetCell(row, col int) int {
	return board.array[row][col]
}

func (board *GameBoard) addRow(newRow []int, hasBarrier bool) {
	for row := 0; row < BOARD_HEIGHT-1; row++ {
		for col := 0; col < BOARD_WIDTH; col++ {
			board.array[row][col] = board.array[row+1][col]
		}
	}

	for col := 0; col < BOARD_WIDTH; col++ {
		board.array[BOARD_HEIGHT-1][col] = newRow[col]
	}

	if hasBarrier {
		board.offsetLastBarrier = 0
	} else {
		board.offsetLastBarrier += 1
	}
}

type GopherCollision struct {
	msg string
	row int
	col int
}

func (board *GameBoard) moveGopher(row, col int) *GopherCollision {
	switch board.array[row][col] {
	case Pipe:
		return &GopherCollision{"Gopher hit a pipe!", row, col }
	case Power:
		return &GopherCollision{"Gopher got grilled!", row, col }
	case Water:
		return &GopherCollision{"Gopher got drowned!", row, col }
	case Enemy:
		return &GopherCollision{"Gopher meal!", row, col }
	default:
		break
	}
	board.array[board.gopherRow][board.gopherCol] = Tunnel
	board.array[row][col] = Gopher
	board.gopherRow = row
	board.gopherCol = col
	return nil
}

func moveBadgerKeepLeftRight(b *Badger, horizontalStep int) {
	for i := 0; i < BADGER_STEP_SIZE; i++ {
		board.array[b.currRow][b.currCol] = Tunnel
		switch board.GetCell(b.currRow, b.currCol + horizontalStep) {
		case Tunnel:
		case Earth:
			b.currCol = b.currCol + horizontalStep
			break
		default:
			if b.currRow+1 < BOARD_HEIGHT-1 {
				b.currRow = b.currRow + 1
				b.remainingRowsDownward -= 1
			}
		}
		board.array[b.currRow][b.currCol] = Enemy
	}
}

func moveBadger(b *Badger) {
	for i := 0; i < BADGER_STEP_SIZE; i++ {
		board.array[b.currRow][b.currCol] = Tunnel
		switch b.direction {
		case down:
			if b.currRow < BOARD_HEIGHT - 1 {
				switch board.array[b.currRow + 1][b.currCol] {
				case Earth:
				case Tunnel:
					b.currRow += 1
					b.remainingRowsDownward -= 1
					break
				default:
					if rand.Float32() <= 0.5 {
						b.direction = left
					} else {
						b.direction = right
					}
					i -= 1
				}
			}
			break

		case left:
			if b.currRow == BOARD_HEIGHT - 1 {
				b.direction = down
				i -= 1
				break
			}
			switch board.array[b.currRow + 1][b.currCol] {
			case Tunnel:
			case Earth:
				b.direction = down
				i -= 1
				break
			default:
				if b.currCol > 0 {
					b.currCol -= 1
				}
			}
			break

		case right:
			if b.currRow == BOARD_HEIGHT - 1 {
				b.direction = down
				i -= 1
				break
			}
			switch board.array[b.currRow + 1][b.currCol] {
			case Tunnel:
			case Earth:
				b.direction = down
				i -= 1
				break
			default:
				if b.currCol < BOARD_WIDTH - 1 {
					b.currCol += 1
				}
			}
		}
		board.array[b.currRow][b.currCol] = Enemy
	}
}

func updateBadgers() {
	// delete existing badgers (if close to the edge)
	for i := 0; i < len(badgers); i++ {
		var b *Badger = badgers[i]
		if b != nil {
			var distance2Edge int
			switch b.direction {
			case left:
			case keepLeft:
				distance2Edge = b.currCol
				break
			case right:
			case keepRight:
				distance2Edge = BOARD_WIDTH - b.currCol - 1
				break
			default:
				distance2Edge = BADGER_STEP_SIZE
			}
			if distance2Edge < BADGER_STEP_SIZE {
				b = nil
				currNBadgers -= 1
			}
		}
	}

	// make existing badgers leave the board to left or right when it is time
	for i := 0; i < len(badgers); i++ {
		var b *Badger = badgers[i]
		if b != nil {
			if b.remainingRowsDownward <= 0 {
				switch b.direction {
				case down:
					if rand.Float32() <= 0.5 {
						b.direction = keepLeft
					} else {
						b.direction = keepRight
					}
					break;
				case left:
					b.direction = keepLeft
					break;
				case right:
					b.direction = keepRight
					break;
				default:
					break;
				}
			}
		}
	}

	// move existing badgers
	for i := 0; i < len(badgers); i++ {
		var b *Badger = badgers[i]
		if b != nil {
			switch b.direction {
			case keepLeft:
				moveBadgerKeepLeftRight(b, -1)
				break;
			case keepRight:
				moveBadgerKeepLeftRight(b, +1)
				break;
			default:
				moveBadger(b)
			}
		}

		// maybe generate new badger

	}
}

func newBoard() GameBoard {
	var array [][]int = make([][]int, BOARD_HEIGHT, BOARD_HEIGHT)
	for i := 0; i < BOARD_HEIGHT; i++ {
		array[i] = make([]int, BOARD_WIDTH, BOARD_WIDTH)
	}
	board := GameBoard{array, 0, 0, BARRIERS_MIN_ROWS_BETWEEN - 1 }

	for row := 0; row < BOARD_HEIGHT; row++ {
		newRow, containsBarrier := genRandRow(board.offsetLastBarrier)
		board.addRow(newRow, containsBarrier)
	}
	board.array[0][0] = Gopher
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
		barrierHoleRight := barrierHoleLeft + rand.Intn(BOARD_WIDTH-barrierHoleLeft)
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
	for i := 0; i < len(badgers); i++ {
		badgers[i] = nil
	}
}

func GetBoard() *GameBoard {
	return &board
}

func Update() {

}

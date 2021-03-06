package game

import (
	"math/rand"
	"time"
	"strconv"
	"fmt"
)

const BOARD_HEIGHT int = 50
const BOARD_WIDTH int = 25
const BARRIERS_MIN_ROWS_BETWEEN int = 2
const P_ROW_HAS_BARRIER float32 = 0.7
const P_PLACE_BARRIER float32 = 0.3
const P_WORMHOLES float32 = 0.10
const P_NEW_BADGER float32 = 0.30
const MAX_N_BADGERS int = 5
const MAX_N_PORTALS int = 6
const BADGER_MAX_VERTICAL_WAY = 100
const BADGER_STEP_SIZE int = 6

const P_SUPER_POWER_FOOD float32 = 0.07

const N_INIT_FREEZING_CYCLES = 5

const MSG_GOPHER_PIPE string = "Gopher hit a pipe!"
const MSG_GOPHER_GRILLED string = "Gopher got grilled!"
const MSG_GOPHER_DROWNED string = "Gopher drowned!"
const MSG_NOM_NOM string = "Om nom nom!"

const MSG_GO_GO string = "Go go go!"
const MSG_BEST_NAVIGATOR string = "Best navigator!"
const MSG_THAT_WAS_CLOSE string = "That was close!"
const MSG_DEAD_BADGER string = "Only a dead badger is a good badger!"

const MSG_PORTAL_USED string = "portal used"
const MSG_ATE_FRUIT string = "ate fruit"

const SCORE_INCR_INTERVAL = 300
const FREQ_SCORE_INCR = 5
const NARROW_PASSAGE_REWARD = 50
const CLOSE_TO_ENEMY_REWARD = 100
const DEAD_BADGER_REWARD = 150

const GOPHER_SUPER_POWER_DURATION = 7

const (
	Earth            int = iota
	Tunnel           int = iota
	Gopher           int = iota
	GopherSuperPower int = iota
	Pipe             int = iota
	Power            int = iota
	Water            int = iota
	Enemy            int = iota
	SuperPowerFood   int = iota
	Wormhole         int = iota
)

const (
	keepLeft  int = iota
	keepRight int = iota
	right     int = iota
	left      int = iota
	down      int = iota
)

type Badger struct {
	currRow               int
	currCol               int
	remainingRowsDownward int
	direction             int
}

type Portal struct {
	colOne int
	colTwo int
	row    int
}

const (
	RegularBonus        = iota
	PassedNarrowPassage = iota
	BarelyEscaped       = iota
	DeadBadger          = iota
)

func (b *Badger) String() string {
	res := "row: " + strconv.Itoa(b.currRow) +
			"\ncol: " + strconv.Itoa(b.currRow) +
			"\nrem rows: " + strconv.Itoa(b.remainingRowsDownward) +
			"\ndirection: "
	switch b.direction {
	case keepLeft:
		res += "keepLeft"
		break
	case keepRight:
		res += "keepRight"
		break
	case right:
		res += "right"
		break
	case left:
		res += "left"
		break
	case down:
		res += "down"
		break
	}
	return res
}

var portals []*Portal = make([]*Portal, MAX_N_PORTALS, MAX_N_PORTALS)
var badgers []*Badger = make([]*Badger, MAX_N_BADGERS, MAX_N_BADGERS)
var currNBadgers int = 0
var currNPortals int = 0

func deleteBadgersAt(row, col int) {
	for i := 0; i < len(badgers); i++ {
		var b *Badger = badgers[i]
		if b != nil && b.currCol == col && b.currRow == row {
			fmt.Println("Killed badger")
			badgers[i] = nil
			board.array[row][col] = Tunnel
			currNBadgers -= 1
			score += DEAD_BADGER_REWARD
			scoreChan <- &ScoreUpdate{MSG_DEAD_BADGER, DeadBadger, score, DEAD_BADGER_REWARD }
		}
	}
}

type Board struct {
	array                      [][]int
	gopherCol                  int
	gopherRow                  int
	offsetLastBarrier          int
	gopherSuperPowerCycleCount int
}

func (board *Board) GetCell(row, col int) int {
	return board.array[row][col]
}

func (board *Board) addRow(newRow []int, hasBarrier bool) {
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

func (board *Board) moveGopher(row, col int) *GopherCollision {
	if row >= 0 && row < BOARD_HEIGHT && col >= 0 && col < BOARD_WIDTH {
		switch board.array[row][col] {
		case Pipe:
			return &GopherCollision{MSG_GOPHER_PIPE, row, col }
		case Power:
			return &GopherCollision{MSG_GOPHER_GRILLED, row, col }
		case Water:
			return &GopherCollision{MSG_GOPHER_DROWNED, row, col }
		case Enemy:
			if board.gopherSuperPowerCycleCount <= 0 {
				return &GopherCollision{MSG_NOM_NOM, row, col }
			} else {
				deleteBadgersAt(row, col)
			}
		case Wormhole:
			effectChan <- &TriggeredEffect{MSG_PORTAL_USED, row, col}
			var p *Portal = getWormholesInGivenRow(row)
			board.array[row][col] = Tunnel
			if col == p.colOne {
				col = p.colTwo
			} else {
				col = p.colOne
			}
		}
		board.array[board.gopherRow][board.gopherCol] = Tunnel
		if board.gopherSuperPowerCycleCount > 0 {
			board.array[row][col] = GopherSuperPower
		} else {
			board.array[row][col] = Gopher
		}
		board.gopherRow = row
		board.gopherCol = col
	}
	return nil
}

func moveBadgerKeepLeftRight(b *Badger, horizontalStep int) *GopherCollision {
	for i := 0; i < BADGER_STEP_SIZE; i++ {
		board.array[b.currRow][b.currCol] = Tunnel
		if b.currCol+horizontalStep >= 0 && b.currCol+horizontalStep < BOARD_WIDTH {
			switch board.GetCell(b.currRow, b.currCol+horizontalStep) {
			case Tunnel, Earth, Gopher, GopherSuperPower, Enemy, Wormhole, SuperPowerFood:
				b.currCol = b.currCol + horizontalStep

			default:
				if b.currRow+1 < BOARD_HEIGHT-1 {
					b.currRow = b.currRow + 1
					b.remainingRowsDownward -= 1
				}
			}
		}
		if isOnWormhole(b) {
			effectChan <- &TriggeredEffect{MSG_PORTAL_USED, b.currRow, b.currCol}
			moveBatcherThroughPortal(b)
		}
		board.array[b.currRow][b.currCol] = Enemy
		if board.gopherCol == b.currCol && board.gopherRow == b.currRow {
			if board.gopherSuperPowerCycleCount <= 0 {
				return &GopherCollision{MSG_NOM_NOM, b.currRow, b.currCol}
			} else {
				deleteBadgersAt(b.currRow, b.currCol)
			}
		}
	}
	return nil
}

func isOnWormhole(b *Badger) bool {
	return board.GetCell(b.currRow, b.currCol) == Wormhole
}

func moveBatcherThroughPortal(b *Badger) {
	var p *Portal = getWormholesInGivenRow(b.currRow)
	if p.colOne == b.currCol {
		b.currCol = p.colTwo
	} else {
		b.currCol = p.colOne
	}
}

func moveBadger(b *Badger) *GopherCollision {
	for i := 0; i < BADGER_STEP_SIZE; i++ {
		board.array[b.currRow][b.currCol] = Tunnel
		switch b.direction {
		case down:
			if b.currRow < BOARD_HEIGHT-1 {
				switch board.array[b.currRow+1][b.currCol] {
				case Gopher, GopherSuperPower, Enemy, Tunnel, Earth, SuperPowerFood:
					b.currRow += 1
					b.remainingRowsDownward -= 1
				default:
					if rand.Float32() <= 0.5 {
						b.direction = left
					} else {
						b.direction = right
					}
					i -= 1
				}
			}

		case left:
			if b.currRow == BOARD_HEIGHT-1 {
				b.direction = down
				i -= 1
			}
			switch board.array[b.currRow+1][b.currCol] {
			case Tunnel, Earth, Enemy, Gopher, GopherSuperPower, SuperPowerFood:
				b.direction = down
				i -= 1
			default:
				if b.currCol > 0 {
					b.currCol -= 1
				}
			}

		case right:
			if b.currRow == BOARD_HEIGHT-1 {
				b.direction = down
				i -= 1
			}
			switch board.array[b.currRow+1][b.currCol] {
			case Tunnel, Earth, Enemy, Gopher, GopherSuperPower, SuperPowerFood:
				b.direction = down
				i -= 1
			default:
				if b.currCol < BOARD_WIDTH-1 {
					b.currCol += 1
				}
			}
		}

		if isOnWormhole(b) {
			moveBatcherThroughPortal(b)
			effectChan <- &TriggeredEffect{MSG_PORTAL_USED, b.currRow, b.currCol}
		}

		board.array[b.currRow][b.currCol] = Enemy
		if board.gopherCol == b.currCol && board.gopherRow == b.currRow {
			if board.gopherSuperPowerCycleCount <= 0 {
				return &GopherCollision{MSG_NOM_NOM, b.currRow, b.currCol}
			} else {
				deleteBadgersAt(b.currRow, b.currCol)
			}
		}
	}
	return nil
}

func updatePortals() {
	//repaint all portals in case they were overwritten by a badger or the gopher
	for i := 0; i < len(portals); i++ {
		var p *Portal = portals[i]
		if p != nil {
			board.array[p.row][p.colOne] = Wormhole
			board.array[p.row][p.colTwo] = Wormhole
		}
	}
}
func updateBadgers() *GopherCollision {
	// delete existing badgers (if close to the edge)
	for i := 0; i < len(badgers); i++ {
		var b *Badger = badgers[i]
		if b != nil {
			var distance2Edge int
			switch b.direction {
			case left, keepLeft:
				distance2Edge = b.currCol
			case right, keepRight:
				distance2Edge = BOARD_WIDTH - b.currCol - 1
			default:
				distance2Edge = BADGER_STEP_SIZE
			}

			if distance2Edge < 1 {
				badgers[i] = nil
				currNBadgers -= 1
				board.array[b.currRow][b.currCol] = Tunnel
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
				case left:
					b.direction = keepLeft
				case right:
					b.direction = keepRight
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
				res := moveBadgerKeepLeftRight(b, -1)
				if res != nil {
					return res
				}
			case keepRight:
				res := moveBadgerKeepLeftRight(b, +1)
				if res != nil {
					return res
				}
			default:
				res := moveBadger(b)
				if res != nil {
					return res
				}
			}
		}
	}
	// maybe generate new badger
	if rand.Float32() <= P_NEW_BADGER {
		var nNewBadgers int = rand.Intn(MAX_N_BADGERS - currNBadgers)
		for gNum := 0; gNum < nNewBadgers; gNum++ {
			var possibleStartPositions []int = getFreeSpotsInRow(2)
			if len(possibleStartPositions) == 0 {
				return nil
			}
			var gStartCol int = possibleStartPositions[rand.Intn(len(possibleStartPositions))]
			var newBadger Badger = Badger{2, gStartCol,
																		rand.Intn(BADGER_MAX_VERTICAL_WAY), down}
			for i := 0; i < len(badgers); i++ {
				if badgers[i] == nil {
					badgers[i] = &newBadger
					break
				}
			}
			board.array[newBadger.currRow][newBadger.currCol] = Enemy
			currNBadgers += 1
		}
	}
	return nil
}

func getFreeSpotsInRow(row int) []int {
	return getFreeSpotsInGivenRow(&board.array[row])
}

func getFreeSpotsInGivenRow(row *[]int) []int {
	nSpots := 0

	for i := 0; i < BOARD_WIDTH; i++ {
		switch (*row)[i] {
		case Tunnel, Earth:
			nSpots += 1
		}
	}
	var res []int = make([]int, nSpots, nSpots)

	cnt := 0
	for i := 0; i < BOARD_WIDTH; i++ {
		switch (*row)[i] {
		case Tunnel, Earth:
			res[cnt] = i
			cnt += 1
		}
	}

	return res
}

func getWormholesInGivenRow(row int) *Portal {
	for i := 0; i < len(portals); i++ {
		var p *Portal = portals[i]
		if p != nil && p.row == row {
			return p
		}
	}
	return nil
}

func newBoard() Board {
	var array [][]int = make([][]int, BOARD_HEIGHT, BOARD_HEIGHT)
	for i := 0; i < BOARD_HEIGHT; i++ {
		array[i] = make([]int, BOARD_WIDTH, BOARD_WIDTH)
	}
	board := Board{array, 0, 0, BARRIERS_MIN_ROWS_BETWEEN - 1,
				   0 }

	for row := 0; row < BOARD_HEIGHT; row++ {
		newRow, containsBarrier := genRandRow(board.offsetLastBarrier, row)
		board.addRow(newRow, containsBarrier)
	}
	board.array[0][0] = Gopher
	return board
}

func genRandRow(offsetLastBarrier int, rowNum int) ([]int, bool) {
	var row []int = make([] int, BOARD_WIDTH)

	for j := 0; j < BOARD_WIDTH; j++ {
		row[j] = Earth
	}

	containsBarrier := offsetLastBarrier >= BARRIERS_MIN_ROWS_BETWEEN && rand.Float32() <= P_ROW_HAS_BARRIER

	if containsBarrier {
		barrierHoleLeft := rand.Intn(BOARD_WIDTH)
		barrierHoleRight := barrierHoleLeft + rand.Intn(BOARD_WIDTH-barrierHoleLeft)

		var currBarrierLen int = 0
		var currBarrierType int = Pipe

		for j := 0; j < BOARD_WIDTH; j++ {
			if j >= barrierHoleLeft && j <= barrierHoleRight {
				currBarrierLen = 0
				row[j] = Earth
				continue
			}

			if currBarrierLen == 0 && rand.Float32() <= P_PLACE_BARRIER {
				currBarrierLen = rand.Intn(BOARD_WIDTH / 2)
				switch rand.Intn(3) {
				case 0:
					currBarrierType = Pipe
				case 1:
					currBarrierType = Water
				case 2:
					currBarrierType = Power
				}
			}

			if currBarrierLen > 0 {
				row[j] = currBarrierType
				currBarrierLen -= 1
			} else {
				row[j] = Earth
			}
		}
	}

	// place special food
	if rand.Float32() <= P_SUPER_POWER_FOOD {
		var freeSpots []int = getFreeSpotsInGivenRow(&row)
		if len(freeSpots) > 0 {
			pos := freeSpots[rand.Intn(len(freeSpots))]
			row[pos] = SuperPowerFood
		}
	}

	// place worm holes
	var freeSpaces []int = getFreeSpotsInGivenRow(&row)
	if rand.Float32() <= P_WORMHOLES && len(freeSpaces) != 1 && currNPortals < MAX_N_PORTALS {
		row = addWormholes(row, rowNum, freeSpaces)
	}

	return row, containsBarrier
}

func addWormholes(row []int, rowNum int, freeSpaces []int) []int {
	var spotOne int = 0
	var spotTwo int = 0
	//take the holes that are furthest from each other if borders are in the row, else random columns
	if len(freeSpaces) < len(row) {
		spotOne = freeSpaces[0]
		spotTwo = freeSpaces[len(freeSpaces)-1]
	} else {
		// spawn the first hole inside the first half of the row
		spotOne = rand.Intn(BOARD_WIDTH / 3)
		fmt.Println("0: ", spotOne)
		// make at least one space between the holes
		for i := 0; i < 1; i++ {
			spotTwo = rand.Intn(BOARD_WIDTH)
			if spotTwo-BOARD_WIDTH/3 < spotOne {
				i--
			}
		}
	}
	row[spotOne] = Wormhole
	row[spotTwo] = Wormhole

	for i := 0; i < len(portals); i++ {
		if portals[i] == nil {
			portals[i] = &Portal{spotOne, spotTwo, rowNum }
			break
		}
	}
	currNPortals += 1
	return row
}

var board Board

var score int

var viewEventChan chan *GopherCollision = make(chan *GopherCollision, 100)
var scoreChan chan *ScoreUpdate = make(chan *ScoreUpdate, 100)
var effectChan chan *TriggeredEffect = make(chan *TriggeredEffect, 100)

func Init() {
	board = newBoard()
	for i := 0; i < len(badgers); i++ {
		badgers[i] = nil
	}
	score = 0
}

func GetBoard() *Board {
	return &board
}

func GetEventChan() *chan *GopherCollision {
	return &viewEventChan
}

func GetScoreUpdateChan() *chan *ScoreUpdate {
	return &scoreChan
}

func GetEffectChan() *chan *TriggeredEffect {
	return &effectChan
}

func shiftBadgers() {
	for i := 0; i < len(badgers); i++ {
		var b *Badger = badgers[i]
		if b != nil {
			b.currRow -= 1
			if b.currRow < 0 {
				badgers[i] = nil
				currNBadgers -= 1
			}
		}
	}
}

func shiftPortals() {
	for i := 0; i < len(portals); i++ {
		var p *Portal = portals[i]
		if p != nil {
			p.row--
			if p.row < 0 {
				portals[i] = nil
				currNPortals--
			}
		}
	}
}

var init_freezing_counter = N_INIT_FREEZING_CYCLES
var score_incr_counter = 0

func Update(dt time.Duration) {
	//Drags the board up
	time.Sleep(dt)
	if init_freezing_counter > 0 {
		init_freezing_counter -= 1
	} else {
		newRow, containsBarrier := genRandRow(board.offsetLastBarrier, BOARD_HEIGHT)
		board.addRow(newRow, containsBarrier)
		board.gopherRow--
		shiftBadgers()
		shiftPortals()
		if board.gopherRow < 0 {
			board.gopherRow = 0
			viewEventChan <- &GopherCollision{"Gopher got dragged outside!",
																				board.gopherRow, board.gopherCol}
		}
	}
	updatePortals()
	res := updateBadgers()
	if res != nil {
		viewEventChan <- res
	}

	score_incr_counter += 1
	if score_incr_counter >= SCORE_INCR_INTERVAL {
		score_incr_counter = 0
		score += FREQ_SCORE_INCR
		scoreChan <- &ScoreUpdate{MSG_GO_GO, RegularBonus, score, FREQ_SCORE_INCR }
	}

	if board.gopherSuperPowerCycleCount > 0 {
		board.gopherSuperPowerCycleCount -= 1
	}
}

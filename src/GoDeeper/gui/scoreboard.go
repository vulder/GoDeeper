package gui

import (
	"GoDeeper/game"
	"fmt"
	"time"
	"sync/atomic"
	"strconv"
	"github.com/go-gl/gl/v2.1/gl"
)

const SCORE_WIDTH = 250
const SCORE_HEIGHT = 80
const BOTTOM_DISTANCE = 30
const RIGHT_DISTANCE = 30

const num_w_size = 9 * 2
const num_h_size = 14 * 2
const num_gap = 1

const (
	black	uint32 = iota
	red		uint32 = iota
	blue	uint32 = iota
	gray	uint32 = iota
	dead	uint32 = iota
)

var state uint32 = black
var currentScore int32 = 0

var resetTimer = time.NewTimer(time.Second * 2)

func checkChannels() {
	for {
		checkEventChan()
		checkScoreChan()
	}
}

func resetState() {
	<-resetTimer.C
	atomic.StoreUint32(&state, black)
}

func checkEventChan() {
	evenChan := game.GetEventChan()
	select {
	case event := <-*evenChan:
		switch event.Msg {
		case game.MSG_GOPHER_PIPE:
			atomic.StoreUint32(&state, gray)
			resetTimer.Reset(time.Second * 2)
			go resetState()
			fmt.Println("pipe")
			break
		case game.MSG_GOPHER_DROWNED:
			atomic.StoreUint32(&state, blue)
			resetTimer.Reset(time.Second * 2)
			go resetState()
			fmt.Println("drowned")
			break
		case game.MSG_GOPHER_GRILLED:
			atomic.StoreUint32(&state, red)
			resetTimer.Reset(time.Second * 2)
			go resetState()
			fmt.Println("grilled")
			break
		case game.MSG_NOM_NOM:
			fmt.Println("om nom nom")
			atomic.StoreUint32(&state, dead)
			break
		}
		break
	default:
	}
}

func checkScoreChan() {
	scoreChan := game.GetScoreUpdateChan()
	select {
	case score := <-*scoreChan:
		atomic.StoreInt32(&currentScore, int32(score.NewScore))
		break
	default:
	}
}

func drawScoreboard() bool {
	var color rgb = rgb{0,0,0}

	switch atomic.LoadUint32(&state) {
	case red:
		color = rgb{255,0,0}
		break
	case blue:
		color = rgb{0,0,255}
		break
	case gray:
		color = rgb{160,160,160}
		break
	case dead:
		DrawTexture(0,0,GetWidth(),GetHigh(), Dead)
		return true
	}

	htop := GetHigh()-SCORE_HEIGHT-BOTTOM_DISTANCE
	wtop := GetWidth()-SCORE_WIDTH-RIGHT_DISTANCE

	NewSquare(htop,
		wtop,
		SCORE_WIDTH,SCORE_HEIGHT,
		color).Draw()

	// Drawing actual score
	cScore := atomic.LoadInt32(&currentScore)

	gl.Enable(gl.BLEND)
	defer gl.Disable(gl.BLEND)

	number_str := strconv.Itoa(int(cScore))
	for i, char := range number_str {
		var num int8
		switch char {
		case '0':
			num = NUM_0
			break
		case '1':
			num = NUM_1
			break
		case '2':
			num = NUM_2
			break
		case '3':
			num = NUM_3
			break
		case '4':
			num = NUM_4
			break
		case '5':
			num = NUM_5
			break
		case '6':
			num = NUM_6
			break
		case '7':
			num = NUM_7
			break
		case '8':
			num = NUM_8
			break
		case '9':
			num = NUM_9
			break
		}

		wtop_ralign := wtop + SCORE_WIDTH - (num_w_size * ((len(number_str) - i)))
		DrawTexture(htop, wtop_ralign, num_w_size, num_h_size, num)
	}

	return false
}

package game

type GopherCollision struct {
	msg string
	row int
	col int
}

func (gc GopherCollision) GetMsg() string {
	return gc.msg
}

func (gc GopherCollision) GetRow() int {
	return gc.row
}

func (gc GopherCollision) GetCol() int {
	return gc.col
}

type ScoreUpdate struct {
	msg       string
	kind      int // is one of RegularBonus, passedNarrowPassage, barelyEscaped
	newScore  int
	increment int
}

func (su ScoreUpdate) GetMsg() string {
	return su.msg
}

func (su ScoreUpdate) GetKind() int {
	return su.kind
}

func (su ScoreUpdate) GetNewScore() int {
	return su.newScore
}

func (su ScoreUpdate) GetIncrement() int {
	return su.increment
}

type TriggeredEffect struct {
	msg string
	row int
	col int
}

func (te TriggeredEffect) GetMsg() string {
	return te.msg
}

func (te TriggeredEffect) GetRow() int {
	return te.row
}

func (te TriggeredEffect) GetCol() int {
	return te.col
}

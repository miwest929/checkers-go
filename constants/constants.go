package constants

type Player int

const (
	WHITE_PLAYER Player = 0
	BLACK_PLAYER        = 1
)

type Move int

const (
	UPLEFT    Move = 0
	UPRIGHT        = 1
	DOWNLEFT       = 2
	DOWNRIGHT      = 3
)

type PossibleMove struct {
	RowIdx int
	ColIdx int
	M      Move
}

func NewPossibleMove(rIdx int, cIdx int, m Move) PossibleMove {
	return PossibleMove{RowIdx: rIdx, ColIdx: cIdx, M: m}
}

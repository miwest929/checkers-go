package main

import (
	"fmt"
	"strings"
)

var (
	EMPTY          = " "
	WHITE          = "w"
	BLACK          = "b"
	WHITE_PROMOTED = "W"
	BLACK_PROMOTED = "B"

	WHITE_PLAYER = 0
	BLACK_PLAYER = 1

	UPLEFT    = 0
	UPRIGHT   = 1
	DOWNLEFT  = 2
	DOWNRIGHT = 3
)

type Board struct {
	board [8][8]string
}

func NewBoard() *Board {
	return &Board{board: initialBoard()}
}

func (b *Board) String() string {
	str := ""
	for _, row := range b.board {
		rowStr := strings.Join(row[:], "")
		rowStr = strings.Replace(rowStr, " ", ".", -1)
		str += rowStr + "\n"
	}

	return str
}

func emptyBoard() [8][8]string {
	var board [8][8]string
	for rowIdx, row := range board {
		for colIdx, _ := range row {
			board[rowIdx][colIdx] = EMPTY
		}
	}

	return board
}

func initialBoard() [8][8]string {
	board := emptyBoard()

	for _, idx := range []int{0, 1, 2} {
		for _, cIdx := range pieceIndices(idx) {
			board[idx][cIdx] = WHITE
		}
	}

	for _, idx := range []int{5, 6, 7} {
		for _, cIdx := range pieceIndices(idx) {
			board[idx][cIdx] = BLACK
		}
	}

	return board
}

func pieceIndices(rowIndex int) []int {
	if rowIndex%2 == 0 {
		return []int{1, 3, 5, 7}
	}

	return []int{0, 2, 4, 6}
}

type CheckersGame struct {
}

func (game *CheckersGame) Start() {
}

func main() {
	b := NewBoard()
	fmt.Println(b)
}

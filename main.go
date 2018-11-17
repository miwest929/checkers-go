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
	grid [8][8]string
}

func getMoveVectors() [4][2]int {
	upperLeft := [2]int{-1, -1}
	upperRight := [2]int{-1, 1}
	downLeft := [2]int{1, -1}
	downRight := [2]int{1, 1}

	return [4][2]int{upperLeft, upperRight, downLeft, downRight}
}

// Naming conventions: board -> refers to type Board. grid is the data structure that houses the pieces.
func NewBoard() *Board {
	return &Board{grid: initialGrid()}
}

func NewBoardWithState(state [8][8]string) *Board {
	return &Board{grid: state}
}

func (b *Board) MakeMove(pieceRowIdx int, pieceColIdx int, move int) *Board {
	if !b.isMoveLegal(pieceRowIdx, pieceColIdx, move, true) {
		fmt.Println("MOVE IS ILLEGAL")
		return b
	}

	moveVectors := getMoveVectors()

	newRowIdx := pieceRowIdx + moveVectors[move][0]
	newColIdx := pieceColIdx + moveVectors[move][1]
	newState := b.Clone()
	if b.grid[newRowIdx][newColIdx] == EMPTY {
		newState[newRowIdx][newColIdx] = b.grid[pieceRowIdx][pieceColIdx]
		newState[pieceRowIdx][pieceColIdx] = EMPTY
		newBoard := NewBoardWithState(newState)
		//newBoard.makeKnightPiece(newRowIdx, newColIdx)
		return newBoard
	}

	nextRowIdx := newRowIdx + moveVectors[move][0]
	nextColIdx := newColIdx + moveVectors[move][1]
	if b.grid[nextRowIdx][nextColIdx] == EMPTY {
		newState[nextRowIdx][nextColIdx] = b.grid[pieceRowIdx][pieceColIdx]
		newState[newRowIdx][newColIdx] = EMPTY
		newState[pieceRowIdx][pieceColIdx] = EMPTY

		newBoard := NewBoardWithState(newState)
		//newBoard.makeKnightPiece(nextRowIdx, nextColIdx)
		return newBoard
	}

	return b
}

func (b *Board) Clone() [8][8]string {
	grid := emptyGrid()
	for rIdx, row := range b.grid {
		for cIdx, _ := range row {
			grid[rIdx][cIdx] = b.grid[rIdx][cIdx]
		}
	}

	return grid
}

func (b *Board) String() string {
	str := " _ _ _ _ _ _ _ _\n"
	for _, row := range b.grid {
		rowStr := strings.Join(row[:], "|")
		rowStr = "|" + rowStr + "|"
		str += rowStr + "\n"
	}

	return str
}

func (b *Board) makeKnightPiece(rowIdx int, colIdx int, move int) {
	if b.grid[rowIdx][colIdx] == "w" && rowIdx == 7 {
		b.grid[rowIdx][colIdx] = "W"
		return
	}

	if b.grid[rowIdx][colIdx] == "b" && rowIdx == 0 {
		b.grid[rowIdx][colIdx] = "B"
		return
	}
}

func (b *Board) nextPossibleMoves(playersTurn int) [][]int {
	piece := "w"
	if playersTurn == BLACK_PLAYER {
		piece = "b"
	}

	moves := make([][]int, 0)
	for rIdx := 0; rIdx <= 7; rIdx++ {
		for cIdx := 0; cIdx <= 7; cIdx++ {
			if strings.ToLower(b.grid[rIdx][cIdx]) != piece {
				continue
			}

			if b.isMoveLegal(rIdx, cIdx, UPLEFT, true) {
				moves = append(moves, []int{rIdx, cIdx, UPLEFT})
			}
			if b.isMoveLegal(rIdx, cIdx, UPRIGHT, true) {
				moves = append(moves, []int{rIdx, cIdx, UPRIGHT})
			}
			if b.isMoveLegal(rIdx, cIdx, DOWNLEFT, true) {
				moves = append(moves, []int{rIdx, cIdx, DOWNLEFT})
			}
			if b.isMoveLegal(rIdx, cIdx, DOWNRIGHT, true) {
				moves = append(moves, []int{rIdx, cIdx, DOWNRIGHT})
			}
		}
	}

	return moves
}

/*func (b *Board) getPieces() map[string][][]int {
  pieces := make(map[string][][]int)
  pieces["white"] = make([][]int, 0)
  pieces["black"] = make([][]int, 0)

  for rIdx := 0; rIdx <= 7; rIdx++ {
    for cIdx := 0; cIdx <= 7; cIdx++ {
      if b.grid[rIdx][cIdx] == 'w' {
        pieces[]
      }
      if b.grid[rIdx][cIdx] == 'W' {
      }
      if b.grid[rIdx][cIdx] == 'b' {
      }
      if b.grid[rIdx][cIdx] == 'B' {
      }
    }
  }

  return pieces
}*/

func (b *Board) isMoveLegal(pieceRowIdx int, pieceColIdx int, move int, recursiveCheck bool) bool {
	piece := b.grid[pieceRowIdx][pieceColIdx]
	if piece == EMPTY {
		return false
	}

	if piece == WHITE && (move == UPLEFT || move == UPRIGHT) {
		return false
	}

	if piece == WHITE && (move == DOWNLEFT || move == DOWNRIGHT) {
		return false
	}

	// -1, -1 == UPPER_LEFT
	if move == UPLEFT && (pieceRowIdx <= 0 || pieceColIdx <= 0) {
		return false
	}

	if move == UPRIGHT && (pieceRowIdx <= 0 || pieceColIdx >= 7) {
		return false
	}

	if move == DOWNLEFT && (pieceRowIdx >= 7 || pieceColIdx <= 0) {
		return false
	}

	if move == DOWNRIGHT && (pieceRowIdx >= 7 || pieceColIdx >= 7) {
		return false
	}

	moveVectors := getMoveVectors()
	dirVector := moveVectors[move]
	newRowIdx := pieceRowIdx + dirVector[0]
	newColIdx := pieceColIdx + dirVector[1]
	if b.grid[newRowIdx][newColIdx] == EMPTY {
		return true
	}

	// can't jump yourself
	if b.grid[pieceRowIdx][pieceColIdx] == b.grid[newRowIdx][newColIdx] {
		return false
	}

	// can piece perform a jump
	nextRowIdx := newRowIdx + dirVector[0]
	nextColIdx := newColIdx + dirVector[1]
	if nextRowIdx < 0 || nextRowIdx > 7 || nextColIdx < 0 || nextColIdx > 7 {
		return false
	}

	if b.grid[nextRowIdx][nextColIdx] == EMPTY {
		return true
	}

	return false
}

func emptyGrid() [8][8]string {
	var grid [8][8]string
	for rowIdx, row := range grid {
		for colIdx, _ := range row {
			grid[rowIdx][colIdx] = EMPTY
		}
	}

	return grid
}

func initialGrid() [8][8]string {
	grid := emptyGrid()

	for _, idx := range []int{0, 1, 2} {
		for _, cIdx := range pieceIndices(idx) {
			grid[idx][cIdx] = BLACK
		}
	}

	for _, idx := range []int{5, 6, 7} {
		for _, cIdx := range pieceIndices(idx) {
			grid[idx][cIdx] = WHITE
		}
	}

	return grid
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
	b2 := b.MakeMove(2, 1, DOWNLEFT)
	fmt.Println(b2)
}

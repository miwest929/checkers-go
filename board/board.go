package board

import (
	"checkers-go/constants"
	"strconv"
	"strings"

	"github.com/gookit/color"
)

var (
	EMPTY          = " "
	WHITE          = "w"
	BLACK          = "b"
	WHITE_PROMOTED = "W"
	BLACK_PROMOTED = "B"
)

const NO_LAST_INDEX = -1

type Board struct {
	grid         [8][8]string
	LastWhiteRow int
	LastWhiteCol int
	LastBlackRow int
	LastBlackCol int
}

func getMoveVectors() [4][2]int {
	upperLeft := [2]int{-1, -1}
	upperRight := [2]int{-1, 1}
	downLeft := [2]int{1, -1}
	downRight := [2]int{1, 1}

	return [4][2]int{upperLeft, upperRight, downLeft, downRight}
}

// NextPiecePos computes the next row,col of piece after applying given move
func NextPiecePos(rowIdx int, colIdx int, move constants.Move) [2]int {
	moveVectors := getMoveVectors()
	return [2]int{rowIdx + moveVectors[move][0], colIdx + moveVectors[move][1]}
}

// Naming conventions: board -> refers to type Board. grid is the data structure that houses the pieces.
func NewBoard() *Board {
	return &Board{grid: initialGrid()}
}

func NewBoardWithState(state [8][8]string) *Board {
	return &Board{
		grid:         state,
		LastWhiteRow: NO_LAST_INDEX,
		LastWhiteCol: NO_LAST_INDEX,
		LastBlackRow: NO_LAST_INDEX,
		LastBlackCol: NO_LAST_INDEX,
	}
}

func (b *Board) MakeMove(pieceRowIdx int, pieceColIdx int, move constants.Move) *Board {
	if !b.IsMoveLegal(pieceRowIdx, pieceColIdx, move) {
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
		return newBoard
	}

	nextRowIdx := newRowIdx + moveVectors[move][0]
	nextColIdx := newColIdx + moveVectors[move][1]
	pieceColor := b.grid[pieceRowIdx][pieceColIdx]
	if b.grid[nextRowIdx][nextColIdx] == EMPTY {
		newState[nextRowIdx][nextColIdx] = pieceColor
		newState[newRowIdx][newColIdx] = EMPTY
		newState[pieceRowIdx][pieceColIdx] = EMPTY

		newBoard := NewBoardWithState(newState)

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

func (b *Board) Display() {
	color.Red.Printf("   0 1 2 3 4 5 6 7\n")
	color.Style{color.FgWhite, color.OpBold}.Printf("   _ _ _ _ _ _ _ _\n")
	for rIdx, row := range b.grid {
		color.Red.Printf(strconv.Itoa(rIdx))
		color.Style{color.FgWhite, color.OpBold}.Printf(" |")
		for cIdx, cell := range row {
			if cell == "w" {
				if b.LastWhiteCol == rIdx && b.LastWhiteRow == cIdx {
					color.Style{color.FgMagenta, color.OpBold}.Printf(cell)
				} else {
					color.Style{color.FgMagenta}.Printf(cell)
				}
			} else if cell == "b" {
				if b.LastBlackCol == cIdx && b.LastBlackRow == rIdx {
					color.Style{color.FgCyan, color.OpBold}.Printf(cell)
				} else {
					color.Style{color.FgCyan}.Printf(cell)
				}
			} else {
				color.Style{color.FgCyan, color.OpBold}.Printf(" ")
			}
			color.Style{color.FgWhite, color.OpBold}.Printf("|")
		}
		color.White.Printf("\n")
	}
}

func (b *Board) String() string {
	str := "   0 1 2 3 4 5 6 7\n"
	str += "   _ _ _ _ _ _ _ _\n"
	for rIdx, row := range b.grid {
		rowStr := strings.Join(row[:], "|")
		rowStr = strconv.Itoa(rIdx) + " |" + rowStr + "|"
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

func (b *Board) NextPossibleMoves(playersTurn constants.Player) []constants.PossibleMove {
	piece := "w"
	if playersTurn == constants.BLACK_PLAYER {
		piece = "b"
	}

	moves := make([]constants.PossibleMove, 0)
	for rIdx := 0; rIdx <= 7; rIdx++ {
		for cIdx := 0; cIdx <= 7; cIdx++ {
			if strings.ToLower(b.grid[rIdx][cIdx]) != piece {
				continue
			}

			if b.IsMoveLegal(rIdx, cIdx, constants.UPLEFT) {
				moves = append(moves, constants.NewPossibleMove(rIdx, cIdx, constants.UPLEFT))
			}
			if b.IsMoveLegal(rIdx, cIdx, constants.UPRIGHT) {
				moves = append(moves, constants.NewPossibleMove(rIdx, cIdx, constants.UPRIGHT))
			}
			if b.IsMoveLegal(rIdx, cIdx, constants.DOWNLEFT) {
				moves = append(moves, constants.NewPossibleMove(rIdx, cIdx, constants.DOWNLEFT))
			}
			if b.IsMoveLegal(rIdx, cIdx, constants.DOWNRIGHT) {
				moves = append(moves, constants.NewPossibleMove(rIdx, cIdx, constants.DOWNRIGHT))
			}
		}
	}

	return moves
}

func (b *Board) getPieces() map[string]map[string][][]int {
	pieces := make(map[string]map[string][][]int)
	pieces["white"] = make(map[string][][]int)
	pieces["white"]["normal"] = make([][]int, 0)
	pieces["white"]["king"] = make([][]int, 0)
	pieces["black"] = make(map[string][][]int)
	pieces["black"]["normal"] = make([][]int, 0)
	pieces["black"]["king"] = make([][]int, 0)

	for rIdx := 0; rIdx <= 7; rIdx++ {
		for cIdx := 0; cIdx <= 7; cIdx++ {
			if b.grid[rIdx][cIdx] == "w" {
				pieces["white"]["normal"] = append(pieces["white"]["normal"], []int{rIdx, cIdx})
			}
			if b.grid[rIdx][cIdx] == "W" {
				pieces["white"]["king"] = append(pieces["white"]["king"], []int{rIdx, cIdx})
			}
			if b.grid[rIdx][cIdx] == "b" {
				pieces["black"]["normal"] = append(pieces["black"]["normal"], []int{rIdx, cIdx})
			}
			if b.grid[rIdx][cIdx] == "B" {
				pieces["black"]["king"] = append(pieces["black"]["king"], []int{rIdx, cIdx})
			}
		}
	}

	return pieces
}

func (b *Board) CalculateScore() int {
	piecesScore := b.getTotalPiecesScore()

	totalScore := piecesScore * 100
	return totalScore + b.whosePiecesAreCloserToKings()
}

func (b *Board) whosePiecesAreCloserToKings() int {
	pieces := b.getPieces()

	whiteScoreSum := 0
	for _, piece := range pieces["white"]["normal"] {
		whiteScoreSum += b.distanceFromBeingKnighted(piece)
	}

	blackScoreSum := 0
	for _, piece := range pieces["black"]["normal"] {
		blackScoreSum += b.distanceFromBeingKnighted(piece)
	}

	max_score := 84 // all 12 pieces are 7 moves away from being knighted
	return (max_score - blackScoreSum) - (max_score - whiteScoreSum)
}

func (b *Board) distanceFromBeingKnighted(piece []int) int {
	rowIdx := piece[0]
	colIdx := piece[1]

	if b.grid[rowIdx][colIdx] == "w" || b.grid[rowIdx][colIdx] == "W" {
		return 7 - rowIdx
	}

	if b.grid[rowIdx][colIdx] == "b" || b.grid[rowIdx][colIdx] == "B" {
		return rowIdx
	}

	return 8 // the max distance
}

func (b *Board) getTotalPiecesScore() int {
	totalScore := 0
	for _, row := range b.grid {
		for _, cell := range row {
			if cell == "b" {
				totalScore -= 1
			} else if cell == "B" {
				totalScore -= 5
			} else if cell == "w" {
				totalScore += 1
			} else if cell == "W" {
				totalScore += 5
			}
		}
	}

	return totalScore
}

func (b *Board) IsMoveLegal(pieceRowIdx int, pieceColIdx int, move constants.Move) bool {
	piece := b.grid[pieceRowIdx][pieceColIdx]
	if piece == EMPTY {
		return false
	}

	if piece == BLACK && (move == constants.UPLEFT || move == constants.UPRIGHT) {
		return false
	}

	if piece == WHITE && (move == constants.DOWNLEFT || move == constants.DOWNRIGHT) {
		return false
	}

	// -1, -1 == UPPER_LEFT
	if move == constants.UPLEFT && (pieceRowIdx <= 0 || pieceColIdx <= 0) {
		return false
	}

	if move == constants.UPRIGHT && (pieceRowIdx <= 0 || pieceColIdx >= 7) {
		return false
	}

	if move == constants.DOWNLEFT && (pieceRowIdx >= 7 || pieceColIdx <= 0) {
		return false
	}

	if move == constants.DOWNRIGHT && (pieceRowIdx >= 7 || pieceColIdx >= 7) {
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

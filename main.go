package main

import (
	"bufio"
	"checkers-go/board"
	"checkers-go/constants"
	"checkers-go/player"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gookit/color"
)

type CheckersGame struct {
	currBoard *board.Board
}

func NewGame() *CheckersGame {
	return &CheckersGame{currBoard: board.NewBoard()}
}

func (game *CheckersGame) Start() {
	ai := player.ComputerPlayer{}

	for {
		game.currBoard.Display()

		var rIdx int
		var cIdx int
		var move constants.Move
		for {
			rIdx, cIdx, move = game.GetHumanInput()

			if game.currBoard.IsMoveLegal(rIdx, cIdx, move) {
				break
			}

			fmt.Println("Your chosen move is illegal. Please choose another one.")
		}
		// human's move
		game.currBoard = game.currBoard.MakeMove(rIdx, cIdx, move)
		nextPos := board.NextPiecePos(rIdx, cIdx, move)
		// TODO: LastWhiteRow/Col is not causing last moved piece to be bolded
		game.currBoard.LastWhiteRow = nextPos[0]
		game.currBoard.LastWhiteCol = nextPos[1]

		// computer's move
		nextMove := ai.NextMove(game.currBoard)
		game.currBoard = game.currBoard.MakeMove(nextMove.RowIdx, nextMove.ColIdx, nextMove.M)
		nextPos = board.NextPiecePos(nextMove.RowIdx, nextMove.ColIdx, nextMove.M)
		// TODO: LastBlackRow/Col is not causing last moved piece to be bolded
		game.currBoard.LastBlackRow = nextMove.RowIdx
		game.currBoard.LastBlackCol = nextMove.ColIdx
	}
}

func (game *CheckersGame) promptHumanChoice(prompt string) string {
	//color.Cyan.Printf("%s ->", prompt)
	color.Style{color.FgWhite, color.OpBold}.Printf("%s ", prompt)
	color.Style{color.FgCyan, color.OpBold}.Printf("-> ")
	//fmt.Print(prompt + " -> ")

	reader := bufio.NewReader(os.Stdin)
	inpt, err := reader.ReadString('\n')
	inpt = strings.TrimSuffix(inpt, "\n")
	if err != nil {
		fmt.Println(err)
	}

	return inpt
}

func (game *CheckersGame) GetHumanInput() (int, int, constants.Move) {
	inpt := game.promptHumanChoice("Which piece you like to move? Format: rowIndex,colIndex")
	indices := strings.Split(inpt, ",")
	if len(indices) <= 1 {
		fmt.Println("Must use a ',' to separate the row and column index")
	}

	rIdx, cnvErr := strconv.Atoi(indices[0])
	if cnvErr != nil {
		fmt.Println(cnvErr)
	}

	cIdx, cnvErr := strconv.Atoi(indices[1])
	if cnvErr != nil {
		fmt.Println(cnvErr)
	}

	inpt = game.promptHumanChoice("Move (UL=0, UR=1, DL=2, DR=3)")
	mv, cnvErr := strconv.Atoi(inpt)
	if cnvErr != nil {
		fmt.Println(cnvErr)
	}
	move := constants.Move(mv)

	return rIdx, cIdx, move
}

func main() {
	game := NewGame()
	game.Start()
}

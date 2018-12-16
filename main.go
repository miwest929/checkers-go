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
		fmt.Println(game.currBoard)

		isHumanMoveLegal := false
		var rIdx int
		var cIdx int
		var move constants.Move
		for !isHumanMoveLegal {
			rIdx, cIdx, move = game.GetHumanInput()
			isHumanMoveLegal = game.currBoard.IsMoveLegal(rIdx, cIdx, move)
			if !isHumanMoveLegal {
				fmt.Println("Your chosen move is illegal. Please choose another one.")
			}
		}
		// human's move
		game.currBoard = game.currBoard.MakeMove(rIdx, cIdx, move)

		// computer's move
		nextMove := ai.NextMove(game.currBoard)
		game.currBoard = game.currBoard.MakeMove(nextMove.RowIdx, nextMove.ColIdx, nextMove.M)
	}
}

func (game *CheckersGame) promptHumanChoice(prompt string) string {
	fmt.Print(prompt + " -> ")

	reader := bufio.NewReader(os.Stdin)
	inpt, err := reader.ReadString('\n')
	inpt = strings.TrimSuffix(inpt, "\n")
	if err != nil {
		fmt.Println(err)
	}

	return inpt
}

func (game *CheckersGame) GetHumanInput() (int, int, constants.Move) {
	inpt := game.promptHumanChoice("Piece Row Index (0-7)")
	rIdx, cnvErr := strconv.Atoi(inpt)
	if cnvErr != nil {
		fmt.Println(cnvErr)
	}

	inpt = game.promptHumanChoice("Piece Col Index (0-7)")
	cIdx, cnvErr := strconv.Atoi(inpt)
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

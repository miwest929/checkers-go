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
		rIdx, cIdx, move := game.GetHumanInput()

		// human's move
		game.currBoard = game.currBoard.MakeMove(rIdx, cIdx, move)

		// computer's move
		nextMove := ai.NextMove(game.currBoard)
		game.currBoard = game.currBoard.MakeMove(nextMove.RowIdx, nextMove.ColIdx, nextMove.M)
	}
}

func (game *CheckersGame) GetHumanInput() (int, int, constants.Move) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Piece Row Index (0-7) -> ")
	inpt, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	inpt = strings.TrimSuffix(inpt, "\n")
	rIdx, cnvErr := strconv.Atoi(inpt)
	if cnvErr != nil {
		fmt.Println(cnvErr)
	}

	fmt.Print("Piece Col Index (0-7) -> ")
	inpt, err = reader.ReadString('\n')
	inpt = strings.TrimSuffix(inpt, "\n")
	if err != nil {
		fmt.Println(err)
	}
	cIdx, cnvErr := strconv.Atoi(inpt)
	if cnvErr != nil {
		fmt.Println(cnvErr)
	}

	fmt.Print("Move (UL=0, UR=1, DL=2, DR=3) -> ")
	inpt, err = reader.ReadString('\n')
	inpt = strings.TrimSuffix(inpt, "\n")
	if err != nil {
		fmt.Println(err)
	}
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

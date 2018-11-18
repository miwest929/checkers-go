package main

import (
	"checkers-go/board"
	"checkers-go/constants"
	"checkers-go/player"
	"fmt"
)

type CheckersGame struct {
}

func (game *CheckersGame) Start() {
}

func main() {
	b := board.NewBoard()
	b2 := b.MakeMove(2, 1, constants.DOWNLEFT)
	fmt.Println(b2)

	n := player.NewNode(b2, 0, 0, constants.UPLEFT)
	fmt.Println(n)
}

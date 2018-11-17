package main

import (
	"checkers-go/board"
	"fmt"
)

type CheckersGame struct {
}

func (game *CheckersGame) Start() {
}

func main() {
	b := board.NewBoard()
	b2 := b.MakeMove(2, 1, board.DOWNLEFT)
	fmt.Println(b2)
}

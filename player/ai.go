package player

import (
	brd "checkers-go/board"
	"checkers-go/constants"
	//	"fmt"
)

type Node struct {
	board    *brd.Board
	player   constants.Player
	level    int
	children []*Node
	move     constants.Move
}

func NewNode(board *brd.Board, player constants.Player, level int, move constants.Move) *Node {
	node := &Node{board: board, player: player, level: level, move: move}
	node.children = make([]*Node, 0)

	return node
}

func (n *Node) String() string {
	return n.board.String()
}

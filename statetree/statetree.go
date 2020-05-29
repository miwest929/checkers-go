package statetree

import (
	brd "github.com/miwest929/checkers-go/board"
	"github.com/miwest929/checkers-go/constants"
)

type Node struct {
	B        *brd.Board
	P        constants.Player
	Level    int
	Children []*Node
	M        *constants.PossibleMove
	Score    int
}

func NewNode(board *brd.Board, player constants.Player, level int, move *constants.PossibleMove) *Node {
	node := &Node{B: board, P: player, Level: level, M: move}
	node.Children = make([]*Node, 0)

	return node
}

func (n *Node) String() string {
	return n.B.String()
}

type Tree struct {
	Root *Node
}

func NewTree(board *brd.Board) *Tree {
	node := NewNode(board, constants.BLACK_PLAYER, 0, nil)
	return &Tree{Root: node}
}

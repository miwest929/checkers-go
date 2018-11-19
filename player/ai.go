package player

import (
	brd "checkers-go/board"
	"checkers-go/constants"
	//	"fmt"

	"container/list"
)

type Node struct {
	board    *brd.Board
	player   constants.Player
	level    int
	children []*Node
	move     *constants.Move
}

func NewNode(board *brd.Board, player constants.Player, level int, move *constants.Move) *Node {
	node := &Node{board: board, player: player, level: level, move: move}
	node.children = make([]*Node, 0)

	return node
}

func (n *Node) String() string {
	return n.board.String()
}

type Tree struct {
	root *Node
}

func NewTree(board *brd.Board) *Tree {
	node := NewNode(board, constants.BLACK_PLAYER, 0, nil)
	return &Tree{root: node}
}

type ComputerPlayer struct {
}

func (cp *ComputerPlayer) constructStateTree(levelsCount int, board *brd.Board) *Tree {
	tree := NewTree(board)

	queue := list.New()
	queue.PushBack(tree.root)
	numberOfStates := 0
	for queue.Len() > 0 {
		state := queue.Front()
		queue.Remove(state)

		if state.level+1 > levelsCount {
			continue
		}

		nextMoves = state.board.nextPossibleMoves(state.player)
		for _, move := range nextMoves {
			nextBoard := state.board.MakeMove(move)
			nextState := NewNode(nextBoard, constants.opponent(state.player), move, state.level+1)
			numberOfStates += 1
			state.children.append(nextState)
			queue.PushBack(nextState)
		}
	}

	fmt.Printf("Generated %d states", numberOfStates)

	return tree
}

func (cp *ComputerPlayer) minimax(node *Node, isMax bool) int {
  if len(node.children) == 0 {
    node.score = scorer(node.Board)
    return node.score
  }

  scores := make([]int, 0)
  if isMax {
    for _, child := range node.children {
      scores = append(scores, cp.minimax(child, false))
    }

    node.score = GetMax(scores)
  } else {
    for _, child := range node.children {
      scores = append(scores, cp.minimax(child, true))
    }

    node.score = GetMin(scores)
  }
}

func GetMax(scores []int) int {
  maxValue := scores[0]
  _, value := range scores[1:] {
    if value > maxValue {
      maxValue = value
    }
  }

  return maxValue
}

func GetMin(scores []int) int {
  minValue := scores[0]
  _, value := range scores[1:] {
    if value < minValue {
      minValue = value
    }
  }

  return minValue
}

func (cp *ComputerPlayer) NextMove(board *brd.Board) *constants.Move {
  stateTree := cp.constructStateTree(3, board)
  bestScore := cp.minimax(stateTree.root, true)

  for _, child := range stateTree.root.children {
    if child.score == bestScore {
      return child.move
    }
  }

  return nil
}

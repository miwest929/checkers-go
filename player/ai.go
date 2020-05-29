package player

import (
	brd "github.com/miwest929/checkers-go/board"
	"github.com/miwest929/checkers-go/constants"
	"github.com/miwest929/checkers-go/queue"
	stree "github.com/miwest929/checkers-go/statetree"
	"fmt"
	"math"
	"time"
)

type ComputerPlayer struct {
}

func (cp *ComputerPlayer) constructStateTree(levelsCount int, board *brd.Board) *stree.Tree {
	tree := stree.NewTree(board)

	q := queue.NewQueue()
	q.Enqueue(tree.Root)
	numberOfStates := 0
	for !q.IsEmpty() {
		state := q.Dequeue()

		if state.Level+1 > levelsCount {
			continue
		}

		nextMoves := state.B.NextPossibleMoves(state.P)
		for _, move := range nextMoves {
			nextBoard := state.B.MakeMove(move.RowIdx, move.ColIdx, move.M)
			nextState := stree.NewNode(nextBoard, constants.Opponent(state.P), state.Level+1, &move)
			numberOfStates++
			state.Children = append(state.Children, nextState)
			q.Enqueue(nextState)
		}
	}

	fmt.Printf("Generated %d states\n", numberOfStates)

	return tree
}

func (cp *ComputerPlayer) minimaxMax(node *stree.Node) float64 {
	return cp.minimax(node, true, math.Inf(-1), math.Inf(1))
}

func (cp *ComputerPlayer) minimaxMin(node *stree.Node) float64 {
	return cp.minimax(node, false, math.Inf(1), math.Inf(-1))
}

func (cp *ComputerPlayer) minimax(node *stree.Node, isMaximizing bool, alpha float64, beta float64) float64 {
	if len(node.Children) == 0 {
		node.Score = int(node.B.CalculateScore())
		return float64(node.Score)
	}

	if isMaximizing {
		bestValue := math.Inf(-1)
		for _, child := range node.Children {
			value := cp.minimax(child, false, alpha, beta)
			bestValue = math.Max(value, bestValue)
			alpha = math.Max(alpha, bestValue)
			if beta <= alpha {
				break
			}
		}
		node.Score = int(bestValue)
		return bestValue
	} else {
		bestValue := math.Inf(1)
		for _, child := range node.Children {
			value := cp.minimax(child, true, alpha, beta)
			bestValue = math.Min(value, bestValue)
			beta = math.Min(beta, bestValue)
			if beta <= alpha {
				break
			}
		}
		node.Score = int(bestValue)
		return bestValue
	}

	return float64(node.Score)
}

func (cp *ComputerPlayer) NextMove(board *brd.Board) *constants.PossibleMove {
	startTime := time.Now()
	stateTree := cp.constructStateTree(7, board)
	fmt.Println(time.Since(startTime))
	startTime = time.Now()
	bestScore := cp.minimax(stateTree.Root, true, math.Inf(-1), math.Inf(1))
	fmt.Println(time.Since(startTime))

	for _, child := range stateTree.Root.Children {
		if child.Score == int(bestScore) {
			return child.M
		}
	}

	return nil
}

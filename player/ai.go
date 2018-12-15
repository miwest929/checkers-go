package player

import (
	brd "checkers-go/board"
	"checkers-go/constants"
	"checkers-go/queue"
	stree "checkers-go/statetree"
	"fmt"
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
			numberOfStates += 1
			state.Children = append(state.Children, nextState)
			q.Enqueue(nextState)
		}
	}

	fmt.Printf("Generated %d states\n", numberOfStates)

	return tree
}

type MinimaxFn func([]int) int

func (cp *ComputerPlayer) minimaxMax(node *stree.Node) int {
	return cp.minimax(node, GetMax, GetMin)
}

func (cp *ComputerPlayer) minimaxMin(node *stree.Node) int {
	return cp.minimax(node, GetMin, GetMax)
}

func (cp *ComputerPlayer) minimax(node *stree.Node, deciderFn MinimaxFn, opponentDecider MinimaxFn) int {
	if len(node.Children) == 0 {
		node.Score = node.B.CalculateScore()
		return node.Score
	}

	scores := make([]int, 0)
	for _, child := range node.Children {
		scores = append(scores, cp.minimax(child, opponentDecider, deciderFn))
	}

	node.Score = deciderFn(scores)
	return node.Score
}

func GetMax(scores []int) int {
	maxValue := scores[0]
	for _, value := range scores[1:] {
		if value > maxValue {
			maxValue = value
		}
	}

	return maxValue
}

func GetMin(scores []int) int {
	minValue := scores[0]
	for _, value := range scores[1:] {
		if value < minValue {
			minValue = value
		}
	}

	return minValue
}

func (cp *ComputerPlayer) NextMove(board *brd.Board) *constants.PossibleMove {
	startTime := time.Now()
	stateTree := cp.constructStateTree(7, board)
	fmt.Println(time.Since(startTime))
	startTime = time.Now()
	bestScore := cp.minimaxMax(stateTree.Root)
	fmt.Println(time.Since(startTime))

	for _, child := range stateTree.Root.Children {
		if child.Score == bestScore {
			return child.M
		}
	}

	return nil
}

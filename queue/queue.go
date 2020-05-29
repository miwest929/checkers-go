// Package queue creates a ItemQueue data structure for the Item type
package queue

import (
	stree "github.com/miwest929/checkers-go/statetree"
	"sync"
)

// Item the type of the queue
type Item *stree.Node

// ItemQueue the queue of Items
type ItemQueue struct {
	nodes []*stree.Node
	lock  sync.RWMutex
}

// New creates a new ItemQueue
func NewQueue() *ItemQueue {
	return &ItemQueue{nodes: make([]*stree.Node, 0)}
}

// Enqueue adds an Item to the end of the queue
func (s *ItemQueue) Enqueue(t *stree.Node) {
	s.lock.Lock()
	s.nodes = append(s.nodes, t)
	s.lock.Unlock()
}

// Dequeue removes an Item from the start of the queue
func (s *ItemQueue) Dequeue() *stree.Node {
	s.lock.Lock()
	node := s.nodes[0]
	s.nodes = s.nodes[1:len(s.nodes)]
	s.lock.Unlock()
	return node
}

// Front returns the item next in the queue, without removing it
func (s *ItemQueue) Front() *stree.Node {
	s.lock.RLock()
	node := s.nodes[0]
	s.lock.RUnlock()
	return node
}

// IsEmpty returns true if the queue is empty
func (s *ItemQueue) IsEmpty() bool {
	return len(s.nodes) == 0
}

// Size returns the number of Items in the queue
func (s *ItemQueue) Size() int {
	return len(s.nodes)
}

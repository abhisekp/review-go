package node

import "sync"

type Node[T any] struct {
	sync.RWMutex
	value T
	next  *Node[T]
	prev  *Node[T]
}

func (n *Node[T]) Value() T {
	return n.value
}

func (n *Node[T]) Next() *Node[T] {
	return n.next
}

func (n *Node[T]) Prev() *Node[T] {
	return n.prev
}

func (n *Node[T]) SetNext(next *Node[T]) {
	n.Lock()
	defer n.Unlock()
	n.next = next
}

func (n *Node[T]) SetPrev(prev *Node[T]) {
	n.Lock()
	defer n.Unlock()
	n.prev = prev
}

func NewNode[T any](value T) *Node[T] {
	return &Node[T]{
		value: value,
	}
}

func (n *Node[T]) SetValue(value T) {
	n.Lock()
	defer n.Unlock()
	n.value = value
}

package linkedlist

import (
	"fmt"
	"iter"
	. "review-go/node"
	"slices"
	"sync"
	"sync/atomic"
)

type LinkedList[T any] struct {
	sync.RWMutex
	head *Node[T]
	tail *Node[T]
	size atomic.Int64
}

func NewLinkedList[T any](data ...T) *LinkedList[T] {
	ll := LinkedList[T]{}
	if len(data) > 0 {
		ll.Add(data...)
	}
	return &ll
}

func (ll *LinkedList[T]) Add(data ...T) {
	for _, d := range data {
		node := NewNode(d)
		ll.Lock()
		if ll.head == nil {
			ll.head = node
			ll.tail = ll.head
		} else {
			ll.tail.SetNext(node)
			ll.tail = node
		}
		ll.Unlock()
		ll.size.Add(1)
	}
}

func (ll *LinkedList[T]) Size() int {
	return int(ll.size.Load())
}

func (ll *LinkedList[T]) IsEmpty() bool {
	return ll.Size() == 0
}

func (ll *LinkedList[T]) SetHead(head *Node[T]) {
	ll.Lock()
	defer ll.Unlock()
	ll.head = head
}

func (ll *LinkedList[T]) SetTail(tail *Node[T]) {
	ll.Lock()
	defer ll.Unlock()
	ll.tail = tail
}

func (ll *LinkedList[T]) Head() *Node[T] {
	ll.RLock()
	defer ll.RUnlock()
	return ll.head
}

func (ll *LinkedList[T]) Tail() *Node[T] {
	ll.RLock()
	defer ll.RUnlock()
	return ll.tail
}

func (ll *LinkedList[T]) Values() iter.Seq[T] {
	return func(yield func(T) bool) {
		for n := range ll.Nodes() {
			if !yield(n.Value()) {
				return
			}
		}
	}
}

func (ll *LinkedList[T]) String() string {
	return fmt.Sprintf("%+v", slices.Collect(ll.Values()))
}

func (ll *LinkedList[T]) Remove(node *Node[T]) {
	ll.Lock()
	defer ll.Unlock()
	defer ll.size.Add(-1)
	if node == ll.head {
		ll.head = ll.head.Next()
		return
	}

	for n := ll.head; n != nil; n = n.Next() {
		if n.Next() == node {
			n.SetNext(node.Next())
			if node == ll.tail {
				ll.tail = n
			}
			return
		}
	}
}

func (ll *LinkedList[T]) InsertAfter(node *Node[T], data T) {
	ll.Lock()
	defer ll.Unlock()
	defer ll.size.Add(1)

	newNode := NewNode(data)
	newNode.SetNext(node.Next())
	node.SetNext(newNode)
	if ll.tail == node {
		ll.tail = newNode
	}
}

func (ll *LinkedList[T]) InsertBefore(node *Node[T], data T) {
	ll.Lock()
	defer ll.Unlock()
	defer ll.size.Add(1)

	newNode := NewNode(data)
	newNode.SetNext(node)

	if ll.head == node {
		ll.head = newNode
		return
	}

	for n := ll.head; n != nil; n = n.Next() {
		if n.Next() == node {
			n.SetNext(newNode)
			return
		}
	}
}

func (ll *LinkedList[T]) Nodes() iter.Seq[*Node[T]] {
	return func(yield func(*Node[T]) bool) {
		for n := ll.head; n != nil; n = n.Next() {
			if !yield(n) {
				return
			}
		}
	}
}

func (ll *LinkedList[T]) Clone() *LinkedList[T] {
	ll.RLock()
	defer ll.RUnlock()
	return NewLinkedList[T](slices.Collect(ll.Values())...)
}

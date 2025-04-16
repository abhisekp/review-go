package stack

import (
	"iter"
	"review-go/queue"
)

type Stack[T any] struct {
	queue1 *queue.Queue[T]
	queue2 *queue.Queue[T]
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		queue1: queue.NewQueue[T](),
		queue2: queue.NewQueue[T](),
	}
}

func (s *Stack[T]) add(data T) {
	if s.queue1.IsEmpty() {
		s.queue1.Enqueue(data)
	} else {
		for d := range s.queue1.Values() {
			s.queue2.Enqueue(d)
			s.queue1.Dequeue()
		}
		s.queue1.Enqueue(data)
		for d := range s.queue2.Values() {
			s.queue1.Enqueue(d)
			s.queue2.Dequeue()
		}
	}
}

func (s *Stack[T]) Push(values ...T) {
	for _, d := range values {
		s.add(d)
	}
}

func (s *Stack[T]) Pop() bool {
	return s.queue1.Dequeue()
}

func (s *Stack[T]) Peek() (T, bool) {
	var empty T
	if s.IsEmpty() {
		return empty, false
	}
	return s.queue1.Peek()
}

func (s *Stack[T]) IsEmpty() bool {
	return s.queue1.IsEmpty()
}

func (s *Stack[T]) Values() iter.Seq[T] {
	return s.queue1.Values()
}

func (s *Stack[T]) Size() int {
	return s.queue1.Size()
}

package queue

import (
	"iter"
	"review-go/linkedlist"
	"review-go/utils"
)

type Queue[T any] struct {
	data *linkedlist.LinkedList[T]
}

type QueueOption struct {
}

func NewQueue[T any](options ...QueueOption) *Queue[T] {
	option := utils.TakeFirst(QueueOption{}, options)
	_ = option
	return &Queue[T]{
		data: linkedlist.NewLinkedList[T](),
	}
}

func (q *Queue[T]) Enqueue(data ...T) {
	q.data.Add(data...)
}

func (q *Queue[T]) Dequeue() bool {
	if q.IsEmpty() {
		return false
	}
	q.data.Remove(q.data.Head())
	return true
}

func (q *Queue[T]) Peek() (T, bool) {
	var empty T
	if q.IsEmpty() {
		return empty, false
	}
	val := q.data.Head()
	if val != nil {
		return val.Value(), true
	}
	return empty, false
}

func (q *Queue[T]) IsEmpty() bool {
	return q.data.IsEmpty()
}

func (q *Queue[T]) Values() iter.Seq[T] {
	return q.data.Values()
}

func (q *Queue[T]) Size() int {
	return q.data.Size()
}

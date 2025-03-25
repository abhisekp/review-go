package queue

import (
	"iter"
	"review-go/utils"
	"slices"
)

type Queue[T any] struct {
	data []T
}

type QueueOption struct {
	Size int
}

func NewQueue[T any](options ...QueueOption) *Queue[T] {
	option := utils.TakeFirst(QueueOption{}, options)
	size := 10
	if option.Size != 0 {
		size = option.Size
	}
	return &Queue[T]{
		data: make([]T, 0, size),
	}
}

func (q *Queue[T]) Enqueue(data ...T) {
	q.data = append(q.data, data...)
}

func (q *Queue[T]) Dequeue() bool {
	if q.IsEmpty() {
		return false
	}
	q.data = q.data[1:]
	return true
}

func (q *Queue[T]) Peek() (T, bool) {
	var empty T
	if q.IsEmpty() {
		return empty, false
	}
	return q.data[0], true
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.data) == 0
}

func (q *Queue[T]) Values() iter.Seq[T] {
	return slices.Values(q.data)
}

func (q *Queue[T]) Size() int {
	return len(q.data)
}

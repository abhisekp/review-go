package bfs

import (
	"review-go/queue"
	"review-go/set"
)

func ShortestPath[T comparable](graph map[T][]T, start T, traverseFn func(T) bool) (T, int, bool) {
	visited := set.NewSet[T]()

	pathQueue := queue.NewQueue[T]()

	level := 0

	pathQueue.Enqueue(graph[start]...)
	if traverseFn(start) {
		return start, level, true
	}

	for !pathQueue.IsEmpty() {
		size := pathQueue.Size()
		for i := 0; i < size; i++ {
			node, _ := pathQueue.Peek()
			pathQueue.Dequeue()
			if visited.Contains(node) {
				continue
			}
			visited.Add(node)
			pathQueue.Enqueue(graph[node]...)
			if traverseFn(node) {
				return node, level + 1, true
			}
		}
		level++
	}

	return start, level, false
}

type Person struct {
	ID     string
	Name   string
	Seller bool
}

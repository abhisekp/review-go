package bfs

import (
	"fmt"
	"review-go/queue"
	"review-go/set"
	"review-go/stackFromQueue"
	"slices"
	"strings"
)

func ShortestPath[T comparable](graph map[T][]T, start T, traverseFn func(T) bool) (T, int, bool) {
	visited := set.NewSet[T]()

	pathQueue := queue.NewQueue[T]()

	levelStack := stack.NewStack[T]()
	levelStack.Push(start)

	level := 0

	pathQueue.Enqueue(graph[start]...)
	if traverseFn(start) {
		return start, level, true
	}

	for !pathQueue.IsEmpty() {
		currLevelSize := pathQueue.Size()
		for i := 0; i < currLevelSize; i++ {
			if node, ok := pathQueue.Peek(); ok {
				pathQueue.Dequeue()
				if visited.Contains(node) {
					continue
				}
				visited.Add(node)
				pathQueue.Enqueue(graph[node]...)
				levelStack.Push(node)
				if traverseFn(node) {
					// Print Level Stack
					stackValues := slices.Collect(levelStack.Values())
					msg := strings.Builder{}
					for j, v := range stackValues {
						if j == len(stackValues)-1 {
							msg.WriteString(fmt.Sprintf("%v", v))
						} else {
							msg.WriteString(fmt.Sprintf("%v -> ", v))
						}
					}
					fmt.Println(msg.String())

					return node, level + 1, true
				}
			}
		}
		levelStack.Pop()
		level++
	}

	return start, level, false
}

type Person struct {
	ID     string
	Name   string
	Seller bool
}

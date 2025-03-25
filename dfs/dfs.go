package dfs

import (
	"fmt"
	"review-go/queue"
	"review-go/set"
)

type FileNode struct {
	IsFile      bool
	Name        string
	IsDirectory bool
	Children    []FileNode
}

var _ fmt.Stringer = (*FileNode)(nil)

func (fn FileNode) String() string {
	id := ""
	if fn.IsDirectory {
		id = fmt.Sprintf("%s/", fn.Name)
	} else {
		id = fmt.Sprintf("%s", fn.Name)
	}
	return id
}

type DFS[T any] struct {
	queue   *queue.Queue[T]
	visited *set.Set[string]
}

func NewDFS[T any]() *DFS[T] {
	return &DFS[T]{
		queue:   queue.NewQueue[T](),
		visited: set.NewSet[string](),
	}
}

func (d *DFS[T]) Run(graph []T, parent *T, traverse func(node T, ID string) (id string, bail bool)) (int, bool) {
	return 0, false
}

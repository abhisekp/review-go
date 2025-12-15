# queue

Generic FIFO queue implemented using the `linkedlist` package. Provides `Enqueue`, `Dequeue`, `Peek` and iteration helpers.

Files

- `queue.go` — queue implementation.

Usage

- Create: `q := queue.NewQueue[int]()`, then `q.Enqueue(1,2)`, `q.Dequeue()`, `q.Peek()`.


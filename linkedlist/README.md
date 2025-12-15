# linkedlist

Thread-safe generic linked list implementation providing add/remove/insert and iterator-style `Values` and `Nodes`.

Files

- `linkedlist.go` — main implementation.

Usage

- Import `review-go/linkedlist` and create a new list: `ll := linkedlist.NewLinkedList[int](1,2,3)`.
- Use `ll.Add`, `ll.Remove`, `ll.InsertAfter`, `ll.Values()` for iteration.


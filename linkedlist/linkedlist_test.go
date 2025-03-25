package linkedlist

import (
	"iter"
	"review-go/utils"
	"testing"
)

func TestNewLinkedList(t *testing.T) {
	t.Run("NewLinkedList", func(t *testing.T) {
		ll := NewLinkedList[int]()
		if ll.head != nil {
			t.Errorf("Expected head to be nil, got %v", ll.head)
		}

		t.Run("NewLinkedList with elements", func(t *testing.T) {
			ll := NewLinkedList[int](1, 2, 3)
			if ll.head.Value() != 1 {
				t.Errorf("Expected head value to be 1, got %v", ll.head.Value())
			}
			if ll.tail.Value() != 3 {
				t.Errorf("Expected tail value to be 3, got %v", ll.tail.Value())
			}
		})
	})

	t.Run("Add", func(t *testing.T) {
		ll := NewLinkedList[int]()
		ll.Add(1, 2, 3)
		if ll.head.Value() != 1 {
			t.Errorf("Expected head value to be 1, got %v", ll.head.Value())
		}
		if ll.tail.Value() != 3 {
			t.Errorf("Expected tail value to be 3, got %v", ll.tail.Value())
		}
	})

	t.Run("Size", func(t *testing.T) {
		ll := NewLinkedList[int](1, 2, 3)
		if ll.Size() != 3 {
			t.Errorf("Expected size to be 3, got %d", ll.Size())
		}
	})
	t.Run("IsEmpty", func(t *testing.T) {
		ll := NewLinkedList[int]()
		if !ll.IsEmpty() {
			t.Errorf("Expected list to be empty, got %v", ll)
		}

		ll.Add(1, 2, 3)
		if ll.IsEmpty() {
			t.Errorf("Expected list to not be empty, got %v", ll)
		}
	})

	t.Run("Contains", func(t *testing.T) {
		ll := NewLinkedList[int](1, 2, 3)
		if !Contains(ll, utils.IsEqual, 2) {
			t.Errorf("Expected list to contain 2, got %v", ll)
		}
		if Contains(ll, utils.IsEqual, 4) {
			t.Errorf("Expected list to not contain 4, got %v", ll)
		}
	})

	t.Run("String", func(t *testing.T) {
		ll := NewLinkedList[int]()
		ll.Add(1, 2, 3)
		if ll.String() != "[1 2 3]" {
			t.Errorf("Expected string representation to be [1 2 3], got %v", ll.String())
		}
	})

	t.Run("Find", func(t *testing.T) {
		ll := NewLinkedList(1, 2, 3)
		if Find(ll, utils.IsEqual, 2) != ll.head.Next() {
			t.Errorf("Expected to find 2, got %v", Find(ll, utils.IsEqual, 2))
		}
	})

	t.Run("Delete", func(t *testing.T) {
		ll := NewLinkedList(1, 2, 3)
		node1 := Find(ll, utils.IsEqual, 2)
		ll.Remove(node1)
		if Find(ll, utils.IsEqual, 2) != nil {
			t.Errorf("Expected to delete 2, got %v", Find(ll, utils.IsEqual, 2))
		}
		if ll.Size() != 2 {
			t.Errorf("Expected size to be 2, got %d", ll.Size())
		}
		if ll.String() != "[1 3]" {
			t.Errorf("Expected string representation to be [1 3], got %v", ll.String())
		}

		t.Run("Delete last element", func(t *testing.T) {
			ll := NewLinkedList(1, 2, 3)
			node1 := Find(ll, utils.IsEqual, 3)
			ll.Remove(node1)
			if Find(ll, utils.IsEqual, 3) != nil {
				t.Errorf("Expected to delete 3, got %v", Find(ll, utils.IsEqual, 2))
			}
			if ll.Size() != 2 {
				t.Errorf("Expected size to be 2, got %d", ll.Size())
			}
			if ll.String() != "[1 2]" {
				t.Errorf("Expected string representation to be [1 2], got %v", ll.String())
			}
		})

		t.Run("Delete last element", func(t *testing.T) {
			ll := NewLinkedList(1, 2, 3)
			node1 := Find(ll, utils.IsEqual, 1)
			ll.Remove(node1)
			if Find(ll, utils.IsEqual, 1) != nil {
				t.Errorf("Expected to delete 1, got %v", Find(ll, utils.IsEqual, 2))
			}
			if ll.Size() != 2 {
				t.Errorf("Expected size to be 2, got %d", ll.Size())
			}
			if ll.String() != "[2 3]" {
				t.Errorf("Expected string representation to be [2 3], got %v", ll.String())
			}
		})
	})

	t.Run("InsertAfter", func(t *testing.T) {
		ll := NewLinkedList(1, 2, 3)
		node1 := Find(ll, utils.IsEqual, 2)
		ll.InsertAfter(node1, 4)
		if Find(ll, utils.IsEqual, 4) == nil {
			t.Errorf("Expected to insert 4 after 2, got %v", Find(ll, utils.IsEqual, 4))
		}
		if ll.Size() != 4 {
			t.Errorf("Expected size to be 4, got %d", ll.Size())
		}
		if ll.String() != "[1 2 4 3]" {
			t.Errorf("Expected string representation to be [1 2 4 3], got %v", ll.String())
		}

		t.Run("InsertAfter at the head", func(t *testing.T) {
			ll := NewLinkedList(1, 2, 3)
			node1 := Find(ll, utils.IsEqual, 1)
			ll.InsertAfter(node1, 4)
			if Find(ll, utils.IsEqual, 4) == nil {
				t.Errorf("Expected to insert 4 after 1, got %v", Find(ll, utils.IsEqual, 4))
			}
			if ll.Size() != 4 {
				t.Errorf("Expected size to be 4, got %d", ll.Size())
			}
			if ll.String() != "[1 4 2 3]" {
				t.Errorf("Expected string representation to be [1 4 2 3], got %v", ll.String())
			}
		})

		t.Run("InsertAfter at the tail", func(t *testing.T) {
			ll := NewLinkedList(1, 2, 3)
			node1 := Find(ll, utils.IsEqual, 3)
			ll.InsertAfter(node1, 4)
			if Find(ll, utils.IsEqual, 4) == nil {
				t.Errorf("Expected to insert 4 after 3, got %v", Find(ll, utils.IsEqual, 4))
			}
			if ll.Size() != 4 {
				t.Errorf("Expected size to be 4, got %d", ll.Size())
			}
			if ll.String() != "[1 2 3 4]" {
				t.Errorf("Expected string representation to be [1 2 3 4], got %v", ll.String())
			}
		})

	})

	t.Run("InsertBefore", func(t *testing.T) {
		ll := NewLinkedList(1, 2, 3)
		node1 := Find(ll, utils.IsEqual, 2)
		ll.InsertBefore(node1, 4)
		if Find(ll, utils.IsEqual, 4) == nil {
			t.Errorf("Expected to insert 4 before 2, got %v", Find(ll, utils.IsEqual, 4))
		}
		if ll.Size() != 4 {
			t.Errorf("Expected size to be 4, got %d", ll.Size())
		}
		if ll.String() != "[1 4 2 3]" {
			t.Errorf("Expected string representation to be [1 4 2 3], got %v", ll.String())
		}

		t.Run("InsertBefore at the head", func(t *testing.T) {
			ll := NewLinkedList(1, 2, 3)
			node1 := Find(ll, utils.IsEqual, 1)
			ll.InsertBefore(node1, 4)
			if Find(ll, utils.IsEqual, 4) == nil {
				t.Errorf("Expected to insert 4 before 1, got %v", Find(ll, utils.IsEqual, 4))
			}
			if ll.Size() != 4 {
				t.Errorf("Expected size to be 4, got %d", ll.Size())
			}
			if ll.String() != "[4 1 2 3]" {
				t.Errorf("Expected string representation to be [4 1 2 3], got %v", ll.String())
			}
		})

		t.Run("InsertBefore at the tail", func(t *testing.T) {
			ll := NewLinkedList(1, 2, 3)
			node1 := Find(ll, utils.IsEqual, 3)
			ll.InsertBefore(node1, 4)
			if Find(ll, utils.IsEqual, 4) == nil {
				t.Errorf("Expected to insert 4 before 3, got %v", Find(ll, utils.IsEqual, 4))
			}
			if ll.Size() != 4 {
				t.Errorf("Expected size to be 4, got %d", ll.Size())
			}
			if ll.String() != "[1 2 4 3]" {
				t.Errorf("Expected string representation to be [1 2 4 3], got %v", ll.String())
			}
		})
	})

	t.Run("Nodes", func(t *testing.T) {
		ll := NewLinkedList(1, 2, 3)
		next, stop := iter.Pull(ll.Nodes())
		node1 := Find(ll, utils.IsEqual, 1)
		node2 := Find(ll, utils.IsEqual, 2)
		node3 := Find(ll, utils.IsEqual, 3)
		if n, ok := next(); ok && n != node1 {
			t.Errorf("Expected first node to be 1, got %v", n)
		}
		if n, ok := next(); ok && n != node2 {
			t.Errorf("Expected second node to be 2, got %v", n)
		}
		if n, ok := next(); ok && n != node3 {
			t.Errorf("Expected third node to be 3, got %v", n)
		}
		stop()
		if n, ok := next(); ok {
			t.Errorf("Expected next to return false, got %v", n)
		}
	})

}

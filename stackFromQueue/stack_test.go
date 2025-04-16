package stack

import "testing"

func TestStack(t *testing.T) {
	t.Run("Add 3 Items", func(t *testing.T) {
		s := NewStack[int]()
		s.Push(1, 2, 3)
		if v, ok := s.Peek(); !ok || v != 3 {
			t.Errorf("Expected %v, got %v", 3, v)
		}
	})

	t.Run("Pop 3 Items", func(t *testing.T) {
		s := NewStack[int]()
		s.Push(1, 2, 3, 4)
		if v, ok := s.Peek(); !ok || v != 4 {
			t.Errorf("Expected %v, got %v", 4, v)
		}
		s.Pop()
		if v, ok := s.Peek(); !ok || v != 3 {
			t.Errorf("Expected %v, got %v", 3, v)
		}
		s.Pop()
		if v, ok := s.Peek(); !ok || v != 2 {
			t.Errorf("Expected %v, got %v", 2, v)
		}
		s.Pop()
		if v, ok := s.Peek(); !ok || v != 1 {
			t.Errorf("Expected %v, got %v", 1, v)
		}

		if s.Size() != 1 {
			t.Errorf("Expected stack size to be 1, got %d", s.Size())
		}
	})
}

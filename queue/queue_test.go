package queue

import (
	"slices"
	"testing"
)

func TestQueue(t *testing.T) {
	t.Run("Enqueue", func(t *testing.T) {
		q := NewQueue[int]()
		q.Enqueue(1, 2, 3, 4, 5)
		expected := []int{1, 2, 3, 4, 5}
		actual := slices.Collect(q.data.Values())
		if !slices.Equal(actual, expected) {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	})

	t.Run("Dequeue", func(t *testing.T) {
		q := NewQueue[int]()
		q.Enqueue(1, 2, 3, 4, 5)
		q.Dequeue()
		expected := []int{2, 3, 4, 5}
		actual := slices.Collect(q.data.Values())
		if !slices.Equal(actual, expected) {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	})

	t.Run("Size", func(t *testing.T) {
		q := NewQueue[int]()
		q.Enqueue(1, 2, 3, 4, 5)
		expected := 5
		actual := q.Size()
		if actual != expected {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	})

	t.Run("IsEmpty", func(t *testing.T) {
		q := NewQueue[int]()
		expected := true
		actual := q.IsEmpty()
		if actual != expected {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	})

	t.Run("Peek", func(t *testing.T) {
		q := NewQueue[int]()
		q.Enqueue(1, 2, 3, 4, 5)
		expected := 1
		actual, _ := q.Peek()
		if actual != expected {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	})

	t.Run("Values", func(t *testing.T) {
		q := NewQueue[int]()
		q.Enqueue(1, 2, 3, 4, 5)
		expected := []int{1, 2, 3, 4, 5}
		actual := slices.Collect(q.Values())
		if !slices.Equal(actual, expected) {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	})
}

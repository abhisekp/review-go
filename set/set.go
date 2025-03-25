package set

import (
	"iter"
	"slices"
)

type empty = struct{}

type Set[T comparable] struct {
	data map[T]empty
}

func NewSet[T comparable](data ...T) *Set[T] {
	s := Set[T]{
		data: map[T]empty{},
	}
	s.Add(data...)
	return &s
}

func (s *Set[T]) Add(data ...T) {
	for _, d := range data {
		s.data[d] = empty{}
	}
}

func (s *Set[T]) Contains(data T) bool {
	_, ok := s.data[data]
	return ok
}

func (s *Set[T]) IsEmpty(data T) bool {
	return len(s.data) == 0
}

func (s *Set[T]) Remove(data T) {
	if s.Contains(data) {
		delete(s.data, data)
	}
}

func (s *Set[T]) Values() iter.Seq[T] {
	return func(yield func(T) bool) {
		for d := range s.data {
			if !yield(d) {
				return
			}
		}
	}
}

func (s *Set[T]) Size() int {
	return len(s.data)
}

func Union[T comparable](s1, s2 *Set[T]) *Set[T] {
	return NewSet(append(slices.Collect(s1.Values()), slices.Collect(s2.Values())...)...)
}

func Intersection[T comparable](s1, s2 *Set[T]) *Set[T] {
	s := NewSet[T]()
	for v := range s1.Values() {
		if s2.Contains(v) {
			s.Add(v)
		}
	}
	return s
}

func Difference[T comparable](s1, s2 *Set[T]) *Set[T] {
	s := NewSet[T]()
	for v := range s1.Values() {
		if !s2.Contains(v) {
			s.Add(v)
		}
	}
	return s
}

func SymmetricDifference[T comparable](s1, s2 *Set[T]) *Set[T] {
	return Difference(Union(s1, s2), Intersection(s1, s2))
}

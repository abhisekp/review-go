package set

import "testing"

func TestUnion(t *testing.T) {
	s1 := NewSet[int](1, 2, 3)
	s2 := NewSet[int](2, 3, 4)
	s3 := Union[int](s1, s2)
	if s3.Size() != 4 {
		t.Errorf("Union should have 4 elements, got %d", s3.Size())
	}
	if !s3.Contains(1) || !s3.Contains(2) || !s3.Contains(3) || !s3.Contains(4) {
		t.Errorf("Union should contain 1, 2, 3, 4, got %v", s3)
	}
}

func TestDifference(t *testing.T) {
	s1 := NewSet[int](1, 2, 3)
	s2 := NewSet[int](2, 3, 4)
	s3 := Difference[int](s1, s2)
	if s3.Size() != 1 {
		t.Errorf("Difference should have 1 element, got %d", s3.Size())
	}
	if !s3.Contains(1) {
		t.Errorf("Difference should contain 1, got %v", s3)
	}
}

func TestIntersection(t *testing.T) {
	s1 := NewSet[int](1, 2, 3)
	s2 := NewSet[int](2, 3, 4)
	s3 := Intersection[int](s1, s2)
	if s3.Size() != 2 {
		t.Errorf("Intersection should have 2 elements, got %d", s3.Size())
	}
	if !s3.Contains(2) || !s3.Contains(3) {
		t.Errorf("Intersection should contain 2, 3, got %v", s3)
	}
	if s3.Contains(1) || s3.Contains(4) {
		t.Errorf("Intersection should not contain 1 or 4, got %v", s3)
	}
}

func TestSymmetricDifference(t *testing.T) {
	s1 := NewSet[int](1, 2, 3)
	s2 := NewSet[int](2, 3, 4)
	s3 := SymmetricDifference[int](s1, s2)
	if s3.Size() != 2 {
		t.Errorf("SymmetricDifference should have 2 elements, got %d", s3.Size())
	}
	if !s3.Contains(1) || !s3.Contains(4) {
		t.Errorf("SymmetricDifference should contain 1, and 4, got %v", s3)
	}
	if s3.Contains(2) || s3.Contains(3) {
		t.Errorf("SymmetricDifference should not contain 2, and 3, got %v", s3)
	}
}

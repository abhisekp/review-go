package recursion

import (
	"math/rand"
	"slices"
	"testing"
)

func TestRecursive(t *testing.T) {
	t.Run("Sum", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		expected := 15
		actual := Sum(nums)
		if actual != expected {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	})

	t.Run("Count", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		expected := 5
		actual := Count(nums)
		if actual != expected {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	})

	t.Run("Max", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		expected := 5
		actual := Max(nums)
		if actual != expected {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	})

	t.Run("BinarySearch", func(t *testing.T) {
		t.Run("Numbers", func(t *testing.T) {

			nums := []int{1, 2, 3, 4, 5}
			expected := 2
			actual := BinarySearch(nums, 3)
			if actual != expected {
				t.Errorf("Expected %v, got %v", expected, actual)
			}
		})

		t.Run("Strings", func(t *testing.T) {
			nums := []string{"a", "b", "c", "d", "e"}
			expected := 2
			actual := BinarySearch(nums, "c")
			if actual != expected {
				t.Errorf("Expected %v, got %v", expected, actual)
			}
		})
	})

	t.Run("QuickSort", func(t *testing.T) {
		t.Run("QuickSort1", func(t *testing.T) {
			nums := []int{5, 4, 15, 45, 42, 14, 45}
			expected := []int{4, 5, 14, 15, 42, 45, 45}
			actual := QuickSort(nums)
			if !slices.Equal(actual, expected) {
				t.Errorf("Expected %v, got %v", expected, actual)
			}
		})

		t.Run("QuickSort2", func(t *testing.T) {
			nums := []int{5, 4, 15, 45, 42, 14, 45}
			expected := []int{4, 5, 14, 15, 42, 45, 45}
			actual := QuickSort2(nums)
			if !slices.Equal(actual, expected) {
				t.Errorf("Expected %v, got %v", expected, actual)
			}
		})
	})
}

func BenchmarkQuickSort(b *testing.B) {
	N := b.N

	arr1 := make([]int, 0, N)

	for range N {
		arr1 = append(arr1, rand.Intn(100_000_000))
	}

	arr2 := slices.Clone(arr1)

	for range N {
		arr2 = append(arr2, rand.Intn(100_000_000))
	}

	b.Run("QuickSort", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			QuickSort(arr1)
		}
	})

	b.Run("QuickSort2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			QuickSort2(arr2)
		}
	})
}

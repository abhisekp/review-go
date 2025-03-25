package utils

import (
	"slices"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	t.Run("SelectionSort", func(t *testing.T) {
		arr := []int{5, 3, 1, 2, 4}
		expected := []int{1, 2, 3, 4, 5}
		actual := SelectionSort(arr, IsLess[int])
		if !slices.Equal(actual, expected) {
			t.Errorf("Expected %v, got %v", expected, arr)
		}
		if !slices.Equal(arr, []int{5, 3, 1, 2, 4}) {
			t.Errorf("Should not modify original array. Expected %v, got %v", []int{5, 3, 1, 2, 4}, arr)
		}
	})

	t.Run("SelectionSort with custom comparison function", func(t *testing.T) {
		arr := []int{5, 3, 1, 2, 4}
		expected := []int{5, 4, 3, 2, 1}
		actual := SelectionSort(arr, IsGreater[int])
		if !slices.Equal(actual, expected) {
			t.Errorf("Expected %v, got %v", expected, arr)
		}
		if !slices.Equal(arr, []int{5, 3, 1, 2, 4}) {
			t.Errorf("Should not modify original array. Expected %v, got %v", []int{5, 3, 1, 2, 4}, arr)
		}
	})

	t.Run("SelectionSortInPlace", func(t *testing.T) {
		t.Run("Ascending", func(t *testing.T) {

			arr := []int{5, 3, 1, 2, 4}
			expected := []int{1, 2, 3, 4, 5}
			actual := slices.Collect(SelectionSortInPlace(arr, IsLess[int]))

			if !slices.Equal(arr, expected) {
				t.Errorf("Expected %v, got %v", expected, arr)
			}

			if !slices.Equal(actual, expected) {
				t.Errorf("Expected %v, got %v", expected, actual)
			}
		})

		t.Run("Descending", func(t *testing.T) {
			arr := []int{5, 3, 1, 2, 4}
			expected := []int{5, 4, 3, 2, 1}
			actual := slices.Collect(SelectionSortInPlace(arr, IsGreater[int]))

			if !slices.Equal(arr, expected) {
				t.Errorf("Expected %v, got %v", expected, arr)
			}

			if !slices.Equal(actual, expected) {
				t.Errorf("Expected %v, got %v", expected, actual)
			}
		})

		t.Run("Descending with repeated numbers", func(t *testing.T) {
			arr := []int{5, 3, 3, 1, 2, 4}
			expected := []int{5, 4, 3, 3, 2, 1}
			actual := slices.Collect(SelectionSortInPlace(arr, IsGreater[int]))

			if !slices.Equal(arr, expected) {
				t.Errorf("Expected %v, got %v", expected, arr)
			}
			if !slices.Equal(actual, expected) {
				t.Errorf("Expected %v, got %v", expected, actual)
			}
		})
	})
}

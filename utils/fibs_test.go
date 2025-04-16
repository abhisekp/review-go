package utils

import (
	"slices"
	"testing"
)

func TestFibs(t *testing.T) {
	t.Run("Nnacci", func(t *testing.T) {
		t.Run("with [0, 1, 1]", func(t *testing.T) {
			actual := slices.Collect(Take(10, Nnacci()))

			expected := []int{0, 1, 1, 2, 4, 7, 13, 24, 44, 81}

			if !slices.Equal(actual, expected) {
				t.Errorf("Expected %v, got %v", expected, actual)
			}
		})

		t.Run("with [1, 2, 3]", func(t *testing.T) {
			actual := slices.Collect(Take(15, Nnacci([]int{1, 2, 3})))

			expected := []int{1, 2, 3, 6, 11, 20, 37, 68, 125, 230, 423, 778, 1431, 2632, 4841}

			if !slices.Equal(actual, expected) {
				t.Errorf("Expected %v, got %v", expected, actual)
			}
		})

		t.Run("with [123, 456, 789]", func(t *testing.T) {
			actual := slices.Collect(Take(20, Nnacci([]int{123, 456, 789})))

			expected := []int{123, 456, 789, 1368, 2613, 4770, 8751, 16134, 29655, 54540, 100329, 184524, 339393, 624246, 1148163, 2111802, 3884211, 7144176, 13140189, 24168576}

			if !slices.Equal(actual, expected) {
				t.Errorf("Expected %v, got %v", expected, actual)
			}
		})
	})
}

package utils

import (
	"iter"
	"slices"
)

func SelectionSort[T Ordered](arr []T, cmpFn func(T, T) bool) []T {
	result := make([]T, 0, len(arr))

	newArr := slices.Clone(arr)

	size := len(newArr)

	for range size {
		maxItemIdx := 0
		maxItem := newArr[maxItemIdx]

		for i, item := range newArr {
			if cmpFn(item, maxItem) {
				maxItemIdx = i
				maxItem = newArr[maxItemIdx]
			}
		}
		result = append(result, maxItem)
		// Delete the max item
		newArr = append(newArr[:maxItemIdx], newArr[maxItemIdx+1:]...)
	}

	return result
}

func SelectionSortInPlace[T Ordered](arr []T, cmpFn func(T, T) bool) iter.Seq[T] {
	swap := func(arr []T, i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	}

	return func(yield func(T) bool) {
		for i := range len(arr) {
			maxItemIdx := i
			maxItem := arr[maxItemIdx]

			for j := i; j < len(arr); j++ {
				item := arr[j]
				if cmpFn(item, maxItem) {
					maxItemIdx = j
					maxItem = arr[maxItemIdx]
					swap(arr, i, maxItemIdx)
				}
			}

			if !yield(arr[i]) {
				return
			}
		}
	}
}

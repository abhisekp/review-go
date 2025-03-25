package recursion

import (
	"cmp"
	"review-go/utils"
	"sync"
)

func Sum[T cmp.Ordered](nums []T) T {
	var empty T
	if len(nums) == 0 {
		return empty
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return nums[0] + nums[1]
	}
	return nums[0] + Sum[T](nums[1:])
}

func Count[T any](arr []T) int {
	if len(arr) == 0 {
		return 0
	}
	return 1 + Count[T](arr[1:])
}

func Max[T cmp.Ordered](nums []T) T {
	var empty T
	if len(nums) == 0 {
		return empty
	}
	return max(nums[0], Max[T](nums[1:]))
}

func BinarySearch[T cmp.Ordered](arr []T, search T) int {
	var binarySearchRecursive func(arr []T, search T, low, high int) int
	binarySearchRecursive = func(arr []T, search T, low, high int) int {
		if low > high {
			return -1
		}
		mid := low + (high-low)/2
		if arr[mid] == search {
			return mid
		}
		if search > arr[mid] {
			return binarySearchRecursive(arr, search, mid+1, high)
		}
		return binarySearchRecursive(arr, search, low, mid-1)
	}
	return binarySearchRecursive(arr, search, 0, len(arr)-1)
}

func QuickSort[T utils.Ordered](arr []T) []T {
	if len(arr) < 2 {
		return arr
	}
	if len(arr) == 2 {
		return []T{min(arr[0], arr[1]), max(arr[0], arr[1])}
	}
	pivot := arr[len(arr)/2]
	lessArr, equalArr, greaterArr := utils.PartitionFunc2[T](arr, func(num T, i int) T {
		return num - pivot
	})

	return append(append(QuickSort(lessArr), equalArr...), QuickSort(greaterArr)...)
}

func QuickSort2[T utils.Ordered](arr []T) []T {
	if len(arr) < 2 {
		return arr
	}
	if len(arr) == 2 {
		return []T{min(arr[0], arr[1]), max(arr[0], arr[1])}
	}
	pivot := arr[len(arr)/2]
	lessArr, equalArr, greaterArr := utils.PartitionFunc2[T](arr, func(num T, i int) T {
		return num - pivot
	})

	var leftWG, rightWG sync.WaitGroup
	var leftSort, rightSort []T

	leftWG.Add(1)
	go func() {
		leftSort = QuickSort2(lessArr)
		leftWG.Done()
	}()

	rightWG.Add(1)
	go func() {
		rightSort = QuickSort2(greaterArr)
		rightWG.Done()
	}()

	leftWG.Wait()
	rightWG.Wait()

	return append(append(leftSort, equalArr...), rightSort...)
}

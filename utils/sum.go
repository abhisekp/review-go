package utils

import (
	"cmp"
	"iter"
)

func Sum[T cmp.Ordered](nums iter.Seq[T]) T {
	var empty T
	sum := empty

	for num := range nums {
		sum += num
	}

	return sum
}

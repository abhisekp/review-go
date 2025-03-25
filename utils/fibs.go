package utils

import (
	"iter"
	"math"
)

func Fibs[T int]() iter.Seq[T] {
	last1 := 0
	last2 := 1
	return func(yield func(T) bool) {
		for i := range math.MaxInt {
			if i == 0 || i == 1 {
				if !yield(T(i)) {
					return
				}
			} else {
				if !yield(T(last1 + last2)) {
					return
				}
				last1, last2 = last2, last1+last2
			}

		}
	}
}

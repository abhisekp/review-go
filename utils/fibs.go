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

// Nnacci n-nacci sequence
func Nnacci[T int](inits ...[]T) iter.Seq[T] {
	initial := []T{0, 1, 1}
	if len(inits) >= 1 {
		initial = inits[0]
	}

	return func(yield func(T) bool) {
		currSum := T(0)

		last := initial[0]

		for i, j := 0, 0; ; i = i + 1 {
			if i < len(initial) {
				curr := initial[i]
				currSum += curr
				if !yield(curr) {
					return
				}
			} else {
				if !yield(currSum) {
					return
				}
				currSum -= last
				j++
				if j < len(initial) {
					// last =
				} else {
					last = initial[j]
				}
			}
		}
	}
}

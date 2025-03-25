package utils

import "iter"

func Take[T int](upto T, seq iter.Seq[T]) iter.Seq[T] {
	next, stop := iter.Pull(seq)
	return func(yield func(T) bool) {
		for range upto {
			if nextVal, ok := next(); ok {
				if !yield(nextVal) {
					stop()
					return
				}
			}
		}
	}
}

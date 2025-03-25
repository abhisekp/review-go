package utils

import (
	"iter"
	"math"
)

func Primes[T int]() iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := range math.MaxInt {
			if IsPrime(i) {
				if !yield(T(i)) {
					return
				}
			}
		}
	}
}

var allPrimes = map[int]bool{
	2: true, 3: true, 5: true,
	7: true, 11: true, 13: true,
	17: true, 19: true, 23: true,
	29: true, 31: true, 37: true,
	39: true, 41: true, 43: true,
	47: true, 51: true, 53: true,
	59: true, 61: true, 67: true,
	71: true, 73: true, 79: true,
	83: true, 87: true, 91: true,
	93: true, 97: true, 101: true,
}

func IsPrime[T int](num T) bool {
	if num < 2 {
		return false
	}
	if _, ok := allPrimes[int(num)]; ok {
		return true
	}
	for i := range allPrimes {
		if int(num)%i == 0 {
			return false
		}
	}
	allPrimes[int(num)] = true
	return true
}

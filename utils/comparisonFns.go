package utils

func IsGreater[T Ordered](a, b T) bool { return a > b }

func IsGreaterThanEqual[T Ordered](a, b T) bool { return a >= b }

func IsLess[T Ordered](a, b T) bool { return a < b }

func IsLessThanEqual[T Ordered](a, b T) bool { return a <= b }

func IsEqual[T comparable](a, b T) bool {
	return a == b
}

package utils

import (
	"cmp"
)

type PartFn[T cmp.Ordered] func(T, int) bool
type PartFn2[T cmp.Ordered] func(T, int) T

func PartitionFunc[T cmp.Ordered](arr []T, fn PartFn[T]) ([]T, []T) {
	truthyArr := make([]T, 0, len(arr))
	falsyArr := make([]T, 0, len(arr))
	for i, v := range arr {
		if fn(v, i) {
			truthyArr = append(truthyArr, v)
		} else {
			falsyArr = append(falsyArr, v)
		}
	}
	return truthyArr, falsyArr
}

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64
}

func PartitionFunc2[T Ordered](arr []T, fn PartFn2[T]) ([]T, []T, []T) {
	greaterArr := make([]T, 0, len(arr))
	lesserArr := make([]T, 0, len(arr))
	equalArr := make([]T, 0, len(arr))

	for i, v := range arr {
		result := fn(v, i)
		if result > 0 {
			greaterArr = append(greaterArr, v)
		} else if result < 0 {
			lesserArr = append(lesserArr, v)
		} else {
			equalArr = append(equalArr, v)
		}
	}
	return lesserArr, equalArr, greaterArr
}

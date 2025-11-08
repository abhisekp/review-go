package pipeline

import (
	"fmt"
	"math"
	"strings"
)

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Float interface {
	~float32 | ~float64
}

type Number interface {
	Integer | Float
}

type PipeFn[T Number] func(<-chan T) <-chan T

func wrapNums[T Number](nums []T) <-chan T {
	out := make(chan T)

	go func() {
		defer close(out)
		for _, num := range nums {
			out <- num
		}
	}()

	return out
}

func double[T Number](in <-chan T) <-chan T {
	out := make(chan T)

	go func() {
		defer close(out)
		for num := range in {
			out <- 2 * num
		}
	}()

	return out
}

func filterBy[T Number](cb func(T) bool) PipeFn[T] {
	return func(in <-chan T) <-chan T {
		out := make(chan T)

		go func() {
			defer close(out)

			for num := range in {
				if cb(num) {
					out <- num
				}
			}
		}()

		return out
	}

}

func square[T Number](in <-chan T) <-chan T {
	out := make(chan T)

	go func() {
		defer close(out)

		for num := range in {
			out <- num * num
		}
	}()

	return out
}

func sum[T Number](in <-chan T) <-chan T {
	out := make(chan T)

	go func() {
		defer close(out)

		currSum := T(0)
		for num := range in {
			currSum += num
		}
		out <- currSum
	}()

	return out
}

func log[T Number](in <-chan T) <-chan T {
	out := make(chan T)

	go func() {
		defer close(out)

		allNums := []T{}

		for num := range in {
			out <- num
			allNums = append(allNums, num)
		}

		fmt.Println(allNums)
	}()

	return out
}

func logMsg[T Number](msg string) PipeFn[T] {
	return func(in <-chan T) <-chan T {
		out := make(chan T)

		go func() {
			defer close(out)

			allNums := []T{}

			for num := range in {
				out <- num
				allNums = append(allNums, num)
			}

			fmt.Println(msg, allNums)
		}()

		return out
	}
}

func Pipe[T Number](pipeFns ...PipeFn[T]) PipeFn[T] {
	return func(in <-chan T) <-chan T {
		var out = in
		for _, fn := range pipeFns {
			out = fn(out)
		}

		return out
	}
}

func Run() {
	const header1 = "|| Without Pipe fn ||"

	fmt.Println(header1)
	WithoutPipe()

	const header2 = "|| With Pipe fn ||"
	fmt.Println(strings.Repeat("-", int(math.Max(float64(len(header1)), float64(len(header2))))))

	fmt.Println(header2)
	WithPipe()
}

func WithoutPipe() {
	type T = int
	nums := []T{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numsCh := wrapNums(nums)

	numsCh = log(numsCh)

	// stage 1 - double the numbers
	// doubledCh := double(numsCh)
	//
	// doubledCh = log(doubledCh)

	// stage 2 - filter by odd numbers
	oddCh := filterBy(isOdd[T])(numsCh)

	oddCh = log(oddCh)

	// stage 3 - filter by divisible by 3
	divCh := filterBy(isDivisibleBy[T](3))(oddCh)

	divCh = log(divCh)

	// stage 4 - square the numbers
	squaresCh := square(divCh)

	squaresCh = log(squaresCh)

	// stage 5 - sum the numbers
	sumCh := sum(squaresCh)

	sumCh = log(sumCh)

	// Print sum
	fmt.Println(<-sumCh)
}

func isEven[T Integer](in T) bool {
	return in%2 == 0
}

func isOdd[T Integer](in T) bool {
	return in%2 != 0
}

func isPrime[T Integer](in T) bool {
	panic("Prime function not yet implemented")
}

func isDivisibleBy[T Integer](div T) func(T) bool {
	return func(t T) bool {
		return t%div == 0
	}
}

func Range[T int](min, max T) <-chan T {
	out := make(chan T)

	go func() {
		defer close(out)

		for i := range max - min + 1 {
			out <- i + min
		}
	}()

	return out
}

func WithPipe() {
	type T = int
	// nums := []T{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numsCh := Range(1, 100)

	filterByDiv35 := Pipe[T](
		filterBy(isDivisibleBy[T](3)), logMsg[T]("DivisibleBy 3:"),
		filterBy(isDivisibleBy[T](5)), logMsg[T]("DivisibleBy 5:"),
	)

	ops := append([]PipeFn[T]{}, logMsg[T]("Nums:"),
		// Stage 1
		// Pipe[T](double, logMsg[T]("Double:")),
		// Stage 2
		Pipe[T](filterBy(isOdd[T]), logMsg[T]("Odds:")),
		// Stage 3
		Pipe[T](filterByDiv35, logMsg[T]("Divisible by both 3 and 5:")),
		// Stage 4
		Pipe[T](square, logMsg[T]("Square:")),
		// Stage 5
		Pipe[T](sum, logMsg[T]("Sum:")),
	)

	sumCh := Pipe[T](ops...)(numsCh)

	fmt.Println(<-sumCh)
}

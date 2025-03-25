package ticker

import (
	"iter"
	"time"
)

func Tick(interval time.Duration, times int) iter.Seq2[int, time.Time] {
	// Create a ticker and use it with a iter.Seq
	ticker := time.NewTicker(interval)
	maxTimes := times

	return func(yield func(int, time.Time) bool) {
		for t := range ticker.C {
			if times == 0 || !yield(maxTimes-times, t) {
				ticker.Stop()
				return
			}
			times--
		}
	}
}

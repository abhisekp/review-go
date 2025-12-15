# ticker

`Tick` helper producing an iterator-style sequence of ticks using `time.Ticker` and an `iter.Seq2` producer.

Files

- `ticker.go` — `Tick(interval time.Duration, times int) iter.Seq2[int, time.Time]`.

Usage

- Use the `Tick` function to iterate timed events in an iterator-style pattern.


package ticker

import (
	"fmt"
	"testing"
	"time"
)

func TestTick(t *testing.T) {
	// Test Tick function with a 1-second interval and 5 ticks
	t.Run("Tick", func(t *testing.T) {
		tick := Tick(1*time.Second, 5)
		tm := time.Now()
		for i, v := range tick {
			timeStr := v.Format(time.RFC3339)
			d := v.Sub(tm)
			if d < 900*time.Millisecond {
				t.Errorf("Tick %d should have been sent after at least 900ms, got %v", i+1, d)
			}
			tm = v
			fmt.Println("Tick:", i+1, ".", timeStr)
		}
	})

	t.Run("Fast forward", func(t *testing.T) {
		tick := Tick(1*time.Second, 5)
		tm := time.Now()
		for i, v := range tick {
			timeStr := v.Format(time.RFC3339)
			d := v.Sub(tm)
			if d < 900*time.Millisecond {
				t.Errorf("Tick %d should have been sent after at least 900ms, got %v", i+1, d)
			}
			tm = v
			fmt.Println("Tick:", i+1, ".", timeStr)

			// Fast forward to the next tick
			time.Sleep(time.Until(v.Add(1 * time.Second)))
		}
	})
}

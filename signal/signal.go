package signal

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
)

func Signal() {
	fmt.Println("Start")
	sig := make(chan os.Signal, 1)
	stopCount := 0
	signal.Notify(sig, os.Interrupt, os.Kill)
	for {
		if stopCount == 3 {
			break
		} else {
			rec := <-sig
			stopCount++
			fmt.Println("Received signal:", rec)
		}
	}
	fmt.Println("Exiting...in 3 seconds")
	time.Sleep(3 * time.Second)
	os.Exit(0)
}

func SignalCtx() {
	fmt.Println("Start")
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	stopCount := 0

	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, os.Kill)
	defer stop()

	for {
		if stopCount == 3 {
			stop()
			cancel()
			break
		} else {
			<-ctx.Done()
			stopCount++
			fmt.Println("Received signal")
		}
	}
	fmt.Println("Exiting...in 3 seconds")
	time.Sleep(3 * time.Second)
	os.Exit(0)
}

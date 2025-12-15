package context

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Run() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stream := make(chan string)

	var wg sync.WaitGroup
	wg.Add(3)
	go process1(ctx, &wg, stream)
	go process2(ctx, &wg, stream)
	go process3(ctx, &wg, stream)

	select {
	case msg, ok := <-stream:
		if ok {
			fmt.Println(msg)
		}
	case <-ctx.Done():
		fmt.Println("Main is done")
	}

	wg.Wait()
}

func process1(ctx context.Context, wg *sync.WaitGroup, stream chan<- string) {
	defer wg.Done()
	p1ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	var wgs sync.WaitGroup
	wgs.Add(3)

	doWork := func() {
		defer wg.Done()
		for {
			stream <- "work1"
		}
	}

	for range 3 {
		go doWork()
	}

	select {
	case <-p1ctx.Done():
		fmt.Println("Process1 is done")
	}

	wgs.Wait()
}

func work2(ctx context.Context, wg *sync.WaitGroup, stream chan<- string) {
	defer wg.Done()

	for {
		stream <- "work2"
	}

}

func subprocess2(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

}

func subprocess3(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

}

func process2(ctx context.Context, wg *sync.WaitGroup, stream chan<- string) {
	defer wg.Done()

	p3ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second*3))
	defer cancel()

	var wgs sync.WaitGroup
	go work2(ctx, &wgs, stream)

	select {
	case <-p3ctx.Done():
		fmt.Println("Process2 is done")
		return
	}
}

func process3(ctx context.Context, wg *sync.WaitGroup, stream chan<- string) {
	defer wg.Done()

}

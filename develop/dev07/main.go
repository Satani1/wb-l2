package main

import (
	"fmt"
	"time"
)

func or[T any](channels ...<-chan T) <-chan T {
	out := make(chan T)
	exitChan := make(chan struct{})

	for _, ch := range channels {
		go func(ch <-chan T) {
			for n := range ch {
				out <- n
			}
			exitChan <- struct{}{}
		}(ch)
	}

	go func() {
		<-exitChan
		close(out)
	}()

	return out
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(5*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
		sig(2*time.Second),
	)
	fmt.Printf("fone after %v\n", time.Since(start))
}

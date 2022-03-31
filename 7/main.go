package main

import (
	"fmt"
	"sync"
	"time"
)

func or(channels ...<-chan interface{}) <-chan interface{} {
	singleCh := struct {
		ch     chan interface{}
		closed bool
		sync.Mutex
	}{
		ch:     make(chan interface{}),
		closed: false,
	}

	wg := sync.WaitGroup{}

	for _, val := range channels {
		wg.Add(1)
		go func(done <-chan interface{}) {
			defer wg.Done()
			for {
				select {
				case _, ok := <-done:

					if ok {
						continue
					}
					singleCh.Lock()

					if singleCh.closed {
						singleCh.Unlock()
						return
					}

					close(singleCh.ch)
					singleCh.closed = true
					singleCh.Unlock()
				case <-singleCh.ch:
					return
				}
			}
		}(val)
	}

	wg.Wait()

	return singleCh.ch
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
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("Done %v", time.Since(start))
}

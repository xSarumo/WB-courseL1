package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	value int64
}

func main() {
	var cnt Counter
	iterations := 10000
	cntWorkers := 3

	var wg sync.WaitGroup

	for j := 0; j < cntWorkers; j++ {
		wg.Add(1)
		go func() {
			for i := 0; i < iterations; i++ {
				atomic.AddInt64(&cnt.value, 1)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(cnt.value)
}

package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

func worker(ch chan int, wg *sync.WaitGroup, id int) {
	defer wg.Done()
	for value := range ch {
		fmt.Printf("Worker %v: %v\n", id, value)
	}
}

func main() {
	const cntWorkers int = 4

	globalCh := make(chan int)
	var wg sync.WaitGroup

	for i := 0; i < cntWorkers; i++ {
		wg.Add(1)
		go worker(globalCh, &wg, i+1)
	}

	sysSignCh := make(chan os.Signal, 1)
	signal.Notify(sysSignCh, os.Interrupt)

	for {
		select {
		case <-sysSignCh:
			close(globalCh)
			wg.Wait()
			fmt.Println("All workers are closed")
			return
		default:
			value := rand.Intn(42)
			globalCh <- value
			time.Sleep(100 * time.Millisecond)
		}
	}

}

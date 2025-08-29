package main

import (
	"fmt"
	"math/rand"
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

	globalCh := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go worker(globalCh, &wg, 1)

	var inputTime int
	_, err := fmt.Scan(&inputTime)
	if err != nil {
		fmt.Println("Неверно введенные значения")
		return
	}

	timerCh := time.After(time.Duration(inputTime) * time.Second)

	for {
		select {
		case <-timerCh:
			close(globalCh)
			wg.Wait()
			fmt.Println("Program has stoped after:", inputTime, "seconds")
			return
		default:
			value := rand.Intn(42)
			globalCh <- value
			time.Sleep(100 * time.Millisecond)
		}
	}

}

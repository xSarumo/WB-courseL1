package main

import (
	"fmt"
	"sync"
)

func evaluate(values []int) {
	var wg sync.WaitGroup
	wg.Add(len(values))
	for _, i := range values {
		go square(i, &wg)
	}
	wg.Wait()
}

func square(n int, wg *sync.WaitGroup) {
	fmt.Println(n, "в квадрате:", n*n)
	wg.Done()
}

func main() {
	values := []int{2, 4, 6, 8, 10}
	evaluate(values)
}

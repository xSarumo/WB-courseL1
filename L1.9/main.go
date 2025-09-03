package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"os/signal"
	"sync"
	"time"
)

func GenerateRandX(ch chan int, sys chan os.Signal, wg *sync.WaitGroup) {
	for {
		select {
		case <-sys:
			close(ch)
			fmt.Println("Генерация случайных x завершена!")
			wg.Done()
			return
		default:
			time.Sleep(100 * time.Millisecond)
			ch <- rand.IntN(101)
		}
	}
}

func EvaluateNums(ch chan int, res chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range ch {
		time.Sleep(300 * time.Millisecond)
		res <- val * 2
	}
	fmt.Println("Рассчет всех х*2 был завершен!")
	close(res)

}

func main() {
	sysSignalCh := make(chan os.Signal, 1)
	signal.Notify(sysSignalCh, os.Interrupt)

	InputCh := make(chan int)
	ResultCh := make(chan int)

	var wg sync.WaitGroup

	wg.Add(2)
	go GenerateRandX(InputCh, sysSignalCh, &wg)
	go EvaluateNums(InputCh, ResultCh, &wg)

	for res := range ResultCh {
		time.Sleep(150 * time.Millisecond)
		fmt.Println(res)
	}
	wg.Wait()
	fmt.Println("Программа завершила свое выполнение!")
}

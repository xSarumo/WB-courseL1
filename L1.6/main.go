package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"time"
)

func FirstMethod() {
	//Остановка через канал уведомления, закрытие и остановка при нажатии Ctrl+C + использование WaitGroup
	signCh := make(chan os.Signal, 1)
	signal.Notify(signCh, os.Interrupt)

	var wg sync.WaitGroup
	tmpCh := make(chan int, 1)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range tmpCh {
			i++
		}
	}()

	for {
		select {
		case <-signCh:
			close(tmpCh)
			wg.Wait()
			fmt.Println("Горутина завершила свое выполнение корректно!")
			return
		default:
			fmt.Println("Горутина работает")
			time.Sleep(100 * time.Millisecond)
			continue
		}
	}
}

func SecondMethod() {
	// Использование своего канала для остановки
	stopCh := make(chan struct{})
	go func() {
		for {
			select {
			case <-stopCh:
				fmt.Println("Горутина остановила свое выполнение корректно!")
				return
			default:
				fmt.Println("Горутина работает")
				time.Sleep(1000 * time.Millisecond)
			}
		}
	}()

	time.Sleep(3 * 1000 * time.Millisecond)
	stopCh <- struct{}{}
	close(stopCh)
	os.Exit(0)
}

func ThirdMethod() {
	// Использование time.After + WaitGroup для остановки
	wait := time.Duration(3) * time.Second
	stopCh := time.After(wait)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for {
			select {
			case <-stopCh:
				wg.Done()
				fmt.Println("Горутина остановила свое выполнение корректно!")
				return
			default:
				fmt.Println("Горутина работает")
				time.Sleep(1000 * time.Millisecond)
			}
		}
	}()

	wg.Wait()
	os.Exit(0)
}

func FourthMethod() {
	// Использование context.WithCancel для остановки
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		tiker := time.NewTicker(1000 * time.Millisecond)
		defer tiker.Stop()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Горутина остановила свое выполнение корректно!")
				return
			case <-tiker.C:
				fmt.Println("Горутина долго работает")
			}
		}
	}()

	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(100 * time.Millisecond)
}

func FifthMethod() {
	// Остановка с помощью context.WithTimeout
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	go func() {
		fmt.Println("Горутина долго работает...")
		select {
		case <-ctx.Done():
			fmt.Println("Горутина остановила свое выполнение корректно!")
			return
		case <-time.After(5 * time.Second):
			fmt.Println("Горутина выполнила свою задачу через долгое время")
		}
	}()

	time.Sleep(4 * time.Second)
	cancel()
	time.Sleep(100 * time.Millisecond)
	os.Exit(0)
}

func SixthMethod() {
	//Остановка горутины через runtime.Goexit()
	go func() {
		defer fmt.Println("defer выполнен.\nГорутина остановленна корректно!") // Все defer при использовании такого похода все равно будут выполнены
		fmt.Println("Горутина работает...")
		time.Sleep(2 * time.Second)
		runtime.Goexit()
		fmt.Println("Горутина завершила свое выплнение") // А вот этот код в отличие от defer не будет выполнен
	}()

	time.Sleep(2500 * time.Millisecond)
	os.Exit(0)
}

func main() {
	var methodNum int
	fmt.Println("Введите номер метода (1-6): ")
	fmt.Scan(&methodNum)

	switch methodNum {
	case 1:
		FirstMethod()
	case 2:
		SecondMethod()
	case 3:
		ThirdMethod()
	case 4:
		FourthMethod()
	case 5:
		FifthMethod()
	case 6:
		SixthMethod()
	default:
		fmt.Println("Нет такого варианта выбора")
	}

}

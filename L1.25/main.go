package main

import (
	"fmt"
	"time"
)

func sleep(duration time.Duration) {
	target := time.Now().Add(duration)
	for time.Now().Before(target) {
	}
}

func sleepChan(duration time.Duration) {
	<-time.After(duration)
}

func main() {
	for i := 0; i < 100; i++ {
		if i == 45 {
			sleep(3 * time.Second)
		}
		fmt.Print(i, " ")
	}

	fmt.Println()
	fmt.Println()
	fmt.Println()

	for i := 0; i < 100; i++ {
		if i == 45 {
			sleepChan(3 * time.Second)
		}
		fmt.Print(i, " ")
	}
}

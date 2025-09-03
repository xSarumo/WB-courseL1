package main

import "fmt"

func main() {
	temperature := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	result := make(map[int][]float32)
	for _, i := range temperature {
		key := (int(i) / 10) * 10
		result[key] = append(result[key], i)
	}

	for group, temps := range result {
		fmt.Printf("%d:%v ", group, temps)
	}
}

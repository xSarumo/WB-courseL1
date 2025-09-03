package main

import "fmt"

func MemoryOptimization() {
	firstUnion := []int{1, 2, 3}
	secondUnion := []int{2, 3, 4}

	resultUnion := make([]int, 0, max(len(firstUnion), len(secondUnion))/2)
	for _, i := range firstUnion {
		for _, j := range secondUnion {
			if (i ^ j) == 0 {
				resultUnion = append(resultUnion, i)
			}
		}
	}

	fmt.Println(resultUnion)
} // O(n^2)

func TimeOptimization() {
	firstUnion := map[int]bool{1: false, 2: false, 3: false}
	secondUnion := []int{2, 3, 4}

	for _, i := range secondUnion {
		if _, hit := firstUnion[i]; hit {
			firstUnion[i] = true
		}
	}

	fmt.Print("{ ")
	for val, hit := range firstUnion {
		if hit {
			fmt.Print(val, " ")
		}
	}
	fmt.Print("}")
} // O(n)

func main() {
	MemoryOptimization()
	TimeOptimization()
}

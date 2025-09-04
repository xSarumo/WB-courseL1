package main

import "fmt"

func PrintFunc(fn func()) {
	fmt.Print("{ ")
	fn()
	fmt.Print("}")
}

func main() {
	input := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]bool)

	for _, key := range input {
		set[key] = true
	}

	iteration := func() {
		for key, value := range set {
			if value {
				fmt.Print(key, " ")
			}
		}
	}

	PrintFunc(iteration)
}

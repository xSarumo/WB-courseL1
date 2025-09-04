package main

import "fmt"

func binSearch(arr []int, hit int) int {
	left := 0
	right := len(arr)

	for left+1 < right {
		middle := left + ((right - left) / 2)
		if hit >= arr[middle] {
			left = middle
		} else {
			right = middle
		}
	}
	if left != hit {
		return -1
	}

	return left
}

func main() {
	arr := []int{1, 2, 4, 5, 14, 15, 89, 100, 105, 115}
	fmt.Println("Выберете число из массива:", arr)
	var hit int
	_, err := fmt.Scan(&hit)

	if err != nil {
		fmt.Println("Ошибка ввода", err)
	}

	result := binSearch(arr, hit)
	if result != -1 {
		fmt.Println("Index of value:", result, "| Check out(arr[result]):", arr[result])
	} else {
		fmt.Println("Число не найдено")
	}
}

package main

import "fmt"

func deleteElem[T any](arr []T, indx int) []T {
	length := len(arr)
	if indx < 0 || indx >= length {
		return arr
	}

	copy(arr[indx:], arr[indx+1:])

	var zero T
	arr[length-1] = zero

	return arr[:length-1]
}

func main() {
	input := []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	fmt.Println("Массив:", input)
	fmt.Print("Введите индекс удаляемого элемента:")

	var index int
	_, err := fmt.Scan(&index)
	deleted := input[index]
	if err != nil {
		fmt.Println("Ошибка ввода")
	}

	fmt.Println("Было удалено:", deleted, "|", deleteElem(input, index))
}

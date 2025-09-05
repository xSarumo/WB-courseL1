package main

import "fmt"

func CheckUnique(str string) bool {
	runes := []rune(str)
	dict := make(map[rune]bool)

	for _, r := range runes {
		if dict[r] {
			return false
		}
		dict[r] = true
	}

	return true
}

func main() {
	fmt.Println("Введите строку:")
	var input string
	_, err := fmt.Scan(&input)

	if err != nil {
		fmt.Println("Ошибка ввода")
		return
	}

	fmt.Println(CheckUnique(input))
}

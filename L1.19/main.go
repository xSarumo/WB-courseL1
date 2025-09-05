package main

import "fmt"

func main() {
	var input string
	_, err := fmt.Scan(&input)

	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	length := len(input)
	result := make([]rune, length)

	for i, r := range []rune(input) {
		result[(length-1)-i] = r
	}

	fmt.Println(string(result))
}

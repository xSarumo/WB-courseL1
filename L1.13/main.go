package main

import "fmt"

func main() {
	firstNum := 5
	secondNum := 3
	fmt.Println(firstNum, secondNum)

	firstNum ^= secondNum
	secondNum ^= firstNum
	firstNum ^= secondNum

	fmt.Println(firstNum, secondNum)
}

package main

import (
	"fmt"
	"math/big"
)

func Add(a, b *big.Int) *big.Int {
	return new(big.Int).Add(a, b)
}

func Subtract(a, b *big.Int) *big.Int {
	return new(big.Int).Sub(a, b)
}

func Multiply(a, b *big.Int) *big.Int {
	return new(big.Int).Mul(a, b)
}

func Divide(a, b *big.Int) *big.Int {
	if b.Cmp(big.NewInt(0)) == 0 {
		return nil
	}
	return new(big.Int).Div(a, b)
}

func main() {
	var aStr, bStr string

	fmt.Println("Введите первое большое число:")
	fmt.Scan(&aStr)

	fmt.Println("Введите второе большое число:")
	fmt.Scan(&bStr)

	a, okA := new(big.Int).SetString(aStr, 10)
	b, okB := new(big.Int).SetString(bStr, 10)

	if !okA || !okB {
		fmt.Println("Ошибка: не удалось прочитать числа. Убедитесь, что вы ввели корректные числовые значения.")
		return
	}

	fmt.Println("\nВыберите операцию:")
	fmt.Println("1 - Сложение (+)")
	fmt.Println("2 - Вычитание (-)")
	fmt.Println("3 - Умножение (*)")
	fmt.Println("4 - Деление (/)")

	var operation int
	_, err := fmt.Scan(&operation)
	if err != nil {
		fmt.Println("Ошибка ввода операции:", err)
		return
	}

	var result *big.Int
	var operationSymbol string

	switch operation {
	case 1:
		result = Add(a, b)
		operationSymbol = "+"
	case 2:
		result = Subtract(a, b)
		operationSymbol = "-"
	case 3:
		result = Multiply(a, b)
		operationSymbol = "*"
	case 4:
		result = Divide(a, b)
		operationSymbol = "/"
		if result == nil {
			fmt.Println("Ошибка: деление на ноль!")
			return
		}
	default:
		fmt.Println("Неизвестная операция.")
		return
	}

	fmt.Printf("\nРезультат: %s %s %s = %s\n", a.String(), operationSymbol, b.String(), result.String())
}

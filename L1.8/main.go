package main

import "fmt"

func InvertBit(value int64, position int8) int64 {
	return value ^ (int64(1) << position)
}

func SetOneBit(value int64, position int8) int64 {
	return value | (int64(1) << position)
}

func SetZeroBit(value int64, position int8) int64 {
	return value &^ (int64(1) << position)
}

func main() {
	var value int64
	fmt.Print("Введите число для int64: ")
	_, err1 := fmt.Scan(&value)
	if err1 != nil {
		fmt.Println("Ошибка ввода")
		return
	}

	var position int8
	fmt.Print("Введите номер бита(0-63): ")
	_, err2 := fmt.Scan(&position)
	if err2 != nil {
		fmt.Println("Ошибка ввода")
		return
	}

	if position > 63 || position < 0 {
		return
	}

	var operation int
	fmt.Print("Выберите тип операции:")
	_, err3 := fmt.Scan(&operation)
	if err3 != nil {
		return
	}

	lastValue := value
	switch operation {
	case 1:
		value = InvertBit(value, position)
	case 2:
		value = SetOneBit(value, position)
	case 3:
		value = SetZeroBit(value, position)
	}

	fmt.Printf("Введенное число: %d   {%064b}\n", lastValue, int64(lastValue))
	fmt.Printf("Введенная позиция: %d {%064b}\n", position, int64(int64(1)<<position))
	fmt.Printf("Результат: %d         {%064b}", value, int64(value))
}

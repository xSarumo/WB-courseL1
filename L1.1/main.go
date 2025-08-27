package main

import (
	"fmt"
)

type Human struct {
	Name   string
	Height float32
	Age    int
	Weight float32
}

func (targ *Human) Introduce() {
	fmt.Printf("Имя: %s,\nВозраст: %v,\nВес: %v,\nРост: %v\n", targ.Name, targ.Age, targ.Weight, targ.Height)
}

type Action struct {
	Human  // Встраивание структуры (Аналог наследования)
	action string
}

func (targ *Human) Dance() {
	fmt.Printf("%s сейчас танцует\n", targ.Name)
}

func main() {
	action := Action{
		Human: Human{
			Name:   "Petr",
			Age:    18,
			Weight: 65.6,
			Height: 150.8,
		},
		action: "run",
	}

	action.Introduce()
	fmt.Println()
	action.Dance()
	fmt.Println()
	fmt.Println("Тип действия: ", action.action)
	fmt.Println("Имя исполнителя: ", action.Name)
}

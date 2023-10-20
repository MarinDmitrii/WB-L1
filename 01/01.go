package main

// Дана структура Human (с произвольным набором полей и методов). Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

import "fmt"

type Human struct {
	name string
}

// Объявляем функцию для типа Human
func (h *Human) Info() {
	fmt.Println("My name is", h.name, "and I'm happy!")
}

// Встраивание структуры Human в структуру Action
type Action struct {
	Human
}

func (a Action) Run() {
	fmt.Println(a.name, "is running!")
}

func main() {
	// Объявляем переменную типа Action{}
	Ashely := Action{Human{"Ashley"}}
	// вызов функции Info, объявленной для типа Human{}, через переменную типа Action{}
	Ashely.Info()
	Ashely.Run()
}

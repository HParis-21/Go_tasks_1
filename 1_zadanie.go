/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры
Human (аналог наследования).
*/

package main

import "fmt"

type Human struct {
	name string
	age  int
}

// Action is a struct which embedds Human
type Action struct {
	Human
}

func (h Human) PrintHuman() {
	fmt.Printf("My name is %s. I am %d years old\n", h.name, h.age)
}

func (h *Human) HappyBirthday() {
	fmt.Printf("Happy birthday %s! You are now %d!", h.name, h.age+1)
}

func main() {
	bob := Human{"Bob", 33}
	bob.PrintHuman()
	bob.HappyBirthday()

	tom := Action{Human{"Tom", 12}}
	tom.PrintHuman()
	tom.HappyBirthday()

}

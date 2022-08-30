/*
Реализовать паттерн «адаптер» на любом примере.
*/
package main

import "fmt"

type pilot interface {
	flay()
}

type duck struct {
	name string
}

func (d duck) flay() {
	fmt.Printf("My name is %s. I'm flying south!\n", d.name)
}

type human struct {
	name string
}

func (h human) airplane() {
	fmt.Printf("My name is %s. I'm flying south!\n", h.name)
}

//adapter
type humanAdapter struct {
	human
}

func (ha humanAdapter) flay() {
	ha.airplane()
}

func main() {
	du := duck{"Сrack"}
	bob := human{"Bob"}
	hAdapter := humanAdapter{human: bob}
	du.flay()
	hAdapter.flay()

}

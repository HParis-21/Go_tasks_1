/*
Разработать программу, которая в рантайме способна определить тип
переменной: int, string, bool, channel из переменной типа interface{}.
*/
package main

import "fmt"

func main() {
	printType("s")
	printType(1)
	printType(0.5)
	printType(true)
	printType(make(chan int))

	printType2("s")
	printType2(1)
	printType2(0.5)
	printType2(true)
	printType2(make(chan int))
}

func printType(i interface{}) {
	switch v := i.(type) {
	case string:
		fmt.Println("string", v)
	case int, int8, int16, int32, int64:
		fmt.Println("integer", v)
	case float32, float64:
		fmt.Println("float", v)
	case bool:
		fmt.Println("bool", v)
	case chan int:
		fmt.Println("chan int", v)
	}
}

func printType2(i interface{}) {
	if val, ok := i.(string); ok {
		fmt.Println("string", val)
		return
	}
	if val, ok := i.(int); ok {
		fmt.Println("int", val)
	}
	if val, ok := i.(float32); ok {
		fmt.Println("float", val)
	}
	if val, ok := i.(bool); ok {
		fmt.Println("bool", val)
	}
	if val, ok := i.(chan int); ok {
		fmt.Println("chan int", val)
	}
}

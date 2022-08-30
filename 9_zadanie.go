/*
Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из
массива, во второй — результат операции x*2, после чего данные из второго
канала должны выводиться в stdout.
*/
package main

import (
	"fmt"
)

func main() {
	nums := make([]int, 20)
	for i := range nums {
		nums[i] = i
	}
	chInput := make(chan int)
	chOutput := make(chan int)

	go writeSource(nums, chInput)
	go multi2(chInput, chOutput)
	for elem := range chOutput {
		fmt.Println(elem)
	}
}

func writeSource(data []int, ch chan<- int) {
	defer close(ch)
	for _, elem := range data {
		ch <- elem
	}
}

func multi2(chIn <-chan int, chOut chan<- int) {
	defer close(chOut)
	for elem := range chIn {
		chOut <- elem * 2
	}
}

/*
Реализовать быструю сортировку массива (quicksort) встроенными методами
языка.
*/
package main

import (
	"fmt"
	"math/rand"
)

func Run() {
	source := []int{6, 1, 23, 33, -2, 5, 34, -15, 6, 2}
	fmt.Println(quickSort(source))
}
func quickSort(source []int) []int {
	if len(source) < 2 {
		return source
	}
	// запоминаем индексы крайних элементов
	left := 0
	right := len(source) - 1
	// выбираем случайный индекс, относительно которого будем вращать (опорный элемент)
	pivotIndex := rand.Uint64() % uint64(len(source))
	// меняем элемент с нашим случайным индексом и крайний правый местами
	source[pivotIndex], source[right] = source[right], source[pivotIndex]
	// отбрасываем элементы, меньшие опорного, левее
	for i := range source {
		if source[i] < source[right] {
			source[i], source[left] = source[left], source[i]
			left++
		}
	}
	// опорник ставим сразу после последнего элемента, который оказался меньше опорного
	source[left], source[right] = source[right], source[left]
	quickSort(source[:left])
	quickSort(source[left+1:])
	return source
}

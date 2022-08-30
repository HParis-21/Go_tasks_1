/*
Написать программу, которая конкурентно рассчитает значение квадратов чисел
взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	array := []int{2, 4, 6, 8, 10}
	//Squares(array)
	//Squares2(array)
	SquaresWG(array)
}

//Squares prints squares of every number of the given nums slice
//using an unbuffered channel
func Squares(nums []int) {
	squares := make(chan int)
	for _, val := range nums {
		go func(val int) {
			squares <- val * val
		}(val)
	}
	for i := 0; i < len(nums); i++ {
		fmt.Println(<-squares)
	}
}

//Squares2 prints squares of every number of the given nums slice
//using a buffered channel
func Squares2(nums []int) {
	squares := make(chan int, 5)
	for _, val := range nums {
		go func(val int) {
			squares <- val * val
		}(val)
	}
	for i := 0; i < len(nums); i++ {
		fmt.Println(<-squares)
	}

}

// SquaresWG prints squares of every number of the given nums slice
// using sync.WaitGroup
func SquaresWG(nums []int) {
	var wg sync.WaitGroup
	wg.Add(5)
	for _, val := range nums {
		go func(val int) {
			defer wg.Done()
			fmt.Println(val * val)
		}(val)
	}
	wg.Wait()
}

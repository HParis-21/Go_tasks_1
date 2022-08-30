/*
Дана последовательность чисел: 2,4,6,8,10. Найти сумму их
квадратов(2*2+3*3+4*4….) с использованием конкурентных вычислений.
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	array := []int{2, 4, 6, 8, 10}
	squares_sum(array)

}

// SquaresSumWG calculates and prints sum of nums squared
// using sync.WaitGroup and sync.Mutex
func squares_sum(nums []int) {
	sum := struct {
		value int
		sync.Mutex
	}{}

	var wg sync.WaitGroup
	wg.Add(5)
	for _, val := range nums {
		go func() {
			defer wg.Done()
			sum.Lock()
			defer sum.Unlock()
			sum.value += val * val
		}()
	}
	wg.Wait()
	fmt.Println("Sum:", sum.value)
}

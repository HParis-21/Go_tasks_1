/*
Реализовать структуру-счетчик, которая будет инкрементироваться в
конкурентной среде. По завершению программа должна выводить итоговоезначение счетчика.
*/
package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	val int
	sync.RWMutex
}

func (c *Counter) increment() {
	c.Lock()
	defer c.Unlock()
	c.val++
}

func (c *Counter) value() int {
	c.RLock()
	defer c.RUnlock()
	return c.val
}

func main() {
	count := Counter{}
	var wg sync.WaitGroup
	wg.Add(6)
	for i := 0; i < 6; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 3; i++ {
				count.increment()
			}
		}()
	}
	wg.Wait()
	fmt.Println("Result:", count.val)
}

/*
Реализовать конкурентную запись данных в map.
*/
package main

import (
	"fmt"
	"sync"
)

type safeMap struct {
	mx sync.Mutex
	m  map[int]int
}

func (s *safeMap) addValue(key int, value int) {
	s.mx.Lock()
	defer s.mx.Unlock()
	s.m[key] = value
}

func main() {
	myMap := &safeMap{m: make(map[int]int)}
	wg := &sync.WaitGroup{}
	for i := 1; i < 21; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			myMap.addValue(i, i*i)
		}(i)
	}
	wg.Wait()
	for key, value := range myMap.m {
		fmt.Printf("%d:%d\n", key, value)
	}
}

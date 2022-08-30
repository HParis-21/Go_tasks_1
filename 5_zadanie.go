/*
Разработать программу, которая будет последовательно отправлять значения в
канал, а с другой стороны канала — читать. По истечению N секунд программа
должна завершаться.
*/
package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Time is not specified.")
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n < 1 {
		fmt.Println("Time is not specified.")
		return
	}

	ch := make(chan int)
	timer := time.NewTimer(time.Duration(n) * time.Second)
	go func() {
		for i := range ch {
			fmt.Println(i)
		}
	}()

	i := 0
	for {
		select {
		case ch <- i:
			time.Sleep(1000 * time.Millisecond)
			i++
		case <-timer.C:
			close(ch)
			return
		}
	}
}

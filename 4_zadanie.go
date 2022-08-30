/*
Реализовать постоянную запись данных в канал (главный поток). Реализовать
набор из N воркеров, которые читают произвольные данные из канала и
выводят в stdout. Необходима возможность выбора количества воркеров при старте.
Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать
способ завершения работы всех воркеров.
*/
package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Specify the number of workers > 0")
		return
	}
	numsWork, err := strconv.Atoi(os.Args[1])
	if err != nil || numsWork < 1 {
		fmt.Println("Specify the number of workers > 0")
		return
	}

	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)

	chanIn := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(numsWork)
	for i := 0; i < numsWork; i++ {
		go worker(chanIn, i+1, wg)
	}

	count := 0
Loop:
	for {
		select {
		case <-stop:
			close(chanIn)
			break Loop
		case chanIn <- count:
			count++
		}
	}

	wg.Wait()
	fmt.Println("Received SIGINT: closing input. Last data:", count-1)
}

func worker(input <-chan int, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range input {
		fmt.Printf("worker: %d, data: %d\n", id, data)
		time.Sleep(1 * time.Second)
	}
	fmt.Printf("worker %d finished\n", id)
}

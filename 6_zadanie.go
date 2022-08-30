/*
Реализовать все возможные способы остановки выполнения горутины.
*/
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	//closeChan()
	//doneChan()
	doneContext()
}

func closeChan() {
	var wg sync.WaitGroup
	wg.Add(1)
	ch := make(chan string)
	go func() {
		for {
			foo, ok := <-ch
			if !ok {
				fmt.Println("closeChan Done")
				wg.Done()
				return
			}
			fmt.Println(foo)
		}
	}()
	ch <- "closeChan running"
	ch <- "closeChan running"
	ch <- "closeChan running"
	close(ch)

	wg.Wait()
}

func doneChan() {
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-quit:
				fmt.Println("doneChan Done")
				//close(quit)
				return
			default:
				fmt.Println("doneChan running")
				time.Sleep(1 * time.Second)
			}
		}
	}()
	time.Sleep(2 * time.Second)
	quit <- struct{}{}
}

func doneContext() {
	ch := make(chan struct{})
	finished := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case <-ctx.Done():
				println("doneContext done")
				finished <- struct{}{}
				return
			case <-ch:
				println("stopContext running")
			}
		}
	}()

	for i := 0; i < 5; i++ {
		ch <- struct{}{}
	}
	cancel()   //signal goroutine to finish
	<-finished //wait for goroutine to finish
}

//func main() {
//	forever := make(chan struct{})
//	ctx, cancel := context.WithCancel(context.Background())
//
//	go func(ctx context.Context) {
//		for {
//			select {
//			case <-ctx.Done():  // if cancel() execute
//				forever <- struct{}{}
//				return
//			default:
//				fmt.Println("for loop")
//			}
//
//			time.Sleep(500 * time.Millisecond)
//		}
//	}(ctx)
//
//	go func() {
//		time.Sleep(3 * time.Second)
//		cancel()
//	}()
//
//	<-forever
//	fmt.Println("finish")
//}

/*
Реализовать собственную функцию sleep
*/
package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Starting timer for 3 seconds, time:", time.Now().String())
	mSleep(3 * time.Second)
	// mSleep2(3 * time.Second)
	println("The timer is stopped. Current time:", time.Now().String())
}

func mSleep(d time.Duration) {
	ticker := time.NewTicker(d)
	<-ticker.C
}

func mSleep2(d time.Duration) {
	<-time.After(d)
}

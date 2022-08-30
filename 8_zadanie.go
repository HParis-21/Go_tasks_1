/*
Дана переменная int64. Разработать программу которая устанавливает i-й бит в
1 или 0.
*/
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) != 4 {
		fmt.Println("Invalid number of arguments.")
		fmt.Println("Specify: number position value")
		return
	}
	nums, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		fmt.Println("Invalid number")
		return
	}
	i, err := strconv.ParseInt(os.Args[2], 10, 8)
	if err != nil || i < 0 || i > 63 {
		fmt.Println("Invalid position(0-63)")
		return
	}
	val, err := strconv.ParseInt(os.Args[3], 2, 2)
	if err != nil {
		fmt.Println("Invalid value(0-1)")
		return
	}
	fmt.Printf("Before:\t%d - %s\n", nums, strconv.FormatInt(nums, 2))
	if val == 1 {
		nums = nums | (1 << i) //поразрядная дизъюнкция(или)
	} else {
		nums = nums &^ (1 << i) //сброс бита(и не)
	}
	fmt.Printf("After:\t%d - %s\n", nums, strconv.FormatInt(nums, 2))
}

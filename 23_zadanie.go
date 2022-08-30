/*
Удалить i-ый элемент из слайса.
*/
package main

import "fmt"

func main() {
	nums := make([]int, 10)
	for i := range nums {
		nums[i] = i
	}
	fmt.Println("Before:", nums)
	// nums = removeFastUnstable(nums, 4)
	nums = removeSlowStable(nums, 4)
	fmt.Println("After:", nums)

}

func removeFastUnstable(nums []int, i int) []int {
	nums[i] = nums[len(nums)-1]
	return nums[:len(nums)-2]
}

func removeSlowStable(nums []int, i int) []int {
	return append(nums[:i], nums[i+1:]...)
}

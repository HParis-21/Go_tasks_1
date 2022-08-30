/*
Реализовать пересечение двух неупорядоченных множеств.
*/
package main

import "fmt"

func main() {
	nums1 := []int{8, 12, 5, 7, 3, 6, 7, 15, 9}
	nums2 := []int{16, 8, 3, 4, 2, 12, 0, 4, 66}

	fmt.Println(intersection(nums1, nums2))
	fmt.Println(intersection2(nums1, nums2))
}

func intersection(nums1, nums2 []int) []int {
	allMap := make(map[int]int)
	for _, elem1 := range nums1 {
		allMap[elem1] = 1
	}
	for _, elem2 := range nums2 {
		if _, ok := allMap[elem2]; ok {
			allMap[elem2] = 2
		}
	}
	result := make([]int, 0)
	for key, value := range allMap {
		if value == 2 {
			result = append(result, key)
		}
	}
	return result

}

func intersection2(nums1, nums2 []int) []int {
	allMap := make(map[int]struct{})
	result := make([]int, 0)
	for _, elem := range nums1 {
		allMap[elem] = struct{}{}
	}
	for _, elem := range nums2 {
		if _, ok := allMap[elem]; ok {
			result = append(result, elem)
		}
	}
	return result
}

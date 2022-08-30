/*
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее
собственное множество.
*/
package main

import "fmt"

func main() {
	src := []string{"cat", "cat", "dog", "cat", "tree"}
	checkMap := make(map[string]struct{})
	for _, str := range src {
		checkMap[str] = struct{}{}
	}

	res := make([]string, 0)
	for key := range checkMap {
		res = append(res, key)
	}
	fmt.Println(res)
}

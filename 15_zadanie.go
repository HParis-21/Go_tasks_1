/*
var justString string
func someFunc() {
v := createHugeString(1 << 10)
justString = v[:100]
}
func main() {
someFunc()
}
*/
package main

import (
	"fmt"
)

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
	fmt.Println(justString)
}

func createHugeString(size int) string {
	res := make([]byte, size)
	for n := range res {
		res[n] = '7'
	}
	return string(res)
}
func main() {
	someFunc()
}

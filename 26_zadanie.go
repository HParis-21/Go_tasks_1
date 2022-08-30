/*
Разработать программу, которая проверяет, что все символы в строке
уникальные (true — если уникальные, false etc). Функция проверки должна быть
регистронезависимой.

Например:
abcd — true
abCdefAaf — false
aabcd — false
*/
package main

import (
	"fmt"
	"os"
	"unicode"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Enter text.")
		return
	}
	fmt.Println(testUnique(os.Args[1]))
}

func testUnique(s string) bool {
	runes := make(map[rune]struct{})
	for _, r := range s {
		r = unicode.ToLower(r)
		if _, exists := runes[r]; exists {
			return false
		}
		runes[r] = struct{}{}
	}
	return true
}

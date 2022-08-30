/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow»
*/
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("string to be reversed")
		return
	}
	fmt.Println(reversedWords(os.Args[1]))
}

func reversedWords(s string) string {
	words := strings.Split(s, " ")
	for i := 0; i < len(words)/2; i++ {
		words[i], words[len(words)-i-1] = words[len(words)-i-1], words[i]
	}
	return strings.Join(words, " ")
}

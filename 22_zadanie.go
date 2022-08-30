/*
Разработать программу, которая перемножает, делит, складывает, вычитает две
числовых переменных a,b, значение которых > 2^20.
*/
package main

import (
	"fmt"
	"math/big"
	"os"
	"strings"
)

const (
	base = 10
	prec = 1024
)

func main() {

	if len(os.Args) != 4 {
		fmt.Println("Invalid number of arguments.")
		return
	}
	a, _, err := big.ParseFloat(os.Args[1], base, prec, big.ToNearestEven)
	if err != nil {
		fmt.Println("Invalid number.")
		return
	}
	b, _, err := big.ParseFloat(os.Args[3], base, prec, big.ToNearestEven)
	if err != nil {
		fmt.Println("Invalid number.")
		return
	}
	sign := os.Args[2]
	if !strings.Contains("+-*/", sign) || len(sign) != 1 {
		fmt.Println("Invalid sign.")
		return
	}

	var result big.Float
	switch sign {
	case "+":
		result.Add(a, b)
	case "-":
		result.Sub(a, b)
	case "*":
		result.Mul(a, b)
	case "/":
		result.Quo(a, b)
	}

	fmt.Println("result = ", result.Text('f', 10))
}

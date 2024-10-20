package code

import (
	"fmt"
)

func Add(a, b float64) float64 {

	return a + b
}

func Multiply(a, b float64) float64 {

	return a * b
}

func PrintNumbersAscending(n int) {
	for i := 1; i <= n; i++ {
		fmt.Printf("%d  ", i)
	}

}

//Функции, которые нужно реализовать в функциональном операторе выше

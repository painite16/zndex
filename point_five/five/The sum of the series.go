package code

import (
	_ "fmt"
)

func CalculateSeriesSum(n int) float64 {
	sum := 0.0

	for i := 1; i <= n; i++ {
		sum += 1 / float64(i)
	}
	return sum
}

//Ваша задача — написать функцию CalculateSeriesSum(n int) float64, которая получает на вход число n и вычисляет сумму указанной последовательности. Затем программа должна вернуть полученное значение.

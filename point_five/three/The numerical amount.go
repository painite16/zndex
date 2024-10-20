package code

import (
	_ "fmt"
)

func SumDigitsRecursive(n int) int {

	if n < 10 {

		return n
	} else {
		return n%10 + SumDigitsRecursive(n/10)
	}
}

//Напишите функцию SumDigitsRecursive(n int) int, которая рекурсивно вычисляет сумму цифр числа. Как и в предыдущем задании, не забудьте объявить пакет main; при этом определять функцию main() и вызывать SumDigitsRecursive() не нужно – это сделает проверяющая система.

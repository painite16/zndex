package code

import (
	"fmt"
	"math"
)

func SqRoots() {
	var a, b, c float64
	fmt.Scanln(&a, &b, &c)

	discriminant := b*b - 4*a*c

	if discriminant > 0 {
		d := (-b + math.Sqrt(discriminant)) / (2 * a)
		e := (-b - math.Sqrt(discriminant)) / (2 * a)
		fmt.Println(e, d)
	} else if discriminant == 0 {
		d := (-b) / (2 * a)
		fmt.Println(d)
	} else {
		fmt.Println("0 0")
	}
}

//Для этого напишите функцию SqRoots(), в которой с клавиатуры через пробел вводятся три вещественных числа, при этом первое число гарантированно не равно нулю. Функция должна вывести на экран через пробел по возрастанию корни уравнения (один или два), или два числа 0, если корней нет.

//Функция ничего не должна возвращать.

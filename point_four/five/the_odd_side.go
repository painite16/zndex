package code

import "fmt"

func the_odd_side() {
	sum := 0
	var a int
	fmt.Scanln(&a)
	if a < 0 {
		fmt.Println("Некорректный ввод")
		return
	}
	for i := 1; i <= a; i++ {

		if i%2 == 1 {
			sum += i
		}

	}
	fmt.Println(sum)
}

//Дома на улице нумеруются по сторонам: с одной стороны — чётные, с другой — нечётные. Гоша идёт с начала улицы по нечётной стороне и считает сумму номеров домов. А мы можем никуда не ходить и написать программу, которая эту сумму вычислит.

//Напишите программу, которая запрашивает у пользователя число и выводит на экран сумму всех нечётных чисел от 1 до этого числа с помощью цикла for. Если число отрицательное, программа должна выводить сообщение Некорректный ввод.

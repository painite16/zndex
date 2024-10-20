package code

import "fmt"

func the_dividing_zero() {
	var number int

	fmt.Scanln(&number)

	if number > 0 {
		fmt.Println("Число положительное")
	} else if number < 0 {
		fmt.Println("Число отрицательное")
	} else {
		fmt.Println("Введен ноль")
	}

}

//Напишите программу, которая запрашивает у пользователя число и выводит на экран сообщение Число положительное, Введен ноль и Число отрицательное по результатам оценки числа.

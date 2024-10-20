package code

import "fmt"

func water_in_the_sieve() {
	sum := 0
	var a int
	fmt.Scanln(&a)
	for i := 0; i <= a; i++ {

		sum += i
		if i >= a {
			fmt.Println(sum)

		}
	}
}

//Напишите программу, которая запрашивает у пользователя число и выводит на экран сумму всех чисел от 1 до этого числа с помощью цикла for

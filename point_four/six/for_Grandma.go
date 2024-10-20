package code

import (
	"fmt"
)

func for_grandma() {
	var n int
	fmt.Scanln(&n)
	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}

	}

}

//Напишите программу, которая запрашивает у пользователя число и выводит на экран все числа от 1 до этого числа, которые делятся на 3 с помощью цикла for.

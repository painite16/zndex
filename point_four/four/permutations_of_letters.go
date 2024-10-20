package code

import "fmt"

func permutations_of_letters() {
	sum := 1
	var a int
	fmt.Scanln(&a)
	for i := 1; i <= a; i++ {

		sum *= i
		if i >= a {
			fmt.Println(sum)

		}
	}
}

//Напишите программу, которая запрашивает у пользователя число — количество букв в алфавите, а потом выводит на экран факториал этого числа (количество перестановок букв) с помощью цикла for.

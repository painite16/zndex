package main

import "fmt"

func the_multiplication_table() {
	var a int
	fmt.Scanln(&a)
	for i := 1; i <= 10; i++ {
		fmt.Println(a, "*", i, "=", a*i)
	}
}

//Таблицу умножения Гоша, конечно, знает, но иногда путает в ней цифры. Он и сам мог бы написать себе программку, которая выводит таблицу умножения для забытого числа, но занят проектами помасштабнее. Помогите Гоше: напишите программу, которая запрашивает у пользователя число и выводит на экран таблицу умножения от 1 до 10 для этого числа с помощью цикла for.

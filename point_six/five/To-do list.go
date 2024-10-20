package code

import (
	"fmt"
	"strconv"
)

func PrettyArrayOutput(array [9]string) {

	var sum int
	for i := 0; i < len(array); i++ {
		sum = i + 1
		var strok string = strconv.Itoa(sum)

		if i < 7 {
			fmt.Println(strok + " я уже сделал: " + array[i])
		} else {
			fmt.Println(strok + " не успел сделать: " + array[i])
		}
	}
}

//Напишите функцию PrettyArrayOutput(array [9]string), которая получает массив из девяти строк и печатает: "1 я уже сделал: *элемент массива*" для первых 7 элементов и "9 не успел сделать: *элемент массива*" для последних двух. У строчки должен быть именно такой вид (без кавычек, а вместо "*элемент массива*" сам элемент) и нумерация с единицы.

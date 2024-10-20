package code

import "strconv"

func ConcatStringsAndInt(str1, str2 string, num int) string {
	var a = strconv.Itoa(num)
	var general = str1 + str2 + a
	return general
}

//Напишите функцию ConcatStringsAndInt(str1, str2 string, num int) string, которая принимает две строки и одно целое число, а затем выполняет конкатенацию строк и числа в одну строку.

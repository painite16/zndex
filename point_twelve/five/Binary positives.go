package code

import (
	"errors"
	"strconv"
)

func IntToBinary(num int) (string, error) {
	if num < 0 {
		return "", errors.New("negative numbers are not allowed")
	}
	x := int64(num)
	bin := strconv.FormatInt(x, 2)
	txt := string(bin)
	return txt, nil
}

//Гоша хочет научиться переводить в двоичную систему положительные целые числа. Напишите функцию IntToBinary(num int) (string, error), которая принимает на вход целое число и возвращает его двоичное представление. Если пользователь ввёл отрицательное число, программа должна возвращать сообщение об ошибке (negative numbers are not allowed).

//Примечания
//Обратите внимание, что функция возвращает строку, а не число. В случае ошибки верните пустую строку.

//Много стандартных операций уже реализовано во внешних пакетах.

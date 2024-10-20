package code

import (
	"errors"
	"strconv"
)

func SumTwoIntegers(a, b string) (int, error) {
	var c, err = strconv.Atoi(a)
	var d, brr = strconv.Atoi(b)
	if err != nil || brr != nil {
		return 0, errors.New("invalid input, please provide two integers")
	}
	sum := c + d
	return sum, nil

}

//Напишите функцию SumTwoIntegers(a, b string) (int, error), которая получает две строки и, если это целые числа, возвращает их сумму. Если пользователь ввёл данные не целого типа, функция должна вернуть сообщение об ошибке (invalid input, please provide two integers).

package code

import (
	"fmt"
)

func DivideIntegers(a, b int) (float64, error) {
	var c float64
	if b == 0 {
		return 0, fmt.Errorf("division by zero is not allowed")
	} else {
		c = float64(a / b)
	}
	return c, nil
}

//Напишите функцию DivideIntegers(a, b int) (float64, error), которая возвращает результат деления первого числа на второе. Если второе число равно нулю, функция должна возвращать в качестве ответа 0 и сообщение об ошибке (division by zero is not allowed). Если второе число не равно нулю – верните в качестве ошибки nil.

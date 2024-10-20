package code

import (
	"errors"
)

func GetCharacterAtPosition(str string, position int) (rune, error) {
	tmp := []rune(str)
	if position < 0 || position >= len(tmp) {
		return 0, errors.New("position out of range")
	}
	return tmp[position], nil
}

//Напишите функцию GetCharacterAtPosition(str string, position int) (rune, error) для робота-помощника, которая получает на вход строку и целое число. Функция должна возвращать символ строки, который находится на позиции, указанной пользователем (и nil в качестве ошибки). Если пользователь ввёл число, которое выходит за пределы длины строки, функция должна возвращать в качестве ответа нулевую руну (0) сообщение об ошибке (position out of range).

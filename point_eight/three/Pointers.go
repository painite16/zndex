package code

import (
	"unicode"
)

func isLatin(input string) bool {
	for _, char := range input {
		if !unicode.Is(unicode.Latin, char) {
			return false
		}

	}
	return true
}

//Напишите функцию isLatin(input string) bool, которая принимает строку и выводит true, если все символы в строке латинские, false, если нет.

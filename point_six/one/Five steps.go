package code

func FiveSteps(array [5]int) [5]int {
	reversed := [5]int{}
	for i := 0; i < 5; i++ {
		reversed[i] = array[4-i]
	}
	return reversed
}

//Напоминаем, что ответ нужно начать с объявления пакета main и далее определить функцию FiveSteps(), определять функцию main() и вызывать FiveSteps() не надо – это сделает проверяющая система.

package code

func SumOfArray(array [6]int) int {

	sum := 0
	for _, value := range array {
		sum += value
	}
	return sum
}

//Напишите функцию SumOfArray(array [6]int) int которая получает массив из шести целочисленных элементов и возвращает сумму всех элементов массива.

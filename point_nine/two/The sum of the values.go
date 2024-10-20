package code

func SumOfValuesInMap(m map[int]int) int {
	var sum = 0
	for _, value := range m {
		sum += value
	}
	return sum
}

//Напишите функцию SumOfValuesInMap(m map[int]int) int, которая возвращает сумму значений в мапе.

//Примечания
//Например, если передать функции SumOfValuesInMap(m map[int]int) int мапу [10:38 3:19], то она должна вернуть число 57.

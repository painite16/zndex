package code

func FindMaxKey(m map[int]int) int {
	maxKey := 0
	for key := range m {
		if key > maxKey {
			maxKey = key
		}
	}
	return maxKey
}

//Напишите функцию FindMaxKey(m map[int]int) int, которая принимает мапу и возвращает значение наибольшего ключа.

//Примечания
//Например, если передать функции FindMaxKey(m map[int]int) int мапу [10:37 3:19], то она должна вернуть число 10

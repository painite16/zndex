package code

func FindMinMaxInArray(array [10]int) (int, int) {
	var min, max int
	max = array[0]
	min = array[0]

	for _, value := range array {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return max, min
}

//Гоша презирает гадания и пользуется генератором случайных чисел, чтобы создавать предсказуемые случайности. Воспользуйтесь им и напишите функцию FindMinMaxInArray(array [10]int) (int, int), которая получает массив из десяти случайных целочисленных значений. Далее функция должна находить максимальное и минимальное значения в массиве и возвращать их.

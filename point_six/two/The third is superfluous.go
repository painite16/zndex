package code

func ThirdElementInArray(array [6]int) int {

	for i := 0; i < len(array); i++ {
		sum := 0
		sum += i
		if sum == 2 {
			return array[2]
		} else if sum > 2 {
			break
		} else {
			continue
		}
	}
	return array[2]
}

//Напишите функцию ThirdElementInArray(array [6]int) int, которая получает массив из шести целочисленных элементов (пусть это будут номера игроков на форме) и возвращает третий по счёту элемент массива.

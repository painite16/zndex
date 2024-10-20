package code

func SliceCopy(nums []int) []int {
	nums = make([]int, len(nums), len(nums))
	return nums
}

//Дан слайс целых чисел nums. Этот слайс имеет емкость больше его длины. Создайте функцию SliceCopy(nums []int) []int, которая вернёт новый слайс длиной и ёмкостью, равной длине nums. Скопируйте в него значения из исходного слайса.

package code

func Clean(nums []int, x int) []int {
	pointer := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != x {
			nums[pointer] = nums[i]
			pointer++
		}
	}
	for i := pointer; i < len(nums); i++ {
		nums[i] = 0
	}
	return nums[:pointer]
}

//Дан неотсортированный слайс целых чисел. Напишите функцию Clean(nums []int, x int) ([]int), которая удаляет из исходного слайса все вхождения x. Важно сохранить порядок следования элементов и не использовать дополнительный слайс.

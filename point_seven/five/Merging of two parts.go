package code

func Mix(nums []int) []int {
	n := len(nums) / 2
	mix := make([]int, len(nums))
	idx := 0
	for i := 0; i < n; i++ {
		mix[idx] = nums[i]
		mix[idx+1] = nums[i+n]
		idx += 2
	}

	return mix
}

//Дан слайс nums, состоящий из 2n элементов в формате [x0,x1,...,xn,y0,y1,...,yn]. Создайте функцию Mix(nums []int) []int, которая вернёт слайс, содержащий значения в следующем порядке: [x0,y0,x1,y1,...,xn,yn].

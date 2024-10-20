package code

func Join(nums1, nums2 []int) []int {
	num := make([]int, 0, len(nums1)+len(nums2))
	for _, uko := range nums1 {
		num = append(num, uko)
	}
	for _, uko := range nums2 {
		num = append(num, uko)
	}
	return num
}

//Даны 2 слайса целых чисел nums1 и nums2. Создайте функцию Join(nums1, nums2 []int) []int, которя создаст новый слайс емкостью, вмещающей в себя ровно два слайса (ёмкость должна быть равна его длине). Скопируйте в него сначала значения nums1 затем nums2 и верните его.

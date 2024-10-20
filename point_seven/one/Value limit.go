package code

import "errors"

func UnderLimit(nums []int, limit int, n int) ([]int, error) {
	if n <= 0 {
		return nil, errors.New("n must be a positive number")
	}
	result := make([]int, 0)
	count := 0
	for _, num := range nums {
		if num < limit {
			result = append(result, num)
			count++
			if count == n {
				break
			}
		}

	}
	return result, nil
}

//Дан неотсортированный слайс целых чисел. Напишите функцию UnderLimit(nums []int, limit int, n int) ([]int, error), которая будет возвращать первые n (либо меньше, если остальные не подходят) элементов, которые меньше limit. В случае ошибки функция должна вернуть nil и описание ошибки.

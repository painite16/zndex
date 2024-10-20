package code

func DeleteLongKeys(m map[string]int) map[string]int {
	back := make(map[string]int)
	for j, i := range m {
		if len(j) >= 6 {
			back[j] = i
		}
	}
	return back
}

//Помогите Гоше почистить телефонную книгу. Создайте функцию DeleteLongKeys(m map[string]int) map[string]int, которая принимает мапу и возвращает новую мапу, где присутствуют только те ключи, у которых длина не меньше 6 символов.

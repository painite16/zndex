package code

func CountingSort(contacts []string) map[string]int {
	back := make(map[string]int)
	for _, i := range contacts {
		back[i]++

	}
	return back
}

//Помогите Гоше найти дубликаты. Для этого напишите функцию CountingSort(contacts []string) map[string]int, которая принимает слайс строк и возвращает мапу, где ключ - это элемент слайса, а значение - количество раз, сколько элемент встречается в слайсе.

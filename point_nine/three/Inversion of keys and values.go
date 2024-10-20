package code

func SwapKeysAndValues(m map[string]string) map[string]string {
	back := make(map[string]string)
	for key, value := range m {
		back[value] = key
	}
	return back
}

//Напишите функцию SwapKeysAndValues(m map[string]string) map[string]string, которая принимает мапу и возвращает новую мапу, где ключи и значения поменялись местами.

//Примечания
//Например, если передать функции SwapKeysAndValues(m map[string]string) мапу [Яндекс:+74957397000 Музей Яндекса:+74991101133], то она должна вернуть мапу [+74957397000:Яндекс +74991101133:Музей Яндекса].

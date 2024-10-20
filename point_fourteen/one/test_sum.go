package code

import (
	"testing"
)

func TestSum(t *testing.T) {
	testCases := []struct {
		a, b, expected int
	}{
		{1, 2, 3},
		{1, -1, 0},
		{0, 0, 0},
		{5, -2, 3},
	}
	for _, tc := range testCases {
		result := Sum(tc.a, tc.b)
		if result != tc.expected {
			t.Errorf("Ошибка")
		}
	}
}

//Функция Sum(a, b int) int (пакет main) возвращает результат суммирования чисел a и b. Напишите тест TestSum(t *testing.T) для проверки корректности работы.

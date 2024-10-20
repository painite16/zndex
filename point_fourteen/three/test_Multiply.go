package code

import (
	"testing"
)

func TestMultiply(t *testing.T) {
	testCases := []struct {
		a, b, expected int
	}{
		{1, 2, 2},
		{1, -1, -1},
		{5, 2, 10},
	}
	for _, tc := range testCases {
		result := Multiply(tc.a, tc.b)
		if result != tc.expected {
			t.Errorf("Ошибка")
		}
	}
}

//Функция Multiply(a, b int) int (пакет main) возвращает произведение двух переданных чисел. Напишите тест TestMultiply(t *testing.T) для проверки корректности работы.

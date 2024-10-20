package code

import (
	"testing"
)

func TestDeleteVowels(t *testing.T) {
	testCases := []struct {
		s, expected string
	}{
		{"apple", "ppl"},
		{"young", "yng"},
		{"people", "ppl"},
		{"life", "lf"},
	}
	for _, tc := range testCases {
		result := DeleteVowels(tc.s)
		if result != tc.expected {
			t.Errorf("Ошибка")
		}
	}
}

//Напишите тест для функции DeleteVowels(s string) string, которая должна удалять все гласные из строки английского языка (y не считается гласной).
//Используйте table driven testing.

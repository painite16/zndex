package code

import (
	"testing"
)

func TestGetUTFLength(t *testing.T) {
	testCases := []struct {
		s []byte
		b int
		e error
	}{
		{[]byte("Hello, 人日"), 9, error(nil)},
		{[]byte{0xff}, 0, ErrInvalidUTF8},
	}
	for _, tc := range testCases {
		result, err := GetUTFLength(tc.s)

		if result != tc.b {

			t.Errorf("Incorect length.Expected %d, got %d", tc.b, result)
		}
		if err != tc.e {
			t.Errorf("Expected errorto be nil, got: %v", err)
		}

	}
}

//Функция GetUTFLength(input []byte) (int, error) возвращает длину строки UTF8 и ошибку ErrInvalidUTF8 (в случае возникновения). Напишите тест, который бы проверял возвращаемые функцией значения.

//var ErrInvalidUTF8 = errors.New("invalid utf8")
//
//func GetUTFLength(input []byte) (int, error) {
// if !utf8.Valid(input) {
//  return 0, ErrInvalidUTF8
// }

// return utf8.RuneCount(input), nil
//}

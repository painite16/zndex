package code

import (
	"fmt"
)

func IsPowerOfTwoRecursive(N int) {
	if N%2 == 0 {
		fmt.Println("YES")
	} else if N == 1 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

//Гоша много раз умножал 2 на 2 и сбился. А теперь ему нужно понять, правильно он посчитал, или придётся начинать заново. Напишите функцию IsPowerOfTwoRecursive(N int), которая выводит на экран слово YES, если число N является точной степенью двойки, или слово NO в противном случае.

package code

import (
	"fmt"
)

func dinunt_treats() {
	var a int
	sum := 0
	fmt.Scanln(&a)
	for i := 1; i <= a; i++ {
		if i%3 != 0 && i%5 != 0 {
			sum += i
		}

	}
	fmt.Println(sum)
}

//В честь завершения проекта Гоша хочет угостить коллег пончиками. Но Гоша немного жадноват и очень хитер, поэтому он решил выкладывать пончики кучками — в каждой на один больше — и съедать те из них, которые поровну на всех не делятся. Он знает, что коллег может быть трое или пятеро.

//Напишите программу для подсчёта, сколько пончиков в итоге съест Гоша. (А вот опасна ли такая жадность для его жизни, придётся определять врачам...) Она должна запрашивать у пользователя число и выводить на экран сумму всех чисел от 1 до этого числа, кроме чисел, которые делятся на 3 или на 5 с помощью цикла for.

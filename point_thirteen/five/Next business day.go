package code

import (
	"time"
)

func NextWorkday(start time.Time) time.Time {
	for {
		start = start.AddDate(0, 0, 1)
		if start.Weekday() != time.Saturday && start.Weekday() != time.Sunday {
			break
		}
	}
	return start
}

//Напишите функцию NextWorkday(start time.Time) time.Time, которая вычисляет дату следующего рабочего дня (исключая выходные). Учитывая дату начала, функция должна возвращать дату следующего рабочего дня.

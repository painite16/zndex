package code

import (
	"strconv"
	"time"
)

func TimeAgo(pastTime time.Time) string {
	dif := time.Now().Sub(pastTime)
	hor := int(dif.Hours())
	days := hor / 24
	month := int(dif.Hours() / (24 * 30))
	var txt string
	switch {
	case month > 0:
		txt = strconv.Itoa(month) + " month ago"
		return txt
	case days > 0:
		txt = strconv.Itoa(days) + " days ago"
		return txt

	case hor > 0:
		txt = strconv.Itoa(hor) + " hours ago"
		return txt
	}
	return txt
}

//Напишите функцию TimeAgo(pastTime time.Time) string, который вычисляет время, прошедшее с момента заданного значения time.Time, и возвращает удобочитаемую строку, указывающую, как давно это было.

//Примечания
//Пример работы функции:

//    pastTime := time.Date(2023, 10, 23, 2, 41, 49, 0, time.UTC)
//    result := TimeAgo(pastTime)
//   fmt.Println(result) // Output: 1 month ago
//Пример работы с часами:

//    pastTime := time.Now().Add(-2 * time.Hour)
///    timeAgo := TimeAgo(pastTime)
//    expected := "2 hours ago"

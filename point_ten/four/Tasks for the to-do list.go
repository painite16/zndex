package code

import (
	"time"
)

type Task struct {
	summary     string
	description string
	deadline    time.Time
	priority    int
}

func (s Task) IsOverdue() bool {
	return time.Now().After(s.deadline)
}
func (s Task) IsTopPriority() bool {
	return s.priority == 4 || s.priority == 5
}

//Создайте структуру Task:

//1. summary - короткий заголовок (тип string)

//2. description - подробное описание (тип string)

//3. deadline - дедлайн для задачи (тип time.Time (стандартный пакет time))
//
//4. priority - приоритет задачи (тип int)

//Для этой структуры реализуйте метод IsOverdue, которая возвращает true, если дедлайн задачи уже прошел и false в ином случае. Подсказка: воспользуйтесь функцией time.Now().

//Ещё запрограммируйте метод IsTopPriority. Функция возвращает true, если приоритет 4 или 5, и false, если меньше.

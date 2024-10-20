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

type Note struct {
	title string
	text  string
}
type ToDoList struct {
	name  string
	tasks []Task
	notes []Note
}

func (td ToDoList) TasksCount() int {
	return len(td.tasks)
}

func (td ToDoList) NotesCount() int {
	return len(td.notes)
}
func (td ToDoList) CountTopPrioritiesTasks() int {
	sum := 0
	for _, task := range td.tasks {
		if task.priority >= 5 {
			sum++
		}
	}
	return sum
}
func (td ToDoList) CountOverdueTasks() int {
	sum := 0
	for _, task := range td.tasks {
		if !time.Now().Before(task.deadline) {
			sum++
		}
	}
	return sum
}

//Напишите структура Note (сущность заметок, у которых в отличие от задач нет чётких дедлайнов и приоритета):

//1. title - заголовок (тип string)

//2. text - текст заметок (тип string)

//Создайте структуру ToDoList с такими полями:

//1. name - название списка (тип string)

//2. tasks - список дел на сегодня (тип слайс структур Task (из предыдущего задания))

//3. notes - список дополнительных заметок (тип слайс структур Note)

//Для этой структуры реализуйте методы TasksCount и NotesCount, которые возвращают общее количество задач и заметок соответственно.

//Также реализуйте метод CountTopPrioritiesTasks, который возвращает количество приоритетных задач. А также метод CountOverdueTasks, который возвращает количество просроченных задач.

//Сама структура Task и все её методы из предыдущего задания также должны быть реализованы в этом.

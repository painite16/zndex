package code

type Student struct {
	name            string
	solvedProblems  int
	scoreForOneTask float64
	passingScore    float64
}

func (s Student) IsExcellentStudent() bool {

	sum := float64(s.solvedProblems) * s.scoreForOneTask
	if sum >= s.passingScore {
		return true
	} else {
		return false
	}
}

//Создайте структуру Student с полями name (строка), solvedProblems — количество решённых задач (целое число), scoreForOneTask— количество баллов за одну задачу; будем считать, что все задачи дают одинаковые баллы (число с плавающей точкой) и passingScore — проходной балл в следующий модуль (число с плавающей точкой).

//Создайте метод IsExcellentStudent для этой структуры, который возвращает true, если ученик проходит по баллам в следующий модуль и false в ином случае.

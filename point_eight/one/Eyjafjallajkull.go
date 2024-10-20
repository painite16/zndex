package code

func StringLength(input string) int {
	var intr int
	if input != "" {
		intr = len(input)

	} else {
		intr = 0
	}
	return intr
}

//Напишите функцию StringLength(input string) int, которая принимает строку и возвращает её длину.
//Например, если передать функции StringLength(input string) int строку "Alaid", то она должна вернуть число 5, а если строку "", то – 0.

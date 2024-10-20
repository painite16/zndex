package code

import (
	"fmt"
)

type Cat struct {
}
type Dog struct {
}

type Animal interface {
	MakeSound() string
}

func (c Dog) MakeSound() {
	fmt.Println("Гав!")
}
func (r Cat) MakeSound() {
	fmt.Println("Мяу!")
}

//"Заведи кота или собаку, будет весело!" — говорили друзья. А Гоша ответил, что надо подумать, и пошёл писать интерфейс.

//Создайте интерфейс Animal с методом MakeSound, который будет выводить звук, издаваемый животным. Создайте структуры Dog и Cat, которые будут реализовывать этот интерфейс и издавать соответствующие звуки (выводить на экран) – "Гав!" и "Мяу!".

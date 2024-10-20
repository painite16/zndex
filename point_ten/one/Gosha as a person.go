package code

import "fmt"

type Person struct {
	name    string
	age     int
	address string
}

func (p Person) Print() {
	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
	fmt.Println("Address:", p.address)
}

//Создайте структуру Person с полями name, age и address. Создайте метод Print для этой структуры, который будет выводить информацию о человеке на экран в формате:

//Name: Гоша

//Age: 21

//Address: Ясногорск

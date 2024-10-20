package code

import (
	"fmt"
	//"math"
)

type Employee struct {
	name     string
	position string
	salary   float64
	bonus    float64
}

func (e Employee) CalculateTotalSalary() {
	fmt.Println("Employee:", e.name)
	fmt.Println("Position:", e.position)
	//sum := e.salary + e.bonus
	//sum = math.Floor((sum * 100) / 100)
	fmt.Printf("Total Salary: %.2f", (e.salary + e.bonus))
}

//Создайте структуру Employee с полями name (string), position (string), salary (float64) и bonus (float64). Создайте метод CalculateTotalSalary для этой структуры, который будет выводить общую зарплату работника (salary + bonus) по следующему формату:

//Employee: Arseniy

//Position: backend developer

//Total Salary: 1000.02

//Обратите внимание, что Total Salary нужно округлять до 2 знаков после запятой.

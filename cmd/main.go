package main

import (
	"awesomeProject2/maths"
	"fmt"
)

func main() {
	var number int
	fmt.Print("Введите число для вычисления факториала: ")
	fmt.Scanln(&number)

	if number < 0 {
		fmt.Println("Факториал отрицательного числа не определен.")
		return
	}

	result := maths.Factorial(number)
	fmt.Printf("Факториал числа %d равен %d\n", number, result)

	fmt.Print("Введите число для вычисления суммы ряда: ")
	fmt.Scanln(&number)

	result = maths.SumOfSeries(number)
	fmt.Printf("Сумма чисел от 1 до %d равна %d\n", number, result)
}

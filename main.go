package main

import "fmt"

func main() {
	//Используя цикл написать программу, которая выводит на экран таблицу значений функции y = 5 - x2/2 на отрезке [-5; 5] с шагом 0.5.

	for i := -5.0; i <= 5; i += 0.5 {
		y := 5 - i*i/2
		fmt.Println("x = ", i, "y = ", y)
	}
}

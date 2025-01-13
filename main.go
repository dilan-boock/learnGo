package main

import "fmt"

func main() {
	//С клавиатуры вводится целое число. Определить, из каких цифр оно состоит, то есть вывести на экран отдельные цифры числа.

	var chislo int
	fmt.Println("Введите число: ")
	fmt.Scan(&chislo)

	ein := chislo % 1e3   // используем деление по модулю, чтобы «считать» с конца
	zwain := chislo / 1e3 // отрезаем после третьей позиции
	fmt.Println("Число: ", chislo)
}

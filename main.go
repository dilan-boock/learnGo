package main

import (
	"fmt"
)

//Написать программу, подсчитывающую количество четных и нечетных цифр числа.
//
//Описание переменных:
//
//even - количество четных цифр
//uneven - количество нечетных цифр

func main() {
	var chislo int
	fmt.Println("Введите число: ")
	fmt.Scan(&chislo)
	even := 0
	uneven := 0

	for chislo != 0 {
		num := chislo % 10
		if num%2 == 0 {
			even += num
		}
		if num%2 != 0 {
			uneven += num
		}
		fmt.Println(num)
		chislo /= 10
		//fmt.Println(chislo)
	}
	fmt.Println("Сумма четных чисел: ", even)
	fmt.Println("Сумма нечетных чисел: ", uneven)

}

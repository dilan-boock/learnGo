package main

import (
	"fmt"
	"math"
)

func main() {
	//Написать программу, которая возводит число в целочисленную степень. Число и степень вводятся с клавиатуры.
	var chislo, stepen int
	fmt.Println("Введите число: ")
	fmt.Scan(&chislo)
	fmt.Println("Введите степень: ")
	fmt.Scan(&stepen)
	result := math.Pow(float64(stepen), float64(chislo)) //хз почему оно хуево степень считает
	fmt.Println("Число ", chislo, " в степени ", stepen, " = ", result)
}

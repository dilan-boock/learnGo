package main

import "fmt"

func main() {
	//Вывести все квадраты натуральных чисел, не превосходящие данного числа N.
	//Например, если N = 50, то на экран должен быть выведен ряд 1 4 9 16 25 36 49

	n := 50

	for i := 1; i < n; i++ {
		qadrat := i * i
		if qadrat <= n {
			fmt.Println(qadrat)
		}
	}
}

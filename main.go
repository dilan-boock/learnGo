package main

import "fmt"

func main() {
	//Вывести на экран кубы чисел от A до B, которые вводит пользователь.
	var A int
	var B int
	fmt.Println("Введите число A: ")
	fmt.Scan(&A)
	fmt.Println("Введите число B: ")
	fmt.Scan(&B)

	for i := A; i <= B; i++ {
		cub := i * i * i
		fmt.Println(cub)
	}
}

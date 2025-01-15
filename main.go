package main

import "fmt"

//Напишите программу, которая считает количество элементов (пар ключ-значение) в карте.

func main() {

	var massiv = map[string]int{
		"One":   1,
		"Two":   2,
		"Three": 3,
		"Four":  4,
		"Five":  5,
		"Ein":   4}
	n := 0

	for i := 0; i < len(massiv); i++ {
		n++
	}
	fmt.Println("В массиве ", n, "пар ключ-значение")
}

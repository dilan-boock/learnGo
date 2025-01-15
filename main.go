package main

import "fmt"

//Реализуйте функцию, которая принимает карту и пороговое значение. Возвращайте новую карту, содержащую только те элементы, где значение больше порога.

func main() {

	var massiv = map[string]float32{
		"Milk":  10,
		"Cooke": 20,
		"Cofie": 30,
		"Tea":   40,
		"Fish":  51}
	fmt.Println("Старый массив: ", massiv)

	slice := map[string]float32{}

	for i := range massiv {
		for massiv[i] > 30 {
			slice[i] = massiv[i]
			break
		}
	}
	fmt.Println("Массив после порогового значения: ", slice)
}

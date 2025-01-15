package main

import "fmt"

//Реализуйте функцию, которая принимает карту и пороговое значение. Возвращайте новую карту, содержащую только те элементы, где значение больше порога.

func returnMap(massiv map[string]float32, n float32) map[string]float32 {
	slice := map[string]float32{}

	for i := range massiv {
		for massiv[i] > n {
			slice[i] = massiv[i]
			break
		}
	}
	return slice
}

func main() {

	var massiv = map[string]float32{
		"Milk":  10,
		"Cooke": 20,
		"Cofie": 30,
		"Tea":   40,
		"Fish":  51}
	fmt.Println("Старый массив: ", massiv)
	n := float32(30)
	slice := returnMap(massiv, n)
	fmt.Println("Массив после порогового значения: ", slice)
}

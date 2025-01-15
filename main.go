package main

import "fmt"

//Напишите функцию, которая создаёт "обратную" карту. Например, из карты вида {"a": 1, "b": 2, "c": 2} сделайте карту вида {1: ["a"], 2: ["b", "c"]}

func returnMap(massiv map[string]float32) map[float32]string {
	var n float32
	var s string
	slice := map[float32]string{}

	for i := range massiv {
		n = massiv[i]
		s = i
		slice[n] = s
	}
	return slice
}

func main() {

	var massiv = map[string]float32{
		"Milk":  100,
		"Cooke": 2540,
		"Cofie": 321,
		"Tea":   405,
		"Fish":  51}
	fmt.Println("Старый массив: ", massiv)
	slice := returnMap(massiv)
	fmt.Println("Новый массив: ", slice)
}

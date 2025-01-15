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

//Напишите функцию, которая создаёт "обратную" карту. Например, из карты вида {"a": 1, "b": 2, "c": 2} сделайте карту вида {1: ["a"], 2: ["b", "c"]}.
// TODO: исправить поведение со значениями-дублями
func reverseMap() {
	// TODO: исправить имена переменных, например sourceMap, res(result)
	var massiv = map[string]float32{
		"Milk":   100,
		"Cooke":  2540,
		"Cofie":  321,
		"Tea":    405,
		"Fish":   51,
		"BlaBla": 321}
	fmt.Println("Старый массив: ", massiv)

	var n float32
	var s string
	slice := map[float32]string{}

	// не забывать, что есть удобный синтаксис for k, v := range ... (key, value)
	for i := range massiv {
		n = massiv[i]
		s = i
		slice[n] = s
	}
	fmt.Println("Новый массив: ", slice)
}

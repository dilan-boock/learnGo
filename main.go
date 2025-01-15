package main

import "fmt"

//Реализуйте функцию, которая принимает карту с названиями продуктов и их ценами, а также скидку в процентах. Функция должна обновлять карту, уменьшая все цены на заданный процент.

func main() {

	var massiv = map[string]float32{
		"Milk":  100,
		"Cooke": 2540,
		"Cofie": 321,
		"Tea":   405,
		"Fish":  51}
	var n float32

	for i := range massiv {
		n = massiv[i]
		massiv[i] = n - (n/100)*10
	}
	fmt.Println("Новый список цен: ", massiv)
}

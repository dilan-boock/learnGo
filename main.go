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

func valueMap(chislo int) (int, int) {
	even := 0
	uneven := 0

	for chislo != 0 {
		num := chislo % 10
		if num%2 == 0 {
			even++
		}
		if num%2 != 0 {
			uneven++
		}
		fmt.Println(num)
		chislo /= 10
		//fmt.Println(chislo)
	}
	return even, uneven
}

func modifyMap(massiv map[string]float32) map[string]float32 {
	var n float32

	for i := 0; i < len(massiv); i++ {
		n++
		for i := range massiv {
			n = massiv[i]
			massiv[i] = n - (n/100)*10
		}
	}
	return massiv
}

func countMap(massiv map[string]int) int {
	n := 0

	for i := 0; i < len(massiv); i++ {
		n++
	}
	return n
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

	var chislo int
	fmt.Println("Введите число: ")
	fmt.Scan(&chislo)
	even, uneven := valueMap(chislo)
	fmt.Println("Сумма четных чисел: ", even)
	fmt.Println("Сумма нечетных чисел: ", uneven) //Подсчет колва

	var massiv2 = map[string]int{
		"One":   1,
		"Two":   2,
		"Three": 3,
		"Four":  4,
		"Five":  5,
		"Ein":   4}
	n := countMap(massiv2)
	fmt.Println("В массиве ", n, "пар ключ-значение") //ключ-значение

	massivRes := modifyMap(massiv)
	fmt.Println("Новый список цен: ", massivRes)
}

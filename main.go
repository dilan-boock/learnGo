package main

import "fmt"

//Напишите функцию, которая создаёт "обратную" карту. Например, из карты вида {"a": 1, "b": 2, "c": 2} сделайте карту вида {1: ["a"], 2: ["b", "c"]}

func returnMap(massiv map[string]float32) (map[float32]string, error) {
	var n float32
	var s string
	slice := map[float32]string{}

	for i := range massiv {
		n = massiv[i]
		s = i
		slice[n] = s
	}
	return slice, nil
}

func valueMap(even int, uneven int, chislo int) (int, int, error) {

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
	return even, uneven, nil
}

func modifyMap(massiv map[string]float32, v float32) (map[string]float32, error) {

	for i := 0; i < len(massiv); i++ {
		v++
		for i := range massiv {
			v = massiv[i]
			massiv[i] = v - (v/100)*10
		}
	}
	return massiv, nil
}

func countMap(massiv2 map[string]int, v int) (int, error) {

	for i := 0; i < len(massiv2); i++ {
		v++
	}
	return v, nil
}

func main() {
	fmt.Println("Подсчёт частоты элементов")
	input := []string{"apple", "banana", "apple", "orange", "banana", "apple"}
	fmt.Println("input: ", input)
	output := periodMap(input)
	fmt.Println("output: ", output)

	fmt.Println("\nОбъединение двух map")
	map1 := map[int]int{1: 10, 2: 20, 3: 30}
	map2 := map[int]int{2: 15, 3: 35, 4: 40}
	fmt.Println("map1: ", map1)
	fmt.Println("map2: ", map2)
	output2 := mergeMap(map1, map2)
	fmt.Println("output: ", output2)

	fmt.Println("\nПроверка на подмножество")
	map3 := map[string]int{"a": 3, "b": 5, "c": 2}
	map4 := map[string]int{"a": 2, "b": 4}
	fmt.Println("map1: ", map3)
	fmt.Println("map2: ", map4)
	output3 := subsetMap(map3, map4)
	fmt.Println("output: ", output3)

	fmt.Println("\nГруппировка по длине строки")
	input2 := []string{"cat", "dog", "apple", "bat", "pear", "down"}
	fmt.Println("input: ", input2)
	output4 := lenghtMap(input2)
	fmt.Println("output: ", output4)

	fmt.Println("\nНайти максимальное значение")
	input3 := map[string]int{"a": 10, "b": 20, "c": 15}
	fmt.Println("input: ", input3)
	outputKey, outputNum := maxMap(input3)
	fmt.Println("output: ", "'", outputKey, "',", outputNum)

	fmt.Println("\nСравнение двух map")
	map5 := map[string]int{"a": 1, "b": 2, "c": 3}
	map6 := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println("map1: ", map5)
	fmt.Println("map2: ", map6)
	output5 := comparisonMap(map5, map6)
	fmt.Println("output: ", output5)
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

//Реализуйте функцию, которая принимает карту с названиями продуктов и их ценами, а также скидку в процентах. Функция должна обновлять карту, уменьшая все цены на заданный процент.

func discounts() {
	var massiv = map[string]float32{
		"Milk":  100,
		"Cooke": 2540,
		"Cofie": 321,
		"Tea":   405,
		"Fish":  51}
	var n float32

	discount := 10
	for i := range massiv { // todo исправить на for k, v :=
		n = massiv[i]
		massiv[i] = n / 100 * (100 - discount) // TODO исправить
	}

	fmt.Println("Новый список цен: ", massiv)
}

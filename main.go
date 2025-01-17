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
	return slice, error()
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
	return even, uneven, error()
}

func modifyMap(massiv map[string]float32, v float32) (map[string]float32, error) {

	for i := 0; i < len(massiv); i++ {
		v++
		for i := range massiv {
			v = massiv[i]
			massiv[i] = v - (v/100)*10
		}
	}
	return massiv, error()
}

func countMap(massiv2 map[string]int, v int) (int, error) {

	for i := 0; i < len(massiv2); i++ {
		v++
	}
	return v, error()
}

func main() {

}

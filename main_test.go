package main

import (
	"fmt"
	"testing"
)

func TestEnglish(t *testing.T) {
	var massiv = map[string]float32{
		"Milk":  100,
		"Cooke": 2540,
		"Cofie": 321,
		"Tea":   405,
		"Fish":  51}
	fmt.Println("Старый массив: ", massiv)

	var massiv2 = map[string]int{
		"One":   1,
		"Two":   2,
		"Three": 3,
		"Four":  4,
		"Five":  5,
		"Ein":   4}

	chislo := 15
	even := 0
	uneven := 0
	v := 0
	var n float32

	test1, err := returnMap(massiv)
	even, uneven, err = valueMap(even, uneven, chislo)
	test3, err := modifyMap(massiv, n)
	test4, err := countMap(massiv2, v)

	if err != nil {
		t.Errorf("Should not produce an error")
		fmt.Println("Новый массив: ", test1)

		fmt.Println("Сумма четных чисел: ", even)
		fmt.Println("Сумма нечетных чисел: ", uneven) //Подсчет колва

		fmt.Println("В массиве ", test4, "пар ключ-значение") //ключ-значение

		fmt.Println("Новый список цен: ", test3)
	}

	if err == nil {
		t.Errorf("Result was incorrect, got: %s, want: %s.", err)
	}
}

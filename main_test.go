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

	n := 0

	v := 0
	test1, err := returnMap(massiv, v)
	test2, err := valueMap(chislo)
	test3, err := modifyMap(massiv, n)
	test4, err := countMap(massiv2)

	if err != nil {
		t.Errorf("Should not produce an error")
	}

	if expected != actual {
		t.Errorf("Result was incorrect, got: %s, want: %s.", actual, expected)
	}
}

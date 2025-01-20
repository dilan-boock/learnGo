package main

func periodMap(input []string) map[string]int {
	//Напишите функцию, которая принимает срез строк и возвращает map, где ключи — уникальные строки, а значения — их частота в срезе.
	//
	//	Пример:
	//input := []string{"apple", "banana", "apple", "orange", "banana", "apple"}
	//output := map[string]int{"apple": 3, "banana": 2, "orange": 1}

	str := " "
	n := 1
	output := map[string]int{}
	for i := 0; i < len(input); i++ {
		str = input[i]
		_, ok := output[str]
		if !ok {
			output[str] = n
		}
		if ok {
			output[str]++
		}
	}

	return output
}

func mergeMap(map1 map[int]int, map2 map[int]int) map[int]int {
	//Реализуйте функцию, которая принимает две мапки с целыми ключами и значениями. Функция должна возвращать новую мапку, где значения с одинаковыми ключами суммируются.
	//
	//	Пример:
	//map1 := map[int]int{1: 10, 2: 20, 3: 30}
	//map2 := map[int]int{2: 15, 3: 35, 4: 40}
	//output := map[int]int{1: 10, 2: 35, 3: 65, 4: 40}

	output := map[int]int{}
	for i := range map1 {
		output[i] = map1[i]
		for i := range map2 {
			_, ok := output[i]
			if !ok {
				output[i] = map2[i]
			}
			if ok {
				output[i] = map1[i] + map2[i]
			}
		}
	}

	return output
}

func subsetMap(map3 map[string]int, map4 map[string]int) bool {
	//Напишите функцию, которая принимает две map[string]int. Верните true, если вторая мапка является подмножеством первой.
	//
	//	Пример:
	//
	//map1 := map[string]int{"a": 3, "b": 5, "c": 2}
	//map2 := map[string]int{"a": 2, "b": 4}
	//output := true

	for key1 := range map3 {
		_, ok := map4[key1]
		if !ok {
			return false
		}
	}
	return true
}

func lenghtMap(input []string) map[int][]string {
	//Реализуйте функцию, которая принимает срез строк и группирует их по длине. Возвращайте карту, где ключ — длина строки, а значение — срез строк этой длины.
	//
	//	Пример:
	//input := []string{"cat", "dog", "apple", "bat", "pear"}
	//output := map[int][]string{3: {"cat", "dog", "bat"}, 5: {"apple", "pear"}}
	output := map[int][]string{}

	for i := range input {
		str := len(input[i])
		_, ok := output[str]
		if !ok {
			output[str] = []string{input[i]}
		}
		if ok {
			output[str] = append(output[str], input[i])
		}
	}

	return output
}

func maxMap(input map[string]int) (string, int) {
	//Напишите функцию, которая принимает карту map[string]int и возвращает ключ и значение с максимальным значением. Если карта пуста, верните ошибку.
	//
	//	Пример:
	//input := map[string]int{"a": 10, "b": 20, "c": 15}
	//output := "b", 20
	n := 0
	var outputKey string
	var outputNum int
	for i, value := range input {
		if value > n {
			n = value
			outputKey = i
		}
		outputNum = n
	}
	return outputKey, outputNum
}

func comparisonMap(map5 map[string]int, map6 map[string]int) bool {
	//Напишите функцию, которая принимает две мапы map[string]int и проверяет, эквивалентны ли они (имеют одинаковые ключи и значения).
	//Пример:
	//map1 := map[string]int{"a": 1, "b": 2, "c": 3}
	//map2 := map[string]int{"a": 1, "b": 2, "c": 3}
	//output := true

	if len(map5) != len(map6) {
		return false
	}

	for key1, value1 := range map5 {
		value2, ok := map6[key1]
		if !ok || value1 != value2 {
			return false
		}
	}

	return true
}

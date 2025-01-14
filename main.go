package main

import (
	"fmt"
	"strconv"
)

func main() {
	//С клавиатуры вводится целое число. Определить, из каких цифр оно состоит, то есть вывести на экран отдельные цифры числа.

	var chislo int
	fmt.Println("Введите число: ")
	fmt.Scan(&chislo)

	// 12345
	str := strconv.Itoa(chislo)
	//m := []byte(str)

	//c := str[len(str)-1]
	m := make([]byte, len(str))
	copy(m, str)
	//for i := 0; i < len(str); i++ {
	//	m[i] = str[len(str)-i-1]
	//}
	fmt.Println(string(m))
	//[len(str)-1] = str[0]
	//str[0] = c

	//fmt.Println(str[0])

	for chislo != 0 {
		num := chislo % 10
		fmt.Println(num)
		chislo /= 10
		//fmt.Println(chislo)
	}
}

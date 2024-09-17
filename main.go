package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func add(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}

func multiply(a int, b int) int {
	return a * b
}

func divide(a int, b int) int {
	return a / b
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func trans_to_rom(target int) string {

	rom_nums := map[string]int{"C": 100, "XC": 90, "L": 50, "XL": 40, "X": 10, "IX": 9, "VIII": 8, "VII": 7, "VI": 6, "V": 5, "IV": 4, "III": 3, "II": 2, "I": 1}

	var result_num string

	for i := 0; target > 0; i++ {

		// Создаем срез для хранения пар ключ-значение
		type kv struct {
			Key   string
			Value int
		}

		var sortedPairs []kv
		for k, v := range rom_nums {
			sortedPairs = append(sortedPairs, kv{k, v})
		}

		// Сортируем срез по значениям от наибольшего
		sort.Slice(sortedPairs, func(i, j int) bool {
			return sortedPairs[i].Value > sortedPairs[j].Value
		})

		// Ищем первое значение, которое меньше или равно заданному
		var result *kv
		for _, pair := range sortedPairs {
			if pair.Value <= target {
				result = &pair
				break
			}
		}

		target = target - result.Value
		result_num = result_num + result.Key

	}

	return result_num

}

func main() {

	arab_nums := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}

	rom_nums := map[string]int{"X": 10, "IX": 9, "VIII": 8, "VII": 7, "VI": 6, "V": 5, "IV": 4, "III": 3, "II": 2, "I": 1}

	operations := []string{"+", "-", "*", "/"}

	fmt.Println("Введите пример используя арабские, либо римские целые числа от 1 до 10 включительно. \nДоступные арифметические операции: +, -, *, /.")

	//Получаем пример и переводим его в массив для разделения по переменным
	reader := bufio.NewReader(os.Stdin)
	sum, _ := reader.ReadString('\n')
	sum2 := strings.Split(sum, " ")

	//Присваиваем переменным значения из массива
	a := sum2[0]
	operation := sum2[1]
	b := strings.TrimSpace(sum2[2])

	//Проверяем входят-ли переменные в допустимые значения
	found_a_arab := contains(arab_nums, a)
	found_b_arab := contains(arab_nums, b)
	found_operation := contains(operations, operation)
	found_a_rom := rom_nums[a]
	found_b_rom := rom_nums[b]

	if found_a_arab == true && found_b_arab == true || found_a_rom != 0 && found_b_rom != 0 {
	} else {
		panic("Можно использовать либо только арабские, либо только римские числа от 1 до 10 включительно.")
	}

	if found_operation == true {
	} else {
		panic("Недопустимая операция")
	}

	//Ну и наконец решаем пример
	result := 0
	if found_a_arab == true {

		a_int, err := strconv.Atoi(a)
		if err != nil {
			panic(err)
		}
		b_int, err2 := strconv.Atoi(b)
		if err2 != nil {
			panic(err)
		}

		switch operation {
		case "+":
			result = add(a_int, b_int)
		case "-":
			result = subtract(a_int, b_int)
		case "*":
			result = multiply(a_int, b_int)
		case "/":
			result = divide(a_int, b_int)
		}
		fmt.Println(result)
	} else {
		a_int := rom_nums[a]
		b_int := rom_nums[b]

		switch operation {
		case "+":
			result = add(a_int, b_int)
		case "-":
			result = subtract(a_int, b_int)
		case "*":
			result = multiply(a_int, b_int)
		case "/":
			result = divide(a_int, b_int)
		}

		if result < 0 {
			panic("В римской системе счета отсутствуют отрицательные числа")
		} else if result > 0 {
			result := trans_to_rom(result)
			fmt.Println(result)
		} else {
			panic("В римской системе счета отсутствует ноль")
		}
	}
}

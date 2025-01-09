package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Maps
	/*
		myMap1 := make(map[int]int) // make
		fmt.Println(myMap1)

		myMap2 := map[string]int{"one": 1, "two": 2} // init
		fmt.Println(myMap2)

		var myMap3 map[string]string // no init
		fmt.Println(myMap3)

		for key, value := range myMap2 {
			fmt.Println(key, value)
		}
	*/

	// Homework

	// 1
	/*
		myMap := make(map[string]int)
		myMap["apple"] = 3
		fmt.Println(myMap)
	*/

	// 2
	/*
		myMap := map[string]int{"one": 1, "two": 2}
		fmt.Println(myMap["one"])
	*/

	// 3
	/*
		myMap := map[string]int{"apple": 3, "orange": 2}
		delete(myMap, "apple")
		fmt.Println(myMap)
	*/

	// 4
	/*
		myMap := map[string]int{"apple": 3, "orange": 2}
		_, ok := myMap["pineapple"]
		if ok == true {
			fmt.Println("Есть")
		} else {
			fmt.Println("Не Есть")
		}
	*/

	// 5
	/*
		myMap := map[string]int{"apple": 3, "orange": 2}
		fmt.Println(len(myMap))
	*/

	// 6
	/*
		myMap := map[string][]int{"even": {2, 4, 6}, "odd": {1, 3, 5}}
		for key, value := range myMap {
			fmt.Printf("%v %v\n", key, value)
		}
	*/

	// 7
	/*
		myMap := map[string][]int{"even": {2, 4, 6}, "odd": {1, 3, 5}}
		fmt.Println(myMap)
	*/

	// 8
	/*
		myMap := map[string][]int{"even": {2, 4, 6}, "odd": {1, 3, 5}}
		for _, value := range myMap {
			for i, _ := range value {
				fmt.Printf("%v\n", value[i])
			}
		}
	*/

	// 9
	/*
		myMap := map[string][]int{"even": {2, 4, 6}, "odd": {1, 3, 5}}
		key := "odd"
		_, ok := myMap[key]
		if ok == true {
			fmt.Println(len(myMap[key]))
		}
	*/

	// 10
	/*
		myMap1 := map[string]int{"apple": 3, "banana": 5}
		myMap2 := map[string]int{"orange": 2, "grape": 4}
		myMap3 := make(map[string]int)

		for key, value := range myMap1 {
			myMap3[key] = value
		}
		for key, value := range myMap2 {
			myMap3[key] = value
		}

		fmt.Print(myMap3)
	*/

	// 11
	/*
		myMap := make(map[string]string)

		for i := 0; i < 3; i++ {
			var key, value string
			fmt.Println("Введите ключ: ")
			fmt.Scan(&key)
			fmt.Println("Введите значение: ")
			fmt.Scan(&value)

			myMap[key] = value
		}

		fmt.Println(myMap)
	*/

	// 12
	/*
		myMap := make(map[int]string)

		for i := 0; i < 5; i++ {
			var value string
			fmt.Println("Введите значение: ")
			fmt.Scan(&value)

			myMap[i] = value
		}

		fmt.Println(myMap)
	*/

	// 13
	/*
		myMap := make(map[string]int)

		for i := 0; i < 5; i++ {
			var value string
			fmt.Print("Введите значение: ")
			fmt.Scan(&value)

			myMap[value]++
		}

		fmt.Println(myMap)
	*/

	// 14
	/*
		var numbers []int // Срез для хранения чисел

		for {
			var num int
			fmt.Print("Введите число: ")
			fmt.Scan(&num)

			if num < 0 {
				break
			}

			numbers = append(numbers, num)
		}

		fmt.Println(numbers)
	*/

	// 15
	myMap := make(map[string]string)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Введите ключ: ")
		scanner.Scan()
		key := scanner.Text()

		if key == "" {
			break
		}

		fmt.Print("Введите значение: ")
		scanner.Scan()
		value := scanner.Text()

		myMap[key] = value
	}

	fmt.Println(myMap)

	// Algorithm
	/*
		myMap := make(map[string]int)

		var input string
		fmt.Print("Введите строку: ")
		fmt.Scanln(&input)

		for _, value := range input {
			myMap[string(value)] += 1
		}

		count := 0
		letter := ""

		fmt.Println(myMap)

		test := 0
		for key, value := range myMap {
			if test == 0 {
				test += 1
				count = value
				letter = key
			} else {
				if value > count {
					letter = key
					count = value
				}
			}
		}
		fmt.Println(letter, count)
	*/
}

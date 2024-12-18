package main

import (
	"fmt"
	"log"
)

func main() {

	// Lesson
	/*
		for i := 1; i <= 10; i++ {
			for ii := 1; ii <= 10; ii++ {
				fmt.Printf("%d ", i*ii)
			}
			fmt.Printf("\n")
		}

		myArray := [5]int{1, 2}
		myArray := [...]int{1, 2}
		len(myArray)
	*/

	// 1
	/*
		var num int
		fmt.Println("Введите индекс: ")

		_, err := fmt.Scan(&num)
		if err != nil {
			log.Println(err)
			return
		}

		nums := [...]int{6, 19, 26, 9, 46, 8, 5, 65, 90, 25}
		if num > 0 && num < len(nums) {
			fmt.Print(nums[num])
		} else {
			fmt.Printf("Error")
		}
	*/

	// 2
	/*
		nums := [...]int{4, 67, 9, 12, 6, 45, 11, 7, 23, 40}
		for i := 0; i < len(nums); i++ {
			fmt.Println(nums[i])
		}
	*/

	// 3
	/*
		var a1, a2, a3, a4, a5 int
		fmt.Println("Введите числа: ")
		fmt.Scan(&a1, &a2, &a3, &a4, &a5)
		nums := [...]int{a1, a2, a3, a4, a5}

		for i, value := range nums {
			fmt.Printf("%d - %d\n", i, value)
		}
	*/

	// 4
	/*
		nums := [10]int{}

		for i := 0; i < 10; i++ {
			fmt.Printf("Введите число %d: ", i+1)
			fmt.Scan(&nums[i])
		}

		for i := 0; i < len(nums); i++ {
			if nums[i] >= 0 {
				fmt.Printf("%d ", nums[i])
			}
		}
	*/

	// 5
	/*
		nums := [10]int{}

		for i := 0; i < 10; i++ {
			fmt.Printf("Введите число %d: ", i+1)
			fmt.Scan(&nums[i])
		}

		for i := 0; i < len(nums); i++ {
			if nums[i]%2 == 0 {
				fmt.Printf("%d ", nums[i])
			}
		}
	*/

	// 6
	/*
		nums := [6]int{}

		for i := 0; i < 6; i++ {
			fmt.Printf("Введите число %d: ", i+1)
			fmt.Scan(&nums[i])
		}

		for i := 0; i < len(nums); i++ {
			fmt.Printf("%d ", nums[i]*nums[i])
		}
	*/

	// 7
	/*
		nums := [8]int{}

		for i := 0; i < 8; i++ {
			fmt.Printf("Введите число %d: ", i+1)
			fmt.Scan(&nums[i])
		}

		for i := 0; i < len(nums); i++ {
			if i%2 == 0 {
				fmt.Printf("%d ", nums[i])
			}
		}
	*/

	// 8
	/*
		nums := [8]int{}

		for i := 0; i < 8; i++ {
			fmt.Printf("Введите число %d: ", i+1)
			fmt.Scan(&nums[i])
		}

		counter := 0
		for i := 0; i < len(nums); i++ {
			if nums[i] >= 0 {
				counter++
			}
		}
		fmt.Println(counter)
	*/

	// 9
	/*
		nums := [8]int{}

		for i := 0; i < 8; i++ {
			fmt.Printf("Введите число %d: ", i+1)
			fmt.Scan(&nums[i])
		}

		counter := 0
		for i := 0; i < len(nums); i++ {
			counter = counter + nums[i]
		}
		fmt.Println(counter)
	*/

	// 10
	/*
		nums := [6]int{}

		for i := 0; i < 6; i++ {
			fmt.Printf("Введите число %d: ", i+1)
			fmt.Scan(&nums[i])
		}

		counter := 1
		for i := 0; i < len(nums); i++ {
			if nums[i] != 0 {
				counter = counter * nums[i]
			}
		}
		fmt.Println(counter)
	*/

	// 11
	/*
		var k int
		nums := [10]int{}

		fmt.Printf("Введите делитель: ")
		fmt.Scan(&k)

		for i := 0; i < 10; i++ {
			fmt.Printf("Введите числа %d: ", i+1)
			fmt.Scan(&nums[i])
		}

		for i := 0; i < len(nums); i++ {
			if nums[i]%k == 0 {
				fmt.Println(nums[i])
			}
		}
	*/

	// 12
	/*
		nums := [10]int{}

		for i := 0; i < 10; i++ {
			fmt.Printf("Введите число %d: ", i+1)
			fmt.Scan(&nums[i])
		}

		value := nums[0]
		for i := 0; i < len(nums); i++ {
			if nums[i] > value {
				value = nums[i]
			}
		}
		fmt.Println(value)
	*/

	// 13
	/*
		nums := [10]int{}

		for i := 0; i < 10; i++ {
			fmt.Printf("Введите число %d: ", i+1)
			fmt.Scan(&nums[i])
		}

		value := nums[0]
		valueIndex := 0
		for i := 0; i < len(nums); i++ {
			if nums[i] < value {
				valueIndex = i
				value = nums[i]
			}
		}
		fmt.Printf("%d %d", valueIndex, value)
	*/

	// 14
	/*
		nums := [10]int{}

		for i := 0; i < 10; i++ {
			fmt.Printf("Введите число %d: ", i+1)
			fmt.Scan(&nums[i])
		}

		min := nums[0]
		max := nums[0]
		for i := 0; i < len(nums); i++ {
			if nums[i] < min {
				min = nums[i]
			}
			if nums[i] > max {
				max = nums[i]
			}
		}
		fmt.Printf("min %d; max %d", min, max)
	*/

	// 15
	/*
		var m int
		nums := [10]int{}

		fmt.Printf("Введите m: ")
		fmt.Scan(&m)

		for i := 0; i < 10; i++ {
			fmt.Printf("Введите числа %d: ", i+1)
			fmt.Scan(&nums[i])
		}

		isNumber := false
		for i := 0; i < len(nums); i++ {
			if nums[i] == m {
				fmt.Printf("Yes, index = %d\n", i)
				isNumber = true
			}
		}

		if isNumber == false {
			fmt.Println("No")
		}
	*/

	// Add
	var membersCount int = 0
	fmt.Print("Введите общее количество участников: ")
	_, err := fmt.Scan(&membersCount)
	if err != nil || membersCount <= 0 {
		log.Println("Ошибка ввода. Введите положительное число участников.")
		return
	}

	names := make([]string, membersCount)
	weights := make([]float64, membersCount)

	fmt.Println("Введите участников в формате 'имя вес':")
	for i := 0; i < membersCount; i++ {
		_, err = fmt.Scan(&names[i], &weights[i])
		if err != nil || weights[i] <= 0 {
			log.Printf("Ошибка ввода данных для участника #%d. Повторите попытку.\n", i+1)
			return
		}
	}

	lightWeights := []int{}
	averageWeights := []int{}
	heavyWeights := []int{}

	for i, weight := range weights {
		switch {
		case weight <= 64:
			lightWeights = append(lightWeights, i)
		case weight <= 78:
			averageWeights = append(averageWeights, i)
		default:
			heavyWeights = append(heavyWeights, i)
		}
	}

	if len(lightWeights) > 0 {
		fmt.Printf("Легкая весовая категория до 64 кг (вкл) - количество участников: %d\n", len(lightWeights))
		for _, value := range lightWeights {
			fmt.Printf("%s - %f\n", names[value], weights[value])
		}
	}

	if len(averageWeights) > 0 {
		fmt.Printf("Средняя весовая категория до 78 кг (вкл) - количество участников: %d\n", len(averageWeights))
		for _, value := range averageWeights {
			fmt.Printf("%s - %f\n", names[value], weights[value])
		}
	}

	if len(heavyWeights) > 0 {
		fmt.Printf("Тяжелая весовая категория от 78 кг - количество участников: %d\n", len(heavyWeights))
		for _, value := range heavyWeights {
			fmt.Printf("%s - %f\n", names[value], weights[value])
		}
	}

}

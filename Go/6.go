package main

import (
	"fmt"
	"log"
)

func main() {

	// 1
	/*
		var n int
		var m int
		fmt.Println("Введите n и m: ")
		_, err := fmt.Scan(&n, &m)
		if err != nil {
			log.Println(err)
			return
		}

		matrix := [4][4]int{{12, 5, 7, 6}, {4, 0, 2, 7}, {9, 1, 3, 2}, {10, -2, 4, 6}}
		fmt.Println(matrix[n][m])
	*/

	// 2
	/*
		var n int
		var m int
		fmt.Println("Введите n и m: ")
		_, err := fmt.Scan(&n, &m)
		if err != nil {
			log.Println(err)
			return
		}

		matrix := make([][]int, n)
		for i := range matrix {
			matrix[i] = make([]int, m)
		}

		fmt.Println("Введите данные в массив: ")

		for i, _ := range matrix {
			for ii, _ := range matrix[i] {
				fmt.Scan(&matrix[i][ii])
			}
		}

		fmt.Println("Созданная матрица: ")
		fmt.Println(matrix)

		sum := 0
		for i, _ := range matrix {
			for _, value := range matrix[i] {
				sum += value
			}
		}
		fmt.Printf("Сумма всех чисел в массиве: %v", sum)
	*/

	// 3
	/*
		var n int
		var m int
		fmt.Println("Введите n и m: ")
		_, err := fmt.Scan(&n, &m)
		if err != nil {
			log.Println(err)
			return
		}

		matrix := make([][]int, n)
		for i := range matrix {
			matrix[i] = make([]int, m)
		}

		fmt.Println("Введите данные в массив: ")

		for i, _ := range matrix {
			for ii, _ := range matrix[i] {
				fmt.Scan(&matrix[i][ii])
			}
		}

		fmt.Println("Созданная матрица: ")
		fmt.Println(matrix)

		min := matrix[0][0]
		for i, _ := range matrix {
			for _, value := range matrix[i] {
				if min > value {
					min = value
				}
			}
		}
		fmt.Printf("Минимальное число в матрице: %v", min)
	*/

	// 4
	/*
		var n int
		var m int
		fmt.Println("Введите n и m: ")
		_, err := fmt.Scan(&n, &m)
		if err != nil {
			log.Println(err)
			return
		}

		matrix := make([][]int, n)
		for i := range matrix {
			matrix[i] = make([]int, m)
		}

		fmt.Println("Введите данные в массив: ")

		for i, _ := range matrix {
			for ii, _ := range matrix[i] {
				fmt.Scan(&matrix[i][ii])
			}
		}

		fmt.Println("Созданная матрица: ")
		fmt.Println(matrix)

		max := matrix[0][0]
		maxIndex := [2]int{0, 0}
		for i, _ := range matrix {
			for ii, value := range matrix[i] {
				if max < value {
					maxIndex[0] = i
					maxIndex[1] = ii
					max = value
				}
			}
		}
		fmt.Printf("Максимальное число в матрице: %v. Его адрес: %v", max, maxIndex)
	*/

	// 5
	/*
		var n int
		var m int
		fmt.Println("Введите n и m: ")
		_, err := fmt.Scan(&n, &m)
		if err != nil {
			log.Println(err)
			return
		}

		matrix := make([][]int, n)
		for i := range matrix {
			matrix[i] = make([]int, m)
		}

		fmt.Println("Введите данные в массив: ")
		for i, _ := range matrix {
			for ii, _ := range matrix[i] {
				fmt.Scan(&matrix[i][ii])
			}
		}

		fmt.Println("Созданная матрица: ")
		fmt.Println(matrix)

		quantity := len(matrix) * len(matrix[0])
		sum := 0
		for i, _ := range matrix {
			for _, value := range matrix[i] {
				sum += value
			}
		}

		var average float64 = float64(sum) / float64(quantity)

		fmt.Printf("Среднее арифметическое всех элементов: %v\n", average)
	*/

	// 6
	/*
		var n int
		var m int
		fmt.Println("Введите n и m: ")
		_, err := fmt.Scan(&n, &m)
		if err != nil {
			log.Println(err)
			return
		}

		matrix := make([][]int, n)
		for i := range matrix {
			matrix[i] = make([]int, m)
		}

		fmt.Println("Введите данные в массив: ")
		for i, _ := range matrix {
			for ii, _ := range matrix[i] {
				fmt.Scan(&matrix[i][ii])
			}
		}

		fmt.Println("Созданная матрица: ")
		fmt.Println(matrix)

		num := 0
		fmt.Println("Введите число которое надо найти: ")
		fmt.Scan(&num)

		for i, _ := range matrix {
			for _, value := range matrix[i] {
				if value == num {
					fmt.Printf("YES")
					break
				}
			}
		}
	*/

	// 7
	/*
		var n int
		var m int
		fmt.Println("Введите n и m: ")
		_, err := fmt.Scan(&n, &m)
		if err != nil {
			log.Println(err)
			return
		}

		matrix := make([][]int, n)
		for i := range matrix {
			matrix[i] = make([]int, m)
		}

		fmt.Println("Введите данные в массив: ")
		for i, _ := range matrix {
			for ii, _ := range matrix[i] {
				fmt.Scan(&matrix[i][ii])
			}
		}

		fmt.Println("Созданная матрица: ")
		fmt.Println(matrix)

		evenNums := 0
		for i, _ := range matrix {
			for _, value := range matrix[i] {
				if value%2 == 0 {
					evenNums += 1
				}
			}
		}

		fmt.Printf("Количество четных элементов: %v\n", evenNums)
	*/

	// 9
	/*
		var n int
		var m int
		fmt.Println("Введите n и m: ")
		_, err := fmt.Scan(&n, &m)
		if err != nil {
			log.Println(err)
			return
		}

		matrix := make([][]int, n)
		for i := range matrix {
			matrix[i] = make([]int, m)
		}

		fmt.Println("Введите данные в массив: ")
		for i, _ := range matrix {
			for ii, _ := range matrix[i] {
				fmt.Scan(&matrix[i][ii])
			}
		}

		fmt.Println("Созданная матрица: ")
		fmt.Println(matrix)

		matrix[0], matrix[len(matrix)-1] = matrix[len(matrix)-1], matrix[0]

		fmt.Println("Обработанная матрица: ")
		fmt.Println(matrix)
	*/

	// 10
	/*
		var n int
		var m int
		fmt.Println("Введите n и m: ")
		_, err := fmt.Scan(&n, &m)
		if err != nil {
			log.Println(err)
			return
		}

		matrix := make([][]int, n)
		for i := range matrix {
			matrix[i] = make([]int, m)
		}

		fmt.Println("Введите данные в массив: ")
		for i, _ := range matrix {
			for ii, _ := range matrix[i] {
				fmt.Scan(&matrix[i][ii])
			}
		}

		fmt.Println("Созданная матрица: ")
		fmt.Println(matrix)

		for i, _ := range matrix {
			matrix[i][0], matrix[i][len(matrix[i])-1] = matrix[i][len(matrix[i])-1], matrix[i][0]
		}

		fmt.Println("Обработанная матрица: ")
		fmt.Println(matrix)
	*/

	// 11
	/*
		var n int
		var m int
		fmt.Println("Введите n и m: ")
		_, err := fmt.Scan(&n, &m)
		if err != nil {
			log.Println(err)
			return
		}

		matrix := make([][]int, n)
		for i := range matrix {
			matrix[i] = make([]int, m)
		}

		fmt.Println("Введите данные в массив: ")

		for i, _ := range matrix {
			for ii, _ := range matrix[i] {
				fmt.Scan(&matrix[i][ii])
			}
		}

		fmt.Println("Созданная матрица: ")
		fmt.Println(matrix)

		for i, _ := range matrix {
			for ii, _ := range matrix[i] {
				if i != ii && matrix[i][ii] != matrix[ii][i] {
					//fmt.Printf("Адрес: %v %v. Значение: %v \n", i, ii, value)
					fmt.Println("NO")
					return
				}
			}
		}
	*/

	// 12
	/*
		var n int
		var m int
		fmt.Println("Введите n и m: ")
		_, err := fmt.Scan(&n, &m)
		if err != nil {
			log.Println(err)
			return
		}

		matrix := make([][]int, n)
		for i := range matrix {
			matrix[i] = make([]int, m)
		}

		fmt.Println("Введите данные в массив: ")

		for i, _ := range matrix {
			for ii, _ := range matrix[i] {
				fmt.Scan(&matrix[i][ii])
			}
		}

		fmt.Println("Созданная матрица: ")
		fmt.Println(matrix)

		maxs := []int{}

		for _, row := range matrix {
			max := row[0]
			for _, value := range row {
				if value > max {
					max = value
				}
			}
			maxs = append(maxs, max)
		}

		fmt.Printf("Максимальные значения в строках %v \n", maxs)
	*/

	// 13
	/*
		var n int
		var m int
		fmt.Println("Введите n и m: ")
		_, err := fmt.Scan(&n, &m)
		if err != nil {
			log.Println(err)
			return
		}

		matrix := make([][]int, n)
		for i := range matrix {
			matrix[i] = make([]int, m)
		}

		fmt.Println("Введите данные в массив: ")

		for i, _ := range matrix {
			for ii, _ := range matrix[i] {
				fmt.Scan(&matrix[i][ii])
			}
		}

		fmt.Println("Созданная матрица: ")
		fmt.Println(matrix)

		mins := []int{}

		for _, row := range matrix {
			min := row[0]
			for _, value := range row {
				if value < min {
					min = value
				}
			}
			mins = append(mins, min)
		}

		fmt.Printf("Минимальное значения в строках %v \n", mins)
	*/

	// 14
	var n int
	var m int
	fmt.Println("Введите n и m: ")
	_, err := fmt.Scan(&n, &m)
	if err != nil {
		log.Println(err)
		return
	}

	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, m)
	}

	fmt.Println("Введите данные в массив: ")

	for i, _ := range matrix {
		for ii, _ := range matrix[i] {
			fmt.Scan(&matrix[i][ii])
		}
	}

	fmt.Println("Созданная матрица: ")
	fmt.Println(matrix)

	sums := []int{}

	for _, row := range matrix {
		sum := 0
		for _, value := range row {
			sum += value
		}
		sums = append(sums, sum)
	}

	fmt.Printf("Суммы строк %v \n", sums)
}

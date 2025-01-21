package main

import "fmt"

func main() {
	// 1
	/*
		func(numbers ...int) int {
			if len(numbers) == 0 {
				fmt.Println("🖕🏿")
				return 0
			}
			count := 0
			sum := 0
			for _, value := range numbers {
				count++
				sum += value
			}
			fmt.Println(sum / count)
			return sum / count
		}(10, 1000)
	*/

	// 2
	/*
		func(numbers ...int) int {
			if len(numbers) == 0 {
				fmt.Println("🖕🏿")
				return 0
			}
			max := numbers[0]
			for _, value := range numbers {
				if value > max {
					max = value
				}
			}
			fmt.Println(max)
			return max
		}(10, 1000, 1)
	*/

	// 3
	/*
		func(numbers ...int) int {
			if len(numbers) == 0 {
				fmt.Println("🖕🏿")
				return 0
			}
			sum := 0
			for _, value := range numbers {
				sum += value
			}
			fmt.Println(sum)
			return sum
		}()
	*/

	// 4
	/*
		fmt.Println("Введите количество элементов: ")
		var inputCount int
		fmt.Scan(&inputCount)

		fmt.Println("Введите элементы: ")
		input := make([]int, inputCount)
		for i := 0; i < inputCount; i++ {
			fmt.Scan(&input[i])
		}

		func(numbers ...int) string {
			if len(numbers) == 0 {
				fmt.Println("🖕🏿")
				return "error"
			}
			for _, value := range numbers {
				if value%2 == 0 {
					fmt.Println("This list contains an even number.")
					return "This list contains an even number."
				}
			}
			fmt.Println("Only odd.")
			return "Only odd."
		}(input...)
	*/

	// 5
	/*
		fmt.Println("Введите количество элементов: ")
		var inputCount int
		fmt.Scan(&inputCount)

		fmt.Println("Введите элементы: ")
		input := make([]int, inputCount)
		for i := 0; i < inputCount; i++ {
			fmt.Scan(&input[i])
		}

		func(numbers ...int) int {
			if len(numbers) == 0 {
				fmt.Println("🖕🏿")
				return 0
			}
			sum := 1
			for _, value := range numbers {
				sum *= value
			}
			fmt.Println(sum)
			return sum
		}(input...)
	*/

	// 6
	/*
		fmt.Println("Введите количество элементов: ")
		var inputCount int
		fmt.Scan(&inputCount)

		fmt.Println("Введите элементы: ")
		input := make([]int, inputCount)
		for i := 0; i < inputCount; i++ {
			fmt.Scan(&input[i])
		}

		func(numbers ...int) string {
			if len(numbers) == 0 {
				fmt.Println("🖕🏿")
				return "error"
			}
			for _, value := range numbers {
				if value < 0 {
					fmt.Println("Some numbers are not positive.")
					return "Some numbers are not positive."
				}
			}
			fmt.Println("Only positive.")
			return "Only positive."
		}(input...)
	*/

	// 7
	/*
		fmt.Println("Введите количество элементов: ")
		var inputCount int
		fmt.Scan(&inputCount)

		fmt.Println("Введите элементы: ")
		input := make([]int, inputCount)
		for i := 0; i < inputCount; i++ {
			fmt.Scan(&input[i])
		}

		func(numbers ...int) int {
			if len(numbers) == 0 {
				fmt.Println("🖕🏿")
				return 0
			}
			sum := 0
			for _, value := range numbers {
				if value > 0 {
					sum += value
				}
			}
			fmt.Println(sum)
			return sum
		}(input...)
	*/

	// 8
	/*
		fmt.Println("Введите количество элементов: ")
		var inputCount int
		fmt.Scan(&inputCount)

		fmt.Println("Введите элементы: ")
		input := make([]int, inputCount)
		for i := 0; i < inputCount; i++ {
			fmt.Scan(&input[i])
		}

		func(numbers ...int) string {
			if len(numbers) == 0 {
				fmt.Println("🖕🏿")
				return "error"
			}
			for _, value := range numbers {
				if value%2 != 0 {
					fmt.Println("Odd.")
					return "Odd."
				}
			}
			fmt.Println("All numbers are even.")
			return "All numbers are even."
		}(input...)
	*/

	// 9
	/*
		fmt.Println("Введите количество элементов: ")
		var inputCount int
		fmt.Scan(&inputCount)

		fmt.Println("Введите элементы: ")
		input := make([]string, inputCount)
		for i := 0; i < inputCount; i++ {
			fmt.Scan(&input[i])
		}

		func(inputs ...string) string {
			check := len(inputs[0])
			for _, value := range inputs {
				if len(value) != check {
					fmt.Println("🖕🏿")
					return "🖕🏿"
				}
			}
			fmt.Println("All strings have the same length.")
			return "All strings have the same length."
		}(input...)
	*/

	// 10
	fmt.Println("Введите количество элементов: ")
	var inputCount int
	fmt.Scan(&inputCount)

	fmt.Println("Введите элементы: ")
	input := make([]string, inputCount)
	for i := 0; i < inputCount; i++ {
		fmt.Scan(&input[i])
	}

	func(inputs ...string) string {
		check := len(inputs[0])
		for _, value := range inputs {
			fmt.Println(value[0])
			if len(value) != check {
				fmt.Println("🖕🏿")
				return "🖕🏿"
			}
		}
		fmt.Println("All strings have the same length.")
		return "All strings have the same length."
	}(input...)
}

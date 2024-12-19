package main

func main() {

	// Lesson Slices

	// Вместимость
	/*
		names := make([]string, 0, 5)
		namesAdd := []string{"12", "2323"}

		fmt.Println(len(names))
		fmt.Println(cap(names))

		// Добавление
		names = append(names, "alex")
		names = append(names, namesAdd...) // variadic
	*/

	// 1
	/*
		var n int
		fmt.Println("Введите N: ")

		_, err := fmt.Scan(&n)
		if err != nil {
			log.Println(err)
			return
		}

		myArray := make([]int, n, n)

		fmt.Println("Введите элементы: ")
		for i := 0; i < n; i++ {
			fmt.Scan(&myArray[i])
		}

		fmt.Println(myArray[3])
	*/

	// 2
	/*
		var n int
		fmt.Println("Введите N: ")

		_, err := fmt.Scan(&n)
		if err != nil {
			log.Println(err)
			return
		}

		myArray := make([]int, n)
		fmt.Println("Введите элементы: ")
		for i := 0; i < n; i++ {
			fmt.Scan(&myArray[i])
		}

		max := myArray[0]

		for _, value := range myArray {
			if value > max {
				max = value
			}
		}

		fmt.Println(max)
	*/

	// 3
	/*
		var n int
		fmt.Println("Введите N: ")

		_, err := fmt.Scan(&n)
		if err != nil {
			log.Println(err)
			return
		}

		myArray := make([]int, n, n)
		fmt.Println("Введите элементы: ")
		for i := 0; i < n; i++ {
			fmt.Scan(&myArray[i])
		}

		num := 1
		for _, value := range myArray {
			if value != 0 {
				num = num * value
			}
		}

		fmt.Println(num)
	*/

	// 4
	/*
		var n int
		fmt.Println("Введите N: ")

		_, err := fmt.Scan(&n)
		if err != nil {
			log.Println(err)
			return
		}

		myArray := make([]int, n, n)
		fmt.Println("Введите элементы: ")
		for i := 0; i < n; i++ {
			fmt.Scan(&myArray[i])
		}

		min := myArray[0] * myArray[0]

		for _, value := range myArray {
			if value*value < min {
				min = value * value
			}
		}

		fmt.Println(min)
	*/

	// 5
	/*
		var n int
		fmt.Println("Введите N: ")

		_, err := fmt.Scan(&n)
		if err != nil {
			log.Println(err)
			return
		}

		myArray := make([]int, n)
		fmt.Println("Введите элементы: ")
		for i := 0; i < n; i++ {
			fmt.Scan(&myArray[i])
		}

		max := myArray[0]
		maxIndex := 0

		for i, value := range myArray {
			if value > max {
				max = value
				maxIndex = i
			}
		}

		fmt.Printf("%d %d\n", max, maxIndex)
	*/

	// 6
	/*
		var n int
		fmt.Println("Введите N: ")

		_, err := fmt.Scan(&n)
		if err != nil {
			log.Println(err)
			return
		}

		myArray := make([]int, n)
		fmt.Println("Введите элементы: ")
		for i := 0; i < n; i++ {
			fmt.Scan(&myArray[i])
		}

		max := myArray[0]
		min := myArray[0]

		for _, value := range myArray {
			if value > max {
				max = value
			}
			if value < min {
				min = value
			}
		}

		fmt.Printf("max %d, min %d", max, min)
	*/

	// 7
	/*
		var n int
		fmt.Println("Введите N: ")

		_, err := fmt.Scan(&n)
		if err != nil {
			log.Println(err)
			return
		}

		myArray := make([]int, n)
		fmt.Println("Введите элементы: ")
		for i := 0; i < n; i++ {
			fmt.Scan(&myArray[i])
		}

		max := myArray[0]
		min := myArray[0]

		for _, value := range myArray {
			if value > max {
				max = value
			}
			if value < min {
				min = value
			}
		}

		fmt.Printf("max %d, min %d, delta %d", max, min, max-min)
	*/

	// 8
	/*
		var n int
		fmt.Println("Введите N: ")
		_, err := fmt.Scan(&n)
		if err != nil {
			log.Println(err)
			return
		}

		myArray := make([]int, n)
		fmt.Println("Введите элементы: ")
		for i := 0; i < n; i++ {
			fmt.Scan(&myArray[i])
		}

		var k int
		fmt.Println("Введите k: ")
		_, err = fmt.Scan(&k)
		if err != nil {
			log.Println(err)
			return
		}

		for _, value := range myArray {
			if value%k == 0 {
				fmt.Printf("%d ", value)
			}
		}
	*/

	// 9
	/*
		var n int
		fmt.Println("Введите N: ")
		_, err := fmt.Scan(&n)
		if err != nil {
			log.Println(err)
			return
		}

		myArray := make([]int, n)
		fmt.Println("Введите элементы: ")
		for i := 0; i < n; i++ {
			fmt.Scan(&myArray[i])
		}

		for i := len(myArray) - 1; i >= 0; i-- {
			fmt.Printf("%d ", myArray[i])
		}
	*/

	// 10
	/*
		var n int
		fmt.Println("Введите N: ")
		_, err := fmt.Scan(&n)
		if err != nil {
			log.Println(err)
			return
		}

		myArray := make([]int, n)
		fmt.Println("Введите элементы: ")
		for i := 0; i < n; i++ {
			fmt.Scan(&myArray[i])
		}

		oddArray := []int{}
		for i, value := range myArray {
			if i%2 == 0 {
				oddArray = append(oddArray, value)
			}
		}

		fmt.Print(oddArray)
	*/

	// 11
	/*
		var n int
		fmt.Println("Введите N: ")
		_, err := fmt.Scan(&n)
		if err != nil {
			log.Println(err)
			return
		}

		myArray := make([]int, n)
		fmt.Println("Введите элементы: ")
		for i := 0; i < n; i++ {
			fmt.Scan(&myArray[i])
		}

		isFirstZero := false
		firstZero := 0
		isSecondZero := false
		secondZero := 0

		for i, value := range myArray {
			if value == 0 {
				if isFirstZero == false {
					firstZero = i
					isFirstZero = true
				} else if isSecondZero == false {
					secondZero = i
					isSecondZero = true
				}
			}
		}

		if isFirstZero == false {
			fmt.Printf("В массиве нет нулей")
		} else if isSecondZero == false {
			fmt.Printf("В массиве нет второго нуля")
		} else {
			sum := 0
			for i := firstZero + 1; i < secondZero; i++ {
				sum = sum + myArray[i]
			}
			fmt.Println(sum)
		}
	*/

	// 12
	/*
		var n int
		fmt.Println("Введите N: ")
		_, err := fmt.Scan(&n)
		if err != nil {
			log.Println(err)
			return
		}

		myArray := make([]int, n)
		fmt.Println("Введите элементы: ")
		for i := 0; i < n; i++ {
			fmt.Scan(&myArray[i])
		}

		isSequence := false

		for i, value := range myArray {
			if i != 0 {
				if value > 0 && myArray[i-1] > 0 {
					isSequence = true
				} else if value < 0 && myArray[i-1] < 0 {
					isSequence = true
				}
			}
		}

		if isSequence {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}
	*/

	// 13
	/*
		initialUsers := [8]string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}
		myArray := initialUsers[0:3]
		fmt.Print(myArray)
	*/

	// 14
	/*
		var n int
		fmt.Println("Введите N: ")
		_, err := fmt.Scan(&n)
		if err != nil {
			log.Println(err)
			return
		}

		myArray := make([]int, n)
		fmt.Println("Введите элементы: ")
		for i := 0; i < n; i++ {
			fmt.Scan(&myArray[i])
		}

		sum := 1
		for i, value := range myArray {
			if i%2 == 0 && value != 0 {
				sum = sum * value
			}
		}

		fmt.Println(sum)
	*/

	// 15
	/*
		var n int
		fmt.Println("Введите N: ")
		_, err := fmt.Scan(&n)
		if err != nil {
			log.Println(err)
			return
		}

		myArray := make([]string, n)
		fmt.Println("Введите элементы: ")
		for i := 0; i < n; i++ {
			fmt.Scan(&myArray[i])
		}

		myArray = append(myArray, "Arman", "Ilyas", "Tore")

		for i := len(myArray) - 1; i >= 0; i-- {
			fmt.Printf("%s ", myArray[i])
		}
	*/

}

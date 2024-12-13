package main

func main() {
	// 1.
	/*
		for i := 1; i <= 10; i++ {
			fmt.Printf("I like\n")
		}
	*/

	// 2.
	/*
		var num int
		fmt.Println("Введите число: ")

		_, err := fmt.Scan(&num)
		if err != nil {
			log.Println(err)
			return
		}

		for i := 1; i <= 10; i++ {
			fmt.Printf("%d\n", num)
		}
	*/

	// 3.
	/*
		var num int
		fmt.Println("Введите число: ")

		_, err := fmt.Scan(&num)
		if err != nil {
			log.Println(err)
			return
		}

		for i := 1; i <= num; i++ {
			fmt.Printf("%d\n", i)
		}
	*/

	// 4.
	/*
		var str string
		var num int

		fmt.Println("Введите число: ")
		_, err := fmt.Scan(&num)
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println("Введите строку: ")
		fmt.Scan(&str)

		for i := 1; i <= num; i++ {
			fmt.Printf("%s\n", str)
		}
	*/

	// 5.
	/*
		var n int
		var m int

		fmt.Println("Введите число n: ")
		_, err := fmt.Scan(&n)
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println("Введите число m: ")
		_, err = fmt.Scan(&m)
		if err != nil {
			log.Println(err)
			return
		}

		for i := n; i <= m; i++ {
			fmt.Printf("%d\n", i)
		}
	*/

	// 6.
	/*
		for i := 1; i <= 10; i++ {
			fmt.Println(i * i)
		}
	*/

	// 7.
	/*
		var num int
		fmt.Println("Введите число: ")

		_, err := fmt.Scan(&num)
		if err != nil {
			log.Println(err)
			return
		}

		for i := 0; i <= num; i++ {
			if i%2 != 0 {
				continue
			}
			fmt.Printf("%d\n", i)
		}
	*/

	// 8.
	/*
		var num int
		sum := 1
		fmt.Println("Введите число: ")

		_, err := fmt.Scan(&num)
		if err != nil {
			log.Println(err)
			return
		}

		for i := 1; i <= num; i++ {
			sum = sum * i
		}

		fmt.Printf("%d\n", sum)
	*/

	// 9.
	/*
		var a int
		var b int
		sum := 0

		fmt.Println("Введите число a: ")
		_, err := fmt.Scan(&a)
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println("Введите число b: ")
		_, err = fmt.Scan(&b)
		if err != nil {
			log.Println(err)
			return
		}

		if a >= 100 || b >= 100 {
			log.Println("Число больше 100")
			return
		}

		switch {
		case a < b:
			for a <= b {
				sum = sum + a
				a++
			}
			fmt.Print(sum)
		case a > b:
			for a >= b {
				sum = sum - a
				a--
			}
			fmt.Print(sum)
		default:
			fmt.Println("Числа равны")
		}
	*/

	// 10
	/*
		var n, c, d int

		fmt.Println("Введите число n: ")
		_, err := fmt.Scan(&n)
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println("Введите число c: ")
		_, err = fmt.Scan(&c)
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println("Введите число d: ")
		_, err = fmt.Scan(&d)
		if err != nil {
			log.Println(err)
			return
		}

		for i := 1; i <= n; i++ {
			if i%c == 0 && i%d != 0 {
				fmt.Printf("%d\n", i)
				break
			}
		}
	*/

	// 11
	/*
		var n, m int

		fmt.Println("Введите число n: ")
		_, err := fmt.Scan(&n)
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println("Введите число m: ")
		_, err = fmt.Scan(&m)
		if err != nil {
			log.Println(err)
			return
		}

		for i := n; i <= m; i++ {
			if i%2 != 0 {
				fmt.Printf("%d\n", i)
			}
		}
	*/

	// 12
	/*
		var n, m int

		fmt.Println("Введите число n: ")
		_, err := fmt.Scan(&n)
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println("Введите число m: ")
		_, err = fmt.Scan(&m)
		if err != nil {
			log.Println(err)
			return
		}

		for i := n; i <= m; i++ {
			if i%3 == 0 {
				fmt.Printf("%d\n", i)
			}
		}
	*/

	// 13
	/*
		var n int

		fmt.Println("Введите число: ")
		_, err := fmt.Scan(&n)
		if err != nil {
			log.Println(err)
			return
		}

		for i := 1; i <= n; i++ {
			if n%i == 0 {
				fmt.Printf("%d ", i)
			}
		}
	*/

	// 14
	/*
		var n int

		fmt.Println("Введите число: ")
		_, err := fmt.Scan(&n)
		if err != nil {
			log.Println(err)
			return
		}

		for i := 1; i <= n; i++ {
			fmt.Printf("%d %d\n", i, i*i)
		}
	*/

	// 15
	/*
		var a, b, c, d int

		fmt.Println("Введите числа: ")
		fmt.Scan(&a, &b, &c, &d)

		for i := a; i <= b; i++ {
			if i%d == c {
				fmt.Printf("%d ", i)
			}
		}
	*/

	// Add 1
	/*
		var num int

		for num < 18 {
			fmt.Println("Сколько тебе лет?")
			fmt.Scan(&num)
		}

		fmt.Println("Приятного просмотра")
	*/

	// Add 2
	/*
		var num int
		var wrongCount int

		for num < 18 {
			if wrongCount >= 3 {
				fmt.Println("Слишком много попыток")
				break
			}
			fmt.Println("Сколько тебе лет?")
			fmt.Scan(&num)

			if num < 18 {
				wrongCount++
			}
		}

		if wrongCount < 3 && num >= 18 {
			fmt.Println("Приятного просмотра")
		}
	*/
}

package main

func main() {
	// 1. На ввод подается целое число. Если число положительное - вывести сообщение "Positive", если число отрицательное - "Negative". Если подается ноль - вывести сообщение "Zero". Выводить сообщение без кавычек.
	/*
		var num int
		fmt.Println("Введите число: ")
		fmt.Scan(&num)

		if num >= 1 {
			fmt.Println("Positive")
		} else if num <= -1 {
			fmt.Println("Negative")
		} else {
			fmt.Println("Zero")
		}
	*/

	// 2. Составьте программу, которая принимает на вход трехзначное число. По данному трехзначному числу определите, все ли его цифры различны.
	/*
		var num int
		fmt.Println("Введите число: ")
		fmt.Scan(&num)

		if num >= 100 && num <= 999 {

			num1 := (num / 100) % 10
			num2 := (num / 10) % 10
			num3 := num % 10

			if num1 != num2 && num1 != num3 && num2 != num3 {
				fmt.Print("YES")
			} else {
				fmt.Print("NO")
			}

		} else {
			return
		}
	*/

	// 3. Напишите программу, которая принимает на вход натуральное число. Необходимо при положительном числе вывести его первую цифру, иначе для отрицательного числа ничего выводить не нужно.
	/*
		var num int
		fmt.Println("Введите число: ")

		_, err := fmt.Scan(&num)
		if err != nil {
			log.Println(err)
			return
		}

		if num <= 0 {
			return
		}

		for num > 10 {
			num = num / 10
		}

		fmt.Print(num)
	*/

	// 4. Напишите программу, которая определяет является ли билет счастливым. Счастливым считается билет, в шестизначном номере которого сумма первых трёх цифр совпадает с суммой трёх последних.
	/*
		var num int
		fmt.Println("Введите число: ")

		_, err := fmt.Scan(&num)
		if err != nil {
			log.Println(err)
			return
		}

		if num >= 100000 && num <= 999999 {
			firstNum := num / 1000
			secondNum := num - (firstNum * 1000)

			firstSum := ((firstNum / 100) % 10) + ((firstNum / 10) % 10) + (firstNum % 10)
			secondSum := ((secondNum / 100) % 10) + ((secondNum / 10) % 10) + (secondNum % 10)
			if firstSum == secondSum {
				fmt.Print("YES")
			} else {
				fmt.Print("NO")
			}
		}
	*/

	// 5. Напишите программу, которая принимает год. Требуется определить, является ли данный год високосным, напомним: Год является високосным, если он соответствует хотя бы одному из нижеперечисленных условий: кратен 400; кратен 4, но не кратен 100.
	/*
		var year int
		fmt.Println("Введите год: ")

		_, err := fmt.Scan(&year)
		if err != nil {
			log.Println(err)
			return
		}

		if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
			fmt.Print("YES")
		} else {
			fmt.Print("NO")
		}
	*/

	// 6. Напишите программу, которая проверяет, делится ли значение введенное с консоли на 17. Если значение делится, нужно DIVISIBLE BY 17, иначе ничего выводить не нужно.
	/*
		var num int
		fmt.Println("Введите число: ")

		_, err := fmt.Scan(&num)
		if err != nil {
			log.Println(err)
			return
		}

		if num%17 == 0 {
			fmt.Print("DIVISIBLE BY 17")
		}
	*/

	// 7. Напишите программу, принимающую два числа. Если первое число меньше второго, то вывести "<", если первое число больше второго - ">", иначе вывести "=".
	/*
		var num1, num2 int

		fmt.Println("Введите числа: ")
		fmt.Scan(&num1, &num2)

		switch {
		case num1 > num2:
			fmt.Print(">")
		case num1 < num2:
			fmt.Print("<")
		default:
			fmt.Print("=")
		}
	*/

	// 8. Напишите программу, которая принимает 3 числа. Если произведение первых двух чисел равно третьему числу, то нужно вывести "YES", иначе вывести "NO".
	/*
		var num1, num2, num3 int

		fmt.Println("Введите числа: ")
		fmt.Scan(&num1, &num2, &num3)

		if num1*num2 == num3 {
			fmt.Print("YES")
		} else {
			fmt.Print("NO")
		}
	*/

	// 9. Напишите программу, которая по заданному номеру месяца в году определяет время года. Выведите для летних месяцев значение "Summer", для зимних – "Winter", для весенних – "Spring", для осенних – "Autumn". Если число не соответствует возможному значению месяца, то в этом случае следует вывести "Error".
	/*
		var num int

		fmt.Println("Введите месяц: ")
		fmt.Scan(&num)

		switch num {
		case 1, 2, 12:
			fmt.Println("Winter")
		case 3, 4, 5:
			fmt.Println("Spring")
		case 6, 7, 8:
			fmt.Println("Summer")
		case 9, 10, 11:
			fmt.Println("Autumm")
		default:
			fmt.Println("Error")
		}
	*/

	// 10. Напишите программу, которая проверяет число на четность/нечетность. Для четного числа необходимо вывести "EVEN", иначе вывести "ODD".
	/*
		var num int
		fmt.Println("Введите число: ")

		_, err := fmt.Scan(&num)
		if err != nil {
			log.Println(err)
			return
		}

		if num%2 == 0 {
			fmt.Print("EVEN")
		} else {
			fmt.Print("ODD")
		}
	*/

	// 11. Напишите программу, которая принимает два целых числа. Если их значения не равны, то выводить в консоль их сумму, иначе просто 0.
	/*
		var num1, num2 int

		fmt.Println("Введите числа: ")
		fmt.Scan(&num1, &num2)

		if num1 != num2 {
			fmt.Print(num1 + num2)
		} else {
			fmt.Print(0)
		}
	*/

	// 12. Напишите программу, где Арман вводит в каком классе он учится, а программа должна определить и выдать результат прописью как показано в примерах ниже. Помните, что с 1-го по 4-ый включительно классы - это Beginners School, с 5-го по 9-ый включительно - это Middle School и c 10-го по 11-ый включительно - это High School.
	/*
		var num int
		fmt.Println("Введите число: ")

		_, err := fmt.Scan(&num)
		if err != nil {
			log.Println(err)
			return
		}

		switch {
		case num <= 0, num >= 12:
			fmt.Println("Error")
		case num <= 4:
			fmt.Println("Beginners School")
		case num <= 9:
			fmt.Println("Middle School")
		case num <= 11:
			fmt.Println("High School")
		}
	*/

	// 13. Напишите программу, которая принимает число. Если оно меньше или равно 2019, то вывести NO, если больше или равно 2021, то вывести YES, если равно 2020, то вывести ERROR.
	/*
		var num int
		fmt.Println("Введите число: ")

		_, err := fmt.Scan(&num)
		if err != nil {
			log.Println(err)
			return
		}

		if num <= 2019 {
			fmt.Println("NO")
		} else if num >= 2021 {
			fmt.Println("YES")
		} else {
			fmt.Println("ERROR")
		}
	*/

	// 14. Напишите программу, которая принимает три числа. Если все числа одинаковые, то выведите YES, иначе NO.
	/*
		var num1, num2, num3 int

		fmt.Println("Введите числа: ")
		fmt.Scan(&num1, &num2, &num3)

		if num1 == num2 && num1 == num3 && num2 == num3 {
			fmt.Print("YES")
		} else {
			fmt.Print("NO")
		}
	*/

	// 15. Создайте программу, которая выводит в консоль YES, если введенное целое число является положительным и четным, иначе NO.
	/*
		var num int
		fmt.Println("Введите число: ")

		_, err := fmt.Scan(&num)
		if err != nil {
			log.Println(err)
			return
		}

		if num%2 == 0 && num > 0 {
			fmt.Print("YES")
		} else {
			fmt.Print("NO")
		}
	*/
}

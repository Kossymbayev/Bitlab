package main

func main() {
	//fmt.Println("Hello, World")

	// 1. Объявите переменную x типа int и присвойте ей значение 42. Выведите значение переменной x.
	/*
		x := 42
		fmt.Println(x)
	*/

	// 2. Напишите программу в которой вы создаете переменную name и присвойте ей значение "George Washington". Выведите значение переменной на экран c дополнительной информацией как на примере.
	/*
		name := "George Washington"
		fmt.Printf("The first president of USA is %s", name)
	*/

	// 3. Создайте программу, в которой будут переменные country значением "Great Britain" и capital со значением "London". Выведите известную для многих фразу как в примере.
	/*
		country := "Great Britain"
		capital := "London"
		fmt.Printf("%s is the capital of %s!", capital, country)
	*/

	// 4. Создайте программу с переменными: brand = "Apple", product = "IPhone", model = "13 Pro Max", price = "1500$", которая будет выводить в консоль полное описание продукта как в примере ниже.
	/*
		brand := "Apple"
		product := "IPhone"
		model := "13 Pro Max"
		price := "1500$"
		fmt.Printf("Brand: %s \nProduct: %s \nModel: %s \nPrice: %s", brand, product, model, price)
	*/

	// 5. Создайте переменную number типа int32 и присвойте ей значение 5. Выведите куб данного числа на экран.
	/*
		var number float64 = 5
		cube := math.Pow(number, 3)
		fmt.Println(cube)
	*/

	// 6. Напишите программу, которая последовательно делает следующие операции с введенным числом: 1. Число умножается на 2; 2. Затем к числу прибавляется 100. После этого должен быть вывод получившегося числа на экран.
	/*
		var number int
		fmt.Scan(&number)

		number *= 2
		number += 100

		fmt.Println(number)
	*/

	// 7. Аскар торопился в школу и неправильно написал программу, которая сначала находит квадраты двух чисел, а затем их суммирует. Напишите для него правильную программу.
	/*
		var num1, num2 int
		fmt.Scan(&num1)

		fmt.Scan(&num2)

		square1 := num1 * num1
		square2 := num2 * num2

		sum := square1 + square2

		fmt.Println(sum)
	*/

	// 8. Дано натуральное число, выведите его последнюю цифру.
	/*
		number := 123
		lastDigit := number % 10
		fmt.Println(lastDigit)
	*/

	// 9. Дано неотрицательное целое число. Найдите число десятков (то есть вторую цифру справа).
	/*
		number := 2010
		tens := (number / 10) % 10
		fmt.Println(tens)
	*/

	// 10. Арман получил 80.5 баллов (examOne float64) на первом экзамене и 55.7 (examTwo float64) баллов на втором. Напишите программу которая вычислит средний балл Армана и выведет на экран.
	/*
		var examOne float64 = 80.5
		var examTwo float64 = 55.7
		var result float64 = (examOne + examTwo) / 2
		fmt.Println(result)
	*/

	// 11. Напишите программ в которой вводится название любой страны (country string) и название любого города (city string). Далее выведите на экран введенные данные как в примере.
	/*
		var country string
		fmt.Scan(&country)

		var city string
		fmt.Scan(&city)

		fmt.Printf("Country: %s \nCity: %s", country, city)
	*/

	// 12. Напишите программу, которая принимает с консоли значение в долларах и выводит в тенге. (Курс: 1 доллар = 448 тенге)
	/*
		var tenge int
		fmt.Scan(&tenge)

		dollar := 448

		result := tenge * dollar
		fmt.Printf("%d, но он уже больше плак плак", result)
	*/

	// 13. Напишите программу, которая получает на вход две строки firstName и lastName. Выведите значения переменных firstName и lastName через пробел.
	/*
		var firstName string
		fmt.Scanln(&firstName)

		var lastName string
		fmt.Scan(&lastName)

		fmt.Printf("%s %s\n", firstName, lastName)
	*/

	// 14. Напишите программу, которая принимает данные о компании (companyName string, profit int, month float64) и выведет в консоль наименование компании и средний доход в месяц.
	/*
		var companyName string
		fmt.Scanln(&companyName)

		var profit int
		fmt.Scan(&profit)

		var month float64
		fmt.Scan(&month)

		averageProfit := float64(profit) / month

		fmt.Printf("%s\n%.2f\n", companyName, averageProfit)
	*/

	// 15. Создайте программу, которая рассчитает полную сумму чека, если цена блюд заказанные юзером soup = 850, salad = 530, tea = 245 (все переменные типа int) и выведет в консоль как в примере.
	/*
		soup := 850
		salad := 530
		tea := 245

		total := soup + salad + tea

		fmt.Println("TO PAY:")
		fmt.Printf("Soup - %d\n", soup)
		fmt.Printf("Salad - %d\n", salad)
		fmt.Printf("Tea - %d\n", tea)
		fmt.Println("- - -")
		fmt.Printf("Total - %d\n", total)
	*/
}

package main

import "fmt"

type Number int

func (n Number) IsSimple() string {
	if n == 2 || n == 3 || n == 5 {
		return "Число простое."
	}
	if n%2 == 0 || n%3 == 0 || n%5 == 0 {
		return "Число нeпростое."
	} else {
		return "Число простое."
	}
}

func main() {

	var num int
	fmt.Print("Введите число: ")
	fmt.Scan(&num)

	var number Number = Number(num)

	fmt.Println(number.IsSimple())

}

package main

import (
	"fmt"
)

func main() {
	var input string
	fmt.Print("Введите строку: ")
	fmt.Scanln(&input)

	count := 0

	for _, value := range input {
		switch string(value) {
		case "у", "е", "ы", "а", "о", "э", "я", "и", "ю":
			count++
		}
	}
	fmt.Print(count)

}

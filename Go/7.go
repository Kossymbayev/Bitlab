package main

import (
	"fmt"
	"unicode"
)

func main() {

	// Lesson strings
	/*
		name := "Эрес"
		// Bad
		for i := 0; i < len(name); i++ {
			fmt.Println(string(name[i]))
		}
		// Good
		for _, val := range name {
			fmt.Println(string(val))
		}
	*/

	// 1
	/*
		var name string
		fmt.Println("Введите имя: ")
		_, err := fmt.Scan(&name)
		if err != nil {
			log.Println(err)
			return
		}

		var runesSlice []rune
		// runesSlice := []rune(text)
		for _, value := range name {
			runesSlice = append(runesSlice, value)
			//fmt.Println(string(value))
		}
		fmt.Println(len(runesSlice))
	*/

	// 2
	/*
		var name string
		fmt.Println("Введите имя: ")
		_, err := fmt.Scan(&name)
		if err != nil {
			log.Println(err)
			return
		}

		var runesSlice []rune
		for _, value := range name {
			runesSlice = append(runesSlice, value)
		}
		fmt.Println(string(runesSlice[0]))
	*/

	// 3
	/*
		var name string
		fmt.Println("Введите строку: ")
		_, err := fmt.Scan(&name)
		if err != nil {
			log.Println(err)
			return
		}

		if name == "BITLAB" {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	*/

	// 4
	/*
		var name string
		fmt.Println("Введите строку: ")
		_, err := fmt.Scan(&name)
		if err != nil {
			log.Println(err)
			return
		}

		if strings.ToLower(name) == "golang" {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	*/

	// 5
	/*
		var s1 string
		fmt.Println("Введите строку: ")
		_, err := fmt.Scan(&s1)
		if err != nil {
			log.Println(err)
			return
		}

		var s2 string
		fmt.Println("Введите строку: ")
		_, err = fmt.Scan(&s2)
		if err != nil {
			log.Println(err)
			return
		}

		if strings.ToLower(s1) == strings.ToLower(s2) {
			fmt.Println("THEY ARE EQUAL")
		} else {
			fmt.Println("THEY ARE NOT EQUAL")
		}
	*/

	// 6
	/*
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Введите текст: ")
		input, _ := reader.ReadString('\n')

		input = strings.TrimSuffix(input, "\n")
		input = strings.TrimSuffix(input, "\r")

		if strings.HasPrefix(input, " ") && strings.HasSuffix(input, " ") {
			input = strings.TrimSpace(input)
			fmt.Printf("(%s)\n", input)
		} else {
			fmt.Println(input)
		}
	*/

	// 7
	/*
		var input string
		fmt.Println("Введите строку: ")
		_, err := fmt.Scan(&input)
		if err != nil {
			log.Println(err)
			return
		}

		switch strings.ToLower(input) {
		case "zero":
			fmt.Println(0)
		case "one":
			fmt.Println(1)
		case "two":
			fmt.Println(2)
		case "three":
			fmt.Println(3)
		case "four":
			fmt.Println(4)
		case "five":
			fmt.Println(5)
		case "six":
			fmt.Println(6)
		case "seven":
			fmt.Println(7)
		case "eight":
			fmt.Println(8)
		case "nine":
			fmt.Println(9)
		}
	*/

	// 8
	/*
		var input string
		fmt.Println("Введите строку: ")
		_, err := fmt.Scan(&input)
		if err != nil {
			log.Println(err)
			return
		}

		for _, char := range input {
			fmt.Println(string(char))
		}
	*/

	// 9
	/*
		var input string
		fmt.Println("Введите строку: ")
		_, err := fmt.Scan(&input)
		if err != nil {
			log.Println(err)
			return
		}

		for _, char := range input {
			fmt.Print(string(char), string(char))
		}
	*/

	// 10
	/*
		var input string
		var letter string

		fmt.Println("Введите строку: ")
		_, err := fmt.Scan(&input)
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println("Введите букву: ")
		fmt.Scan(&letter)

		count := 0

		for _, char := range input {
			if string(char) == letter {
				count++
			}
		}
		fmt.Printf("Встречается %v раз", count)
	*/

	// 11
	/*
		var input string
		fmt.Println("Введите строку: ")
		_, err := fmt.Scan(&input)
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println(strings.ReplaceAll(input, "a", "u"))
	*/

	// 12
	/*
		var input string
		fmt.Println("Введите строку: ")
		_, err := fmt.Scan(&input)
		if err != nil {
			log.Println(err)
			return
		}

		reversed := ""

		for i := len(input) - 1; i >= 0; i-- {
			reversed += string(input[i])
		}

		fmt.Println(reversed)
	*/

	// 13
	/*
		var s1, s2 string

		fmt.Print("Введите строку s1: ")
		fmt.Scanln(&s1)

		fmt.Print("Введите строку s2: ")
		fmt.Scanln(&s2)

		if strings.Contains(s1, s2) {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	*/

	// 14
	/*
		var input string

		fmt.Print("Введите строку: ")
		fmt.Scanln(&input)

		input = strings.ToLower(input)
		targets := "aeiou"
		count := 0

		for _, char := range input {
			if strings.ContainsRune(targets, char) {
				count++
			}
		}

		fmt.Println(count)
	*/

	// 15
	/*
		var input string

		fmt.Print("Введите строку: ")
		fmt.Scanln(&input)

		for i := 0; i < len(input); i++ {
			if i%2 == 1 {
				fmt.Printf(string(input[i]))
			}
		}
	*/

	// 16
	/*
		var input string
		fmt.Print("Введите строку: ")
		fmt.Scanln(&input)

		frequency := make(map[rune]int)
		for _, char := range input {
			frequency[char]++
		}

		result := ""
		for _, char := range input {
			if frequency[char] == 1 {
				result += string(char)
			}
		}

		fmt.Println(result)
	*/

	// 17
	/*
		var input string
		fmt.Print("Введите строку: ")
		fmt.Scanln(&input)

		input = strings.ToLower(input)

		isPalindrome := true
		for i := 0; i < len(input)/2; i++ {
			if input[i] != input[len(input)-1-i] {
				isPalindrome = false
				break
			}
		}

		if isPalindrome {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	*/

	// 18
	/*
		var input string
		fmt.Print("Введите строку: ")
		fmt.Scanln(&input)

		for _, char := range input {
			if unicode.IsUpper(char) {
				fmt.Printf(string(unicode.ToLower(char)))
			} else {
				fmt.Printf(string(unicode.ToUpper(char)))
			}
		}
	*/

	// 19
	/*
		var input string
		fmt.Print("Введите строку: ")
		fmt.Scanln(&input)

		for _, char := range input {
			if unicode.IsDigit(char) {
				fmt.Print("0")
			} else {
				fmt.Print(string(char))
			}
		}
	*/

	// 20
	var input string
	fmt.Print("Введите строку: ")
	fmt.Scanln(&input)

	count := 0

	for _, char := range input {
		if unicode.IsLetter(char) {
			count++
		}
	}

	fmt.Print(count)
}

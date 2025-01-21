package main

import (
	"fmt"
)

func getImt(weight, height float32) (float32, string) {
	imt := calculateImt(weight, height)

	var result string

	if imt <= 18.5 {
		result = "Недостаточная масса"
	} else if imt <= 24.99 {
		result = "Нормальное масса"
	} else if imt <= 29.99 {
		result = "Избыточная масса"
	} else {
		result = "Жирная масса"
	}

	return imt, result
}

func calculateImt(weight, height float32) float32 {
	imt := weight / (height * height)
	return imt
}

func main() {
	imt, diagnos := getImt(80, 1.80)
	fmt.Printf("Ваш индекс массы тела: %.2f. У вас %v", imt, diagnos)
}

package main

import "fmt"

type Student struct {
	name   string
	grades []float64
}

func GetAverageGrade(student Student) {
	allValues := 0.0
	for _, value := range student.grades {
		allValues += value
	}
	fmt.Printf("%v", allValues/float64(len(student.grades)))
}

func main() {
	student := Student{
		name:   "Marc",
		grades: []float64{85, 90, 92, 88, 95},
	}
	GetAverageGrade(student)
}

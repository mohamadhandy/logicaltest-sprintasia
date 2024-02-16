package main

import (
	"fmt"
	"math"
)

func main() {
	var totalInput int
	fmt.Scan(&totalInput)
	var grades []int
	for i := 0; i < totalInput; i++ {
		var input int
		fmt.Scan(&input)
		grades = append(grades, input)
	}
	grades = gradingStudents(grades)
	fmt.Println("Output: ")
	for _, grade := range grades {
		fmt.Println(grade)
	}
}

func gradingStudents(grades []int) []int {
	for i, grade := range grades {
		if grade < 38 {
			continue
		}
		nextMultipleOf5 := int(math.Ceil(float64(grade)/5.0)) * 5
		diff := nextMultipleOf5 - grade
		if diff < 3 {
			grades[i] = nextMultipleOf5
		}
	}
	return grades
}

package main

import "fmt"

func main() {
	grades := map[string]float64{
		"Alice":   85.5,
		"Bob":     78.0,
		"Charlie": 92.3,
		"Diana":   88.8,
	}

	//print grades
	for student, grade := range grades {
		fmt.Printf("%s's grade: %.1f\n", student, grade)
	}

	//
}

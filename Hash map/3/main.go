package main

import "fmt"

func main() {
	grades := map[string]float64{
		"Alice":   85.5,
		"Bob":     78.0,
		"Charlie": 92.7,
		"Diana":   88.8,
	}

	for student, grade := range grades {
		fmt.Printf("%s's grade: %.1f\n",student,grade)
	}

	grades["Eve"] = 98.9

	total := 0.0
	for _,grade := range grades{
		total += grade
	}
	average := total/float64(len(grades))

	fmt.Printf("Class average grade is: %.2f\n",average)
}
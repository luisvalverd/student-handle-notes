package main

import (
	"fmt"

	"example.com/students/models"
)


func main() {
	var students []models.Student

	students = readStudents(&students)
	students = readGrades(students)	

	fmt.Printf("%35s -- %-15s -- %-8s %-8s %-8s %-8s %-8s --> %s\n\n",
		"FULL NAME", "GENDER", "MATHS", "LENG", "GEO", "ENGLISH", "HISTORY", "APROVES")

	for i := 0; i <= 25; i++ {
		fmt.Println(students[i])
	}
}

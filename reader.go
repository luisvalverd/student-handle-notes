package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"example.com/students/models"
)

func readStudents(student *[]models.Student) []models.Student  {
	file, err := os.Open("STUDENTS.csv")	

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close() 

	scanner := bufio.NewScanner(file)

	scanner.Scan() // move firts line

	for scanner.Scan() {
		text := strings.Split(scanner.Text(), ",")
		id, _ := strconv.Atoi(text[0])

		newStudent := models.Student {
			Id: id,
			Name: text[1],
			Gender: text[2],
			Signatures: []models.Signature{},
		} 

		*student = append(*student, newStudent)
		
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	
	file.Close()

	return *student
}

func readGrades(student []models.Student) []models.Student {
	file, err := os.Open("GRADES.csv")	
	var index int = 0

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()

	for scanner.Scan() {

		var grades[3]float64
		text := strings.Split(scanner.Text(), ",")
		id, err:= strconv.Atoi(text[0])

		if err != nil {
			log.Fatal(err)
		}

		for i, grade := range text[1:4] {
			aux, err := strconv.ParseFloat(grade, 64)
			
			if err != nil {
				log.Fatal(err)
			}

			grades[i] = aux
		}

		student[index].Signatures = append(student[index].Signatures, models.Signature{Notes: grades})

		if  id % 5 == 0{
			index++
		}
	}

	file.Close()

	return student
}
package models

import "fmt"

type Student struct {
	Id         int
	Name       string
	Gender     string
	Signatures []Signature
}

func (student Student) GetNameStudent() string {
	return student.Name
}

// return if all signatures are approved
func (student Student) ApprovedAll() bool {
	for _, signature := range student.Signatures {
		if !signature.Approved() {
			return false
		}
	}

	return true
}

func (student Student) String() string {
	var averages[5]float64
	var s = "APPROVED"

	for i, average := range student.Signatures {
		averages[i] = average.GetAverage()
	}

	if !student.ApprovedAll() {
		s = "REPROVED"
	}

	return fmt.Sprintf("%35s -- %-15s -- %-8.2f %-8.2f %-8.2f %-8.2f %-8.2f --> %s",
		student.Name, student.Gender, averages[0], averages[1], averages[2], averages[3], averages[4], s)
}
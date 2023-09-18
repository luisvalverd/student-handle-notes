package models

type Signature struct {
	Notes [3]float64
}

// obtener lista notas
func (s Signature) GetNotes() []float64 {
	return s.Notes[:]
}

// promedio de notas
func (s Signature) GetAverage() float64 {
	var sum float64 = 0
	for _, n := range s.Notes {
		sum += n
	}

	return sum / 3
}

func (s Signature) Approved() bool {
	return s.GetAverage() >= 7
}
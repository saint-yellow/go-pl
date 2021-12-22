package business

import (
	"fmt"
	"sync"
)

type Student struct {
	ID int
	Name string
	Grades []Grade
}

func (s Student) Average() float32 {
	var result float32
	for _, grade := range s.Grades {
		result += grade.Score
	} 
	return result / float32(len(s.Grades))
}

type Students []Student

func (s Students) GetByID(id int) (*Student, error) {
	for i := range s {
		if s[i].ID == id {
			return &s[i], nil
		}
	}

	return nil, fmt.Errorf("student with ID %d not found", id)
}

var (
	students Students
	studentsMutex *sync.Mutex
)

type GradeType string

const (
	GradeQuiz = GradeType("Quiz")
	GradeTest = GradeType("Test")
	GradeExam = GradeType("Exam")
)

type Grade struct {
	Title string
	Type GradeType
	Score float32
}

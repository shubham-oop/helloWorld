package service

import (
	"database/sql"

	"github.com/shubham-oop/helloWorld/internal/model"
	"github.com/shubham-oop/helloWorld/internal/repository"
)

func FetchStudents() ([]model.Student, error) {
	return repository.GetAllStudents()
}

func FetchStudentByRollNumber(rollNumber int) (*model.Student, error) {
	return repository.GetStudentByRollNumber(rollNumber)
}

func CreateStudent(student model.Student) (int, error) {
	existing, err := repository.GetStudentByRollNumber(student.RollNumber)
	if err == nil && existing != nil {
		// Roll number exists
		return 0, sql.ErrNoRows // We'll use this in the handler to return conflict
	}

	return repository.InsertStudent(student)
}

package repository

import (
	"github.com/shubham-oop/helloWorld/db"
	"github.com/shubham-oop/helloWorld/internal/model"
)

func GetAllStudents() ([]model.Student, error) {
	var students []model.Student
	query := `SELECT name, roll_number, class, age, phone_number, email, address
        FROM students ORDER BY roll_number ASC`
	err := db.DB.Select(&students, query)
	return students, err
}

func GetStudentByRollNumber(rollNumber int) (*model.Student, error) {
	var student model.Student
	query := `
        SELECT name, roll_number, class, age, phone_number, email, address
        FROM students
        WHERE roll_number = $1
    `
	err := db.DB.Get(&student, query, rollNumber)
	if err != nil {
		return nil, err
	}
	return &student, nil
}

func InsertStudent(student model.Student) (int, error) {
	query := `
        INSERT INTO students (name, roll_number, class, age, phone_number, email, address)
        VALUES (:name, :roll_number, :class, :age, :phone_number, :email, :address)
        RETURNING id
    `

	stmt, err := db.DB.PrepareNamed(query)
	if err != nil {
		return 0, err
	}

	var id int
	err = stmt.Get(&id, student)
	return id, err
}

func DeleteStudentByRollNumber(rollNumber int) (bool, error) {
	query := `DELETE FROM students WHERE roll_number = $1`
	result, err := db.DB.Exec(query, rollNumber)
	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}

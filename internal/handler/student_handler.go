package handler

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shubham-oop/helloWorld/internal/model"
	"github.com/shubham-oop/helloWorld/internal/service"
)

func GetStudents(c *gin.Context) {
	students, err := service.FetchStudents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch students"})
		return
	}
	c.JSON(http.StatusOK, students)
}

func GetStudentByRollNumber(c *gin.Context) {
	rollNumberStr := c.Param("roll_number")
	rollNumber, err := strconv.Atoi(rollNumberStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid roll number"})
		return
	}

	student, err := service.FetchStudentByRollNumber(rollNumber)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	c.JSON(http.StatusOK, student)
}

func AddStudent(c *gin.Context) {
	var student model.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	id, err := service.CreateStudent(student)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusConflict, gin.H{"error": "Student with this roll number already exists"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create student"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Student created", "id": id})
}
